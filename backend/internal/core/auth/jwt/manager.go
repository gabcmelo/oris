package jwt

import legacy "safeguild/backend/internal/platform/jwtutil"

type Manager = legacy.Manager

func New(secret string) *Manager {
	return legacy.New(secret)
}
