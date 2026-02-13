package main

import (
	"context"
	"encoding/csv"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
	bootstrap "oris/backend/internal/app"
	httpapp "oris/backend/internal/app/http"
	httpmiddleware "oris/backend/internal/app/http/middleware"
	corejwt "oris/backend/internal/core/auth/jwt"
	appconfig "oris/backend/internal/core/config"
	authmodule "oris/backend/internal/modules/auth"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}

type Community struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	OwnerUserID     string    `json:"ownerUserId"`
	SafeModeEnabled bool      `json:"safeModeEnabled"`
	CreatedAt       time.Time `json:"createdAt"`
}

type Channel struct {
	ID          string    `json:"id"`
	CommunityID string    `json:"communityId"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CommunityMember struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Muted    bool   `json:"muted"`
	Banned   bool   `json:"banned"`
}

type Message struct {
	ID        string    `json:"id"`
	ChannelID string    `json:"channelId"`
	AuthorID  string    `json:"authorUserId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type AuditLog struct {
	ID          string                 `json:"id"`
	CommunityID string                 `json:"communityId"`
	ActorUserID string                 `json:"actorUserId"`
	ActionType  string                 `json:"actionType"`
	Target      map[string]interface{} `json:"target"`
	CreatedAt   time.Time              `json:"createdAt"`
}

type TelemetryState struct {
	Enabled    bool      `json:"enabled"`
	LastSentAt time.Time `json:"lastSentAt"`
}

type wsHub struct {
	mu       sync.RWMutex
	channels map[string]map[*websocket.Conn]string
}

func newHub() *wsHub { return &wsHub{channels: map[string]map[*websocket.Conn]string{}} }

func (h *wsHub) register(channelID, uid string, c *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.channels[channelID]; !ok {
		h.channels[channelID] = map[*websocket.Conn]string{}
	}
	h.channels[channelID][c] = uid
}

func (h *wsHub) unregister(channelID string, c *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.channels[channelID]; !ok {
		return
	}
	delete(h.channels[channelID], c)
	if len(h.channels[channelID]) == 0 {
		delete(h.channels, channelID)
	}
}

func (h *wsHub) broadcastMessage(channelID string, msg Message) {
	h.mu.RLock()
	clients := h.channels[channelID]
	h.mu.RUnlock()
	payload := gin.H{"type": "message.created", "data": msg}
	for c := range clients {
		_ = c.WriteJSON(payload)
	}
}

func (h *wsHub) userIDs(channelID string) []string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	clients := h.channels[channelID]
	seen := map[string]bool{}
	ids := make([]string, 0, len(clients))
	for _, uid := range clients {
		if uid == "" || seen[uid] {
			continue
		}
		seen[uid] = true
		ids = append(ids, uid)
	}
	return ids
}

func (h *wsHub) broadcastPresence(channelID string, users []map[string]string) {
	h.mu.RLock()
	clients := h.channels[channelID]
	h.mu.RUnlock()
	payload := gin.H{"type": "presence.updated", "data": gin.H{"channelId": channelID, "users": users}}
	for c := range clients {
		_ = c.WriteJSON(payload)
	}
}

