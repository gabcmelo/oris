package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"oris/backend/internal/modules/auth/domain"
)

type Service struct {
	repo   domain.Repository
	tokens domain.TokenIssuer
}

func NewService(repo domain.Repository, tokens domain.TokenIssuer) *Service {
	return &Service{repo: repo, tokens: tokens}
}

func (s *Service) Register(ctx context.Context, email, username, password string) (domain.User, string, string, error) {
	if username == "" || len(password) < 6 {
		return domain.User{}, "", "", domain.ErrInvalidPayload
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, "", "", domain.ErrHashingFailed
	}

	id := uuid.NewString()
	if err := s.repo.CreateUser(ctx, id, email, username, string(hash)); err != nil {
		return domain.User{}, "", "", domain.ErrUsernameTaken
	}

	access, err := s.tokens.GenerateUserToken(id, 15*time.Minute)
	if err != nil {
		return domain.User{}, "", "", domain.ErrTokenIssueFailed
	}
	refresh, err := s.tokens.GenerateUserToken(id, 24*time.Hour)
	if err != nil {
		return domain.User{}, "", "", domain.ErrTokenIssueFailed
	}
	_ = s.repo.StoreRefreshToken(ctx, refresh, id, time.Now().Add(24*time.Hour))

	return domain.User{
		ID:       id,
		Email:    email,
		Username: username,
	}, access, refresh, nil
}

func (s *Service) Login(ctx context.Context, username, password string) (string, string, error) {
	uid, hash, err := s.repo.FindCredentialsByUsername(ctx, username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return "", "", domain.ErrInvalidCredentials
	}

	access, err := s.tokens.GenerateUserToken(uid, 15*time.Minute)
	if err != nil {
		return "", "", domain.ErrTokenIssueFailed
	}
	refresh, err := s.tokens.GenerateUserToken(uid, 24*time.Hour)
	if err != nil {
		return "", "", domain.ErrTokenIssueFailed
	}
	_ = s.repo.StoreRefreshToken(ctx, refresh, uid, time.Now().Add(24*time.Hour))
	return access, refresh, nil
}

func (s *Service) Refresh(ctx context.Context, refreshToken string) (string, string, error) {
	if refreshToken == "" {
		return "", "", domain.ErrInvalidPayload
	}
	uid, err := s.repo.FindRefreshTokenOwner(ctx, refreshToken, time.Now())
	if err != nil {
		return "", "", domain.ErrInvalidRefreshToken
	}
	_ = s.repo.DeleteRefreshToken(ctx, refreshToken)

	access, err := s.tokens.GenerateUserToken(uid, 15*time.Minute)
	if err != nil {
		return "", "", domain.ErrTokenIssueFailed
	}
	newRefresh, err := s.tokens.GenerateUserToken(uid, 24*time.Hour)
	if err != nil {
		return "", "", domain.ErrTokenIssueFailed
	}
	_ = s.repo.StoreRefreshToken(ctx, newRefresh, uid, time.Now().Add(24*time.Hour))
	return access, newRefresh, nil
}

func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	if refreshToken == "" {
		return nil
	}
	return s.repo.DeleteRefreshToken(ctx, refreshToken)
}

func (s *Service) Me(ctx context.Context, userID string) (domain.User, error) {
	u, err := s.repo.FindUserByID(ctx, userID)
	if err != nil {
		return domain.User{}, domain.ErrUserNotFound
	}
	return u, nil
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
