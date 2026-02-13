package main

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func buildAllowedOrigins(origins []string) map[string]struct{} {
	allowed := make(map[string]struct{}, len(origins))
	for _, origin := range origins {
		normalized, ok := normalizeOrigin(origin)
		if !ok {
			continue
		}
		allowed[normalized] = struct{}{}
	}
	return allowed
}

func normalizeOrigin(origin string) (string, bool) {
	origin = strings.TrimSpace(origin)
	if origin == "" {
		return "", false
	}
	if origin == "*" {
		return "*", true
	}
	u, err := url.Parse(origin)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return "", false
	}
	if u.Path != "" && u.Path != "/" {
		return "", false
	}
	switch strings.ToLower(u.Scheme) {
	case "http", "https":
		return strings.ToLower(u.Scheme) + "://" + strings.ToLower(u.Host), true
	default:
		return "", false
	}
}

func isOriginAllowed(allowedOrigins map[string]struct{}, origin string) bool {
	if len(allowedOrigins) == 0 {
		return false
	}
	if _, ok := allowedOrigins["*"]; ok {
		return true
	}
	normalized, ok := normalizeOrigin(origin)
	if !ok {
		return false
	}
	_, ok = allowedOrigins[normalized]
	return ok
}

func corsMiddleware(allowedOrigins map[string]struct{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" {
			if !isOriginAllowed(allowedOrigins, origin) {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "origin not allowed"})
				return
			}
			if _, ok := allowedOrigins["*"]; ok {
				c.Header("Access-Control-Allow-Origin", "*")
			} else {
				normalized, _ := normalizeOrigin(origin)
				c.Header("Access-Control-Allow-Origin", normalized)
			}
			c.Header("Vary", "Origin")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		}
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
