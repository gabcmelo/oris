package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"safeguild/backend/internal/modules/auth/domain"
	"safeguild/backend/internal/modules/auth/usecase"
)

type Handler struct {
	svc *usecase.Service
}

func NewHandler(svc *usecase.Service) *Handler {
	return &Handler{svc: svc}
}

func userID(c *gin.Context) string {
	v, _ := c.Get("userID")
	s, _ := v.(string)
	return s
}

func (h *Handler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	user, access, refresh, err := h.svc.Register(c.Request.Context(), req.Email, req.Username, req.Password)
	if err != nil {
		switch err {
		case domain.ErrInvalidPayload:
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		case domain.ErrUsernameTaken:
			c.JSON(http.StatusConflict, gin.H{"error": "username taken"})
		case domain.ErrHashingFailed:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":         gin.H{"id": user.ID, "email": user.Email, "username": user.Username},
		"accessToken":  access,
		"refreshToken": refresh,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	access, refresh, err := h.svc.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		if err == domain.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accessToken": access, "refreshToken": refresh})
}

func (h *Handler) Refresh(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refreshToken"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.RefreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	access, refresh, err := h.svc.Refresh(c.Request.Context(), req.RefreshToken)
	if err != nil {
		if err == domain.ErrInvalidRefreshToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accessToken": access, "refreshToken": refresh})
}

func (h *Handler) Logout(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refreshToken"`
	}
	_ = c.ShouldBindJSON(&req)
	_ = h.svc.Logout(c.Request.Context(), req.RefreshToken)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) Me(c *gin.Context) {
	u, err := h.svc.Me(c.Request.Context(), userID(c))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": u})
}
