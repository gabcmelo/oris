package jwtutil

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	secret []byte
}

func New(secret string) *Manager {
	return &Manager{secret: []byte(secret)}
}

func (m *Manager) GenerateUserToken(userID string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(ttl).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(m.secret)
}

func (m *Manager) ParseUserToken(tokenRaw string) (string, error) {
	tok, err := jwt.Parse(tokenRaw, func(t *jwt.Token) (any, error) {
		return m.secret, nil
	})
	if err != nil || !tok.Valid {
		return "", err
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return "", jwt.ErrTokenMalformed
	}
	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return "", jwt.ErrTokenInvalidClaims
	}
	return sub, nil
}
