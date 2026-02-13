package auth

import (
	"github.com/jackc/pgx/v5/pgxpool"
	corejwt "safeguild/backend/internal/core/auth/jwt"
	"safeguild/backend/internal/modules/auth/infra"
	authhttp "safeguild/backend/internal/modules/auth/transport/http"
	"safeguild/backend/internal/modules/auth/usecase"
)

type Handler = authhttp.Handler

func NewHandler(db *pgxpool.Pool, tokens *corejwt.Manager) *Handler {
	repo := infra.NewPostgresRepository(db)
	svc := usecase.NewService(repo, tokens)
	return authhttp.NewHandler(svc)
}
