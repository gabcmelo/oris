package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Register         gin.HandlerFunc
	Login            gin.HandlerFunc
	Refresh          gin.HandlerFunc
	Logout           gin.HandlerFunc
	Me               gin.HandlerFunc
	CreateCommunity  gin.HandlerFunc
	ListCommunities  gin.HandlerFunc
	GetCommunity     gin.HandlerFunc
	ListMembers      gin.HandlerFunc
	CreateInvite     gin.HandlerFunc
	JoinInvite       gin.HandlerFunc
	CreateChannel    gin.HandlerFunc
	ListChannels     gin.HandlerFunc
	ListMessages     gin.HandlerFunc
	ChannelPresence  gin.HandlerFunc
	PostMessage      gin.HandlerFunc
	ModerationKick   gin.HandlerFunc
	ModerationMute   gin.HandlerFunc
	ModerationBan    gin.HandlerFunc
	ListAuditLogs    gin.HandlerFunc
	Export           gin.HandlerFunc
	VoiceToken       gin.HandlerFunc
	IntegrationEvent gin.HandlerFunc
	TelemetryOptIn   gin.HandlerFunc
	TelemetryStatus  gin.HandlerFunc
	TelemetryPolicy  gin.HandlerFunc
	WS               gin.HandlerFunc
}

type Dependencies struct {
	AppVersion     string
	AppChannel     string
	AuthMiddleware gin.HandlerFunc
	Handlers       Handlers
}

func Register(engine *gin.Engine, deps Dependencies) {
	engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	v1 := engine.Group("/api/v1")
	v1.POST("/auth/register", deps.Handlers.Register)
	v1.POST("/auth/login", deps.Handlers.Login)
	v1.POST("/auth/refresh", deps.Handlers.Refresh)
	v1.POST("/auth/logout", deps.Handlers.Logout)
	v1.GET("/system/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": deps.AppVersion, "channel": deps.AppChannel})
	})
	v1.POST("/system/upgrade/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"latestVersion": deps.AppVersion, "channel": deps.AppChannel, "updateAvailable": false})
	})

	auth := v1.Group("/")
	auth.Use(deps.AuthMiddleware)
	auth.GET("/me", deps.Handlers.Me)
	auth.POST("/communities", deps.Handlers.CreateCommunity)
	auth.GET("/communities", deps.Handlers.ListCommunities)
	auth.GET("/communities/:communityId", deps.Handlers.GetCommunity)
	auth.GET("/communities/:communityId/members", deps.Handlers.ListMembers)
	auth.POST("/communities/:communityId/invites", deps.Handlers.CreateInvite)
	auth.POST("/invites/:code/join", deps.Handlers.JoinInvite)
	auth.POST("/communities/:communityId/channels", deps.Handlers.CreateChannel)
	auth.GET("/communities/:communityId/channels", deps.Handlers.ListChannels)
	auth.GET("/channels/:channelId/messages", deps.Handlers.ListMessages)
	auth.GET("/channels/:channelId/presence", deps.Handlers.ChannelPresence)
	auth.POST("/channels/:channelId/messages", deps.Handlers.PostMessage)
	auth.POST("/communities/:communityId/moderation/kick", deps.Handlers.ModerationKick)
	auth.POST("/communities/:communityId/moderation/mute", deps.Handlers.ModerationMute)
	auth.POST("/communities/:communityId/moderation/ban", deps.Handlers.ModerationBan)
	auth.GET("/communities/:communityId/audit-logs", deps.Handlers.ListAuditLogs)
	auth.GET("/communities/:communityId/exports", deps.Handlers.Export)
	auth.POST("/voice/token", deps.Handlers.VoiceToken)
	auth.POST("/integrations/events", deps.Handlers.IntegrationEvent)
	auth.POST("/admin/telemetry/opt-in", deps.Handlers.TelemetryOptIn)
	auth.GET("/admin/telemetry/status", deps.Handlers.TelemetryStatus)
	auth.GET("/admin/telemetry/policy", deps.Handlers.TelemetryPolicy)
	auth.GET("/ws/:channelId", deps.Handlers.WS)
}
