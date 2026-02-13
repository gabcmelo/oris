package middleware

import (
	legacy "oris/backend/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func Auth(parseToken func(string) (string, error)) gin.HandlerFunc {
	return legacy.Auth(parseToken)
}
