package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNormalizeOrigin_ValidHTTPURL_Normalized(t *testing.T) {
	got, ok := normalizeOrigin("HTTP://LOCALHOST:5173")
	if !ok {
		t.Fatalf("expected valid origin")
	}
	if got != "http://localhost:5173" {
		t.Fatalf("unexpected normalized origin: %s", got)
	}
}

func TestNormalizeOrigin_WithPath_ReturnsInvalid(t *testing.T) {
	_, ok := normalizeOrigin("http://localhost:5173/app")
	if ok {
		t.Fatalf("expected invalid origin when path is present")
	}
}

func TestIsOriginAllowed_ExactMatch_ReturnsTrue(t *testing.T) {
	allowed := buildAllowedOrigins([]string{"http://localhost:5173"})
	if !isOriginAllowed(allowed, "http://localhost:5173") {
		t.Fatalf("expected origin to be allowed")
	}
}

func TestIsOriginAllowed_NotListed_ReturnsFalse(t *testing.T) {
	allowed := buildAllowedOrigins([]string{"http://localhost:5173"})
	if isOriginAllowed(allowed, "http://localhost:3000") {
		t.Fatalf("expected origin to be denied")
	}
}

func TestCorsMiddleware_DisallowedOrigin_ReturnsForbidden(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(corsMiddleware(buildAllowedOrigins([]string{"http://localhost:5173"})))
	r.GET("/ok", func(c *gin.Context) { c.Status(http.StatusOK) })

	req := httptest.NewRequest(http.MethodGet, "/ok", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Fatalf("expected status 403, got %d", w.Code)
	}
}

func TestCorsMiddleware_PreflightAllowedOrigin_ReturnsNoContent(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(corsMiddleware(buildAllowedOrigins([]string{"http://localhost:5173"})))
	r.OPTIONS("/ok", func(c *gin.Context) { c.Status(http.StatusOK) })

	req := httptest.NewRequest(http.MethodOptions, "/ok", nil)
	req.Header.Set("Origin", "http://localhost:5173")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", w.Code)
	}
	if got := w.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:5173" {
		t.Fatalf("unexpected allow origin header: %q", got)
	}
}