type App struct {
	db               *pgxpool.Pool
	hub              *wsHub
	jwt              *corejwt.Manager
	livekitURL       string
	livekitPublicURL string
	livekitAPIKey    string
	livekitAPISecret string
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func main() {
	cfg := appconfig.Load()
	allowedOrigins := buildAllowedOrigins(cfg.AllowedOrigins)

	db, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &App{
		db:               db,
		hub:              newHub(),
		jwt:              corejwt.New(cfg.JWTSecret),
		livekitURL:       cfg.LivekitURL,
		livekitPublicURL: cfg.LivekitPublicURL,
		livekitAPIKey:    cfg.LivekitAPIKey,
		livekitAPISecret: cfg.LivekitAPISecret,
	}
	authHandler := authmodule.NewHandler(db, app.jwt)

	r := gin.Default()
	r.Use(corsMiddleware(allowedOrigins))
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return isOriginAllowed(allowedOrigins, r.Header.Get("Origin"))
	}

	bootstrap.BuildHTTPRouter(r, httpapp.Dependencies{
		AppVersion:     cfg.AppVersion,
		AppChannel:     cfg.AppChannel,
		AuthMiddleware: httpmiddleware.Auth(app.parseUserFromJWT),
		Handlers: httpapp.Handlers{
			Register:         authHandler.Register,
			Login:            authHandler.Login,
			Refresh:          authHandler.Refresh,
			Logout:           authHandler.Logout,
			Me:               authHandler.Me,
			CreateCommunity:  app.createCommunityHandler,
			ListCommunities:  app.listCommunitiesHandler,
			GetCommunity:     app.getCommunityHandler,
			ListMembers:      app.listMembersHandler,
			CreateInvite:     app.createInviteHandler,
			JoinInvite:       app.joinInviteHandler,
			CreateChannel:    app.createChannelHandler,
			ListChannels:     app.listChannelsHandler,
			ListMessages:     app.listMessagesHandler,
			ChannelPresence:  app.channelPresenceHandler,
			PostMessage:      app.postMessageHandler,
			ModerationKick:   app.moderationHandler("kick"),
			ModerationMute:   app.moderationHandler("mute"),
			ModerationBan:    app.moderationHandler("ban"),
			ListAuditLogs:    app.listAuditLogsHandler,
			Export:           app.exportHandler,
			VoiceToken:       app.voiceTokenHandler,
			IntegrationEvent: app.integrationEventHandler,
			TelemetryOptIn:   app.telemetryOptInHandler,
			TelemetryStatus:  app.telemetryStatusHandler,
			TelemetryPolicy:  app.telemetryPolicyHandler,
			WS:               app.wsHandler,
		},
	})

	log.Printf("Oris API on :8080")
	if err := r.Run(cfg.HTTPAddr); err != nil {
		log.Fatal(err)
	}
}

func (a *App) parseUserFromJWT(tokenRaw string) (string, error) {
	return a.jwt.ParseUserToken(tokenRaw)
}

func userID(c *gin.Context) string {
	v, _ := c.Get("userID")
	s, _ := v.(string)
	return s
}

func (a *App) addAudit(ctx context.Context, communityID, actor, action string, target string) {
	_, _ = a.db.Exec(ctx, `insert into audit_log(id, community_id, actor_user_id, action_type, target) values($1,$2,$3,$4,$5::jsonb)`, uuid.NewString(), communityID, actor, action, target)
}

func roleAllowed(actual, min string) bool {
	return roleRank(actual) >= roleRank(min)
}

func roleRank(role string) int {
	order := map[string]int{"member": 1, "moderator": 2, "admin": 3, "owner": 4}
	return order[role]
}

func (a *App) membershipRole(ctx context.Context, communityID, uid string) (string, bool, bool) {
	var role string
	var banned bool
	var muted bool
	err := a.db.QueryRow(ctx, `select role, banned, muted from community_members where community_id=$1 and user_id=$2`, communityID, uid).Scan(&role, &banned, &muted)
	if err != nil {
		return "", false, false
	}
	return role, banned, muted
}

func (a *App) hasRole(ctx context.Context, communityID, uid, minRole string) bool {
	role, banned, _ := a.membershipRole(ctx, communityID, uid)
	if banned || role == "" {
		return false
	}
	return roleAllowed(role, minRole)
}

func (a *App) hasMembership(ctx context.Context, communityID, uid string) bool {
	role, banned, _ := a.membershipRole(ctx, communityID, uid)
	return role != "" && !banned
}

func (a *App) createCommunityHandler(c *gin.Context) {
	var body struct {
		Name            string `json:"name"`
		SafeModeEnabled *bool  `json:"safeModeEnabled"`
	}
	if err := c.ShouldBindJSON(&body); err != nil || body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	safe := true
	if body.SafeModeEnabled != nil {
		safe = *body.SafeModeEnabled
	}
	uid := userID(c)
	cid := uuid.NewString()
	tx, err := a.db.Begin(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	defer tx.Rollback(c.Request.Context())
	_, err = tx.Exec(c.Request.Context(), `insert into communities(id,name,owner_user_id,safe_mode_enabled) values($1,$2,$3,$4)`, cid, body.Name, uid, safe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create failed"})
		return
	}
	_, err = tx.Exec(c.Request.Context(), `insert into community_members(community_id,user_id,role) values($1,$2,'owner')`, cid, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "member create failed"})
		return
	}
	_, _ = tx.Exec(c.Request.Context(), `insert into audit_log(id,community_id,actor_user_id,action_type,target) values($1,$2,$3,$4,$5::jsonb)`, uuid.NewString(), cid, uid, "community.create", `{"communityId":"`+cid+`"}`)
	if err = tx.Commit(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "commit failed"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": cid, "name": body.Name, "ownerUserId": uid, "safeModeEnabled": safe, "createdAt": time.Now()})
}

