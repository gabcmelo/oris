package jwt

import legacy "oris/backend/internal/platform/jwtutil"

type Manager = legacy.Manager

func New(secret string) *Manager {
	return legacy.New(secret)
}
