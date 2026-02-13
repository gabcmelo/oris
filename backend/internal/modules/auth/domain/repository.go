package domain

import (
	"context"
	"time"
)

type Repository interface {
	CreateUser(ctx context.Context, id, email, username, passwordHash string) error
	FindCredentialsByUsername(ctx context.Context, username string) (userID string, passwordHash string, err error)
	StoreRefreshToken(ctx context.Context, token, userID string, expiresAt time.Time) error
	FindRefreshTokenOwner(ctx context.Context, token string, now time.Time) (userID string, err error)
	DeleteRefreshToken(ctx context.Context, token string) error
	FindUserByID(ctx context.Context, userID string) (User, error)
}

type TokenIssuer interface {
	GenerateUserToken(userID string, ttl time.Duration) (string, error)
}
