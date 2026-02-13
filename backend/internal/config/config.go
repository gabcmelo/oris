package config

import (
	"os"
	"strings"
)

type Config struct {
	HTTPAddr         string
	DatabaseURL      string
	JWTSecret        string
	AppVersion       string
	AppChannel       string
	LivekitURL       string
	LivekitPublicURL string
	LivekitAPIKey    string
	LivekitAPISecret string
	AllowedOrigins   []string
}

func Load() Config {
	return Config{
		HTTPAddr:         envOr("HTTP_ADDR", ":8080"),
		DatabaseURL:      envOr("DATABASE_URL", "postgres://safeguild:safeguild@postgres:5432/safeguild?sslmode=disable"),
		JWTSecret:        envOr("JWT_SECRET", "dev-secret"),
		AppVersion:       envOr("APP_VERSION", "0.1.0"),
		AppChannel:       envOr("APP_CHANNEL", "stable"),
		LivekitURL:       envOr("LIVEKIT_URL", "ws://livekit:7880"),
		LivekitPublicURL: envOr("LIVEKIT_PUBLIC_URL", "ws://localhost:7880"),
		LivekitAPIKey:    envOr("LIVEKIT_API_KEY", "devkey"),
		LivekitAPISecret: envOr("LIVEKIT_API_SECRET", "secret"),
		AllowedOrigins:   splitCSV(envOr("APP_ALLOWED_ORIGINS", "http://localhost:5173")),
	}
}

func envOr(k, fallback string) string {
	v := os.Getenv(k)
	if v == "" {
		return fallback
	}
	return v
}

func splitCSV(v string) []string {
	parts := strings.Split(v, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		s := strings.TrimSpace(p)
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