func (a *App) listCommunitiesHandler(c *gin.Context) {
	uid := userID(c)
	rows, err := a.db.Query(c.Request.Context(), `select c.id,c.name,c.owner_user_id,c.safe_mode_enabled,c.created_at from communities c join community_members m on m.community_id=c.id where m.user_id=$1 and m.banned=false order by c.created_at desc`, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	defer rows.Close()
	items := []Community{}
	for rows.Next() {
		var it Community
		if err := rows.Scan(&it.ID, &it.Name, &it.OwnerUserID, &it.SafeModeEnabled, &it.CreatedAt); err == nil {
			items = append(items, it)
		}
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (a *App) getCommunityHandler(c *gin.Context) {
	cid := c.Param("communityId")
	var it Community
	err := a.db.QueryRow(c.Request.Context(), `select id,name,owner_user_id,safe_mode_enabled,created_at from communities where id=$1`, cid).Scan(&it.ID, &it.Name, &it.OwnerUserID, &it.SafeModeEnabled, &it.CreatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, it)
}

func (a *App) listMembersHandler(c *gin.Context) {
	cid := c.Param("communityId")
	if !a.hasMembership(c.Request.Context(), cid, userID(c)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	rows, err := a.db.Query(
		c.Request.Context(),
		`select m.user_id, u.username, m.role, m.muted, m.banned
		 from community_members m
		 join users u on u.id = m.user_id
		 where m.community_id=$1
		 order by m.joined_at asc`,
		cid,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	defer rows.Close()

	items := []CommunityMember{}
	for rows.Next() {
		var m CommunityMember
		if err := rows.Scan(&m.UserID, &m.Username, &m.Role, &m.Muted, &m.Banned); err == nil {
			items = append(items, m)
		}
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (a *App) createInviteHandler(c *gin.Context) {
	cid := c.Param("communityId")
	if !a.hasRole(c.Request.Context(), cid, userID(c), "moderator") {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	var body struct {
		Hours   int `json:"hours"`
		MaxUses int `json:"maxUses"`
	}
	_ = c.ShouldBindJSON(&body)
	if body.Hours <= 0 {
		body.Hours = 24
	}
	if body.MaxUses <= 0 {
		body.MaxUses = 50
	}
	code := strings.ToUpper(strings.ReplaceAll(uuid.NewString()[:8], "-", ""))
	_, err := a.db.Exec(c.Request.Context(), `insert into invites(code,community_id,expires_at,max_uses,uses_count) values($1,$2,now()+make_interval(hours => ($3)::int),$4,0)`, code, cid, body.Hours, body.MaxUses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invite create failed"})
		return
	}
	a.addAudit(c.Request.Context(), cid, userID(c), "invite.create", `{"code":"`+code+`"}`)
	var exp time.Time
	_ = a.db.QueryRow(c.Request.Context(), `select expires_at from invites where code=$1`, code).Scan(&exp)
	c.JSON(http.StatusCreated, gin.H{"code": code, "communityId": cid, "expiresAt": exp, "maxUses": body.MaxUses, "usesCount": 0})
}

func (a *App) joinInviteHandler(c *gin.Context) {
	code := c.Param("code")
	uid := userID(c)
	tx, err := a.db.Begin(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	defer tx.Rollback(c.Request.Context())
	var cid string
	var exp time.Time
	var maxUses, uses int
	err = tx.QueryRow(c.Request.Context(), `select community_id,expires_at,max_uses,uses_count from invites where code=$1 for update`, code).Scan(&cid, &exp, &maxUses, &uses)
	if err != nil || exp.Before(time.Now()) || uses >= maxUses {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invite"})
		return
	}
	memberInsertResult, err := tx.Exec(c.Request.Context(), `insert into community_members(community_id,user_id,role) values($1,$2,'member') on conflict (community_id,user_id) do nothing`, cid, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "member join failed"})
		return
	}
	joined := memberInsertResult.RowsAffected() == 1
	if joined {
		if _, err = tx.Exec(c.Request.Context(), `update invites set uses_count=uses_count+1 where code=$1`, code); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invite update failed"})
			return
		}
		if _, err = tx.Exec(c.Request.Context(), `insert into audit_log(id,community_id,actor_user_id,action_type,target) values($1,$2,$3,$4,$5::jsonb)`, uuid.NewString(), cid, uid, "member.join", `{"userId":"`+uid+`"}`); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "audit log failed"})
			return
		}
	}
	if err = tx.Commit(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "commit failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "communityId": cid, "joined": joined})
}

func (a *App) createChannelHandler(c *gin.Context) {
	cid := c.Param("communityId")
	uid := userID(c)
	if !a.hasRole(c.Request.Context(), cid, uid, "moderator") {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	var body struct{ Name, Type string }
	if err := c.ShouldBindJSON(&body); err != nil || body.Name == "" || (body.Type != "text" && body.Type != "voice") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	id := uuid.NewString()
	_, err := a.db.Exec(c.Request.Context(), `insert into channels(id,community_id,name,type) values($1,$2,$3,$4)`, id, cid, body.Name, body.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create channel failed"})
		return
	}
	a.addAudit(c.Request.Context(), cid, uid, "channel.create", `{"channelId":"`+id+`"}`)
	c.JSON(http.StatusCreated, gin.H{"id": id, "communityId": cid, "name": body.Name, "type": body.Type, "createdAt": time.Now()})
}

func (a *App) listChannelsHandler(c *gin.Context) {
	cid := c.Param("communityId")
	if !a.hasMembership(c.Request.Context(), cid, userID(c)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	rows, err := a.db.Query(c.Request.Context(), `select id,community_id,name,type,created_at from channels where community_id=$1 order by created_at asc`, cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	defer rows.Close()
	items := []Channel{}
	for rows.Next() {
		var ch Channel
		if err := rows.Scan(&ch.ID, &ch.CommunityID, &ch.Name, &ch.Type, &ch.CreatedAt); err == nil {
			items = append(items, ch)
		}
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (a *App) listMessagesHandler(c *gin.Context) {
	chID := c.Param("channelId")
	var cid string
	if err := a.db.QueryRow(c.Request.Context(), `select community_id from channels where id=$1`, chID).Scan(&cid); err != nil || !a.hasMembership(c.Request.Context(), cid, userID(c)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	rows, err := a.db.Query(c.Request.Context(), `select id,channel_id,author_user_id,content,created_at from messages where channel_id=$1 order by created_at asc limit 200`, chID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	defer rows.Close()
	items := []Message{}
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.ChannelID, &m.AuthorID, &m.Content, &m.CreatedAt); err == nil {
			items = append(items, m)
		}
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (a *App) resolvePresenceUsers(ctx context.Context, userIDs []string) []map[string]string {
	if len(userIDs) == 0 {
		return []map[string]string{}
	}
	rows, err := a.db.Query(ctx, `select id, username from users where id = any($1)`, userIDs)
	if err != nil {
		return []map[string]string{}
	}
	defer rows.Close()
	out := []map[string]string{}
	for rows.Next() {
		var id, username string
		if err := rows.Scan(&id, &username); err == nil {
			out = append(out, map[string]string{"userId": id, "username": username})
		}
	}
	return out
}

func (a *App) channelPresenceHandler(c *gin.Context) {
	chID := c.Param("channelId")
	var cid string
	if err := a.db.QueryRow(c.Request.Context(), `select community_id from channels where id=$1`, chID).Scan(&cid); err != nil || !a.hasMembership(c.Request.Context(), cid, userID(c)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	users := a.resolvePresenceUsers(c.Request.Context(), a.hub.userIDs(chID))
	c.JSON(http.StatusOK, gin.H{"channelId": chID, "users": users})
}

func (a *App) postMessageHandler(c *gin.Context) {
	chID := c.Param("channelId")
	uid := userID(c)
	var body struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&body); err != nil || strings.TrimSpace(body.Content) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	var cid string
	if err := a.db.QueryRow(c.Request.Context(), `select community_id from channels where id=$1`, chID).Scan(&cid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "channel not found"})
		return
	}
	role, banned, muted := a.membershipRole(c.Request.Context(), cid, uid)
	if role == "" || banned || muted {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	mid := uuid.NewString()
	createdAt := time.Now()
	_, err := a.db.Exec(c.Request.Context(), `insert into messages(id,channel_id,author_user_id,content,created_at) values($1,$2,$3,$4,$5)`, mid, chID, uid, body.Content, createdAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "message create failed"})
		return
	}
	msg := Message{ID: mid, ChannelID: chID, AuthorID: uid, Content: body.Content, CreatedAt: createdAt}
	a.hub.broadcastMessage(chID, msg)
	c.JSON(http.StatusCreated, msg)
}

func (a *App) moderationHandler(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cid := c.Param("communityId")
		actor := userID(c)
		actorRole, actorBanned, _ := a.membershipRole(c.Request.Context(), cid, actor)
		if actorRole == "" || actorBanned || !roleAllowed(actorRole, "moderator") {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		var body struct {
			TargetUserID string `json:"targetUserId"`
		}
		if err := c.ShouldBindJSON(&body); err != nil || body.TargetUserID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		if body.TargetUserID == actor {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cannot moderate yourself"})
			return
		}
		targetRole, _, _ := a.membershipRole(c.Request.Context(), cid, body.TargetUserID)
		if targetRole == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "target member not found"})
			return
		}
		if roleRank(actorRole) <= roleRank(targetRole) {
			c.JSON(http.StatusForbidden, gin.H{"error": "cannot moderate same or higher role"})
			return
		}
		sql := ""
		switch action {
		case "kick", "ban":
			sql = `update community_members set banned=true where community_id=$1 and user_id=$2`
		case "mute":
			sql = `update community_members set muted=true where community_id=$1 and user_id=$2`
		}
		_, err := a.db.Exec(c.Request.Context(), sql, cid, body.TargetUserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "moderation failed"})
			return
		}
		a.addAudit(c.Request.Context(), cid, actor, "moderation."+action, `{"targetUserId":"`+body.TargetUserID+`"}`)
		c.JSON(http.StatusOK, gin.H{"ok": true})
	}
}

func (a *App) listAuditLogsHandler(c *gin.Context) {
	cid := c.Param("communityId")
	if !a.hasRole(c.Request.Context(), cid, userID(c), "moderator") {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	rows, err := a.db.Query(c.Request.Context(), `select id,community_id,actor_user_id,action_type,target,created_at from audit_log where community_id=$1 order by created_at desc limit 200`, cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	defer rows.Close()
	items := []AuditLog{}
	for rows.Next() {
		var aLog AuditLog
		var target map[string]interface{}
		if err := rows.Scan(&aLog.ID, &aLog.CommunityID, &aLog.ActorUserID, &aLog.ActionType, &target, &aLog.CreatedAt); err == nil {
			aLog.Target = target
			items = append(items, aLog)
		}
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (a *App) exportHandler(c *gin.Context) {
	cid := c.Param("communityId")
	if !a.hasRole(c.Request.Context(), cid, userID(c), "moderator") {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	format := strings.ToLower(c.DefaultQuery("format", "json"))
	rows, err := a.db.Query(c.Request.Context(), `select id,community_id,actor_user_id,action_type,target,created_at from audit_log where community_id=$1 order by created_at desc limit 200`, cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	defer rows.Close()
	items := []AuditLog{}
	for rows.Next() {
		var aLog AuditLog
		var target map[string]interface{}
		if err := rows.Scan(&aLog.ID, &aLog.CommunityID, &aLog.ActorUserID, &aLog.ActionType, &target, &aLog.CreatedAt); err == nil {
			aLog.Target = target
			items = append(items, aLog)
		}
	}
	if format == "csv" {
		c.Header("Content-Disposition", "attachment; filename=audit.csv")
		c.Header("Content-Type", "text/csv")
		w := csv.NewWriter(c.Writer)
		_ = w.Write([]string{"id", "community_id", "actor_user_id", "action_type", "created_at"})
		for _, l := range items {
			_ = w.Write([]string{l.ID, l.CommunityID, l.ActorUserID, l.ActionType, l.CreatedAt.Format(time.RFC3339)})
		}
		w.Flush()
		return
	}
	c.JSON(http.StatusOK, gin.H{"communityId": cid, "auditLogs": items})
}

func (a *App) voiceTokenHandler(c *gin.Context) {
	var body struct {
		ChannelID string `json:"channelId"`
	}
	if err := c.ShouldBindJSON(&body); err != nil || body.ChannelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	uid := userID(c)
	var communityID string
	var channelType string
	if err := a.db.QueryRow(c.Request.Context(), `select community_id,type from channels where id=$1`, body.ChannelID).Scan(&communityID, &channelType); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "channel not found"})
		return
	}
	if channelType != "voice" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel is not voice"})
		return
	}
	if !a.hasMembership(c.Request.Context(), communityID, uid) {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	claims := jwt.MapClaims{
		"iss": a.livekitAPIKey,
		"sub": uid,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(1 * time.Hour).Unix(),
		"video": map[string]any{
			"roomJoin":     true,
			"room":         body.ChannelID,
			"canPublish":   true,
			"canSubscribe": true,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(a.livekitAPISecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create livekit token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"serverUrl": a.livekitPublicURL, "apiKey": a.livekitAPIKey, "room": body.ChannelID, "identity": uid, "token": token})
}

func (a *App) integrationEventHandler(c *gin.Context) {
	var payload map[string]any
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"received": true, "next": "integration stub queued"})
}

func (a *App) telemetryOptInHandler(c *gin.Context) {
	if !a.hasAdminAccess(c.Request.Context(), userID(c)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	var body struct {
		Enabled bool `json:"enabled"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	_, err := a.db.Exec(c.Request.Context(), `insert into telemetry_settings(id,enabled,updated_at) values(true,$1,now()) on conflict (id) do update set enabled=excluded.enabled, updated_at=excluded.updated_at`, body.Enabled)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "enabled": body.Enabled})
}

func (a *App) telemetryStatusHandler(c *gin.Context) {
	if !a.hasAdminAccess(c.Request.Context(), userID(c)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	var st TelemetryState
	err := a.db.QueryRow(c.Request.Context(), `select enabled,updated_at from telemetry_settings where id=true`).Scan(&st.Enabled, &st.LastSentAt)
	if err != nil {
		st.Enabled = false
		st.LastSentAt = time.Time{}
	}
	c.JSON(http.StatusOK, st)
}

func (a *App) telemetryPolicyHandler(c *gin.Context) {
	if !a.hasAdminAccess(c.Request.Context(), userID(c)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"allow": []string{
			"cpu", "memory", "disk", "io", "network",
			"api_latency_p95", "api_latency_p99", "api_error_rate", "login_rate", "rate_limit_hits",
			"ws_active_connections", "ws_reconnections", "channel_join_latency",
			"voice_active_rooms", "voice_participants", "voice_turn_ratio", "voice_packet_loss_aggregate",
		},
		"deny":      []string{"message_content", "audio_content", "email", "username", "raw_ip", "personal_ids"},
		"transport": "otlp",
	})
}

func (a *App) hasAdminAccess(ctx context.Context, uid string) bool {
	var allowed bool
	if err := a.db.QueryRow(ctx, `select exists(select 1 from community_members where user_id=$1 and banned=false and role in ('admin','owner'))`, uid).Scan(&allowed); err != nil {
		return false
	}
	return allowed
}

func (a *App) wsHandler(c *gin.Context) {
	channelID := c.Param("channelId")
	tokenRaw := c.Query("token")
	if tokenRaw == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}
	uid, err := a.parseUserFromJWT(tokenRaw)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}
	var cid string
	if err := a.db.QueryRow(c.Request.Context(), `select community_id from channels where id=$1`, channelID).Scan(&cid); err != nil || !a.hasMembership(c.Request.Context(), cid, uid) {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	a.hub.register(channelID, uid, conn)
	a.hub.broadcastPresence(channelID, a.resolvePresenceUsers(c.Request.Context(), a.hub.userIDs(channelID)))
	defer func() {
		a.hub.unregister(channelID, conn)
		a.hub.broadcastPresence(channelID, a.resolvePresenceUsers(c.Request.Context(), a.hub.userIDs(channelID)))
		_ = conn.Close()
	}()
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}
