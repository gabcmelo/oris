package auth

import (
	"github.com/jackc/pgx/v5/pgxpool"
	corejwt "oris/backend/internal/core/auth/jwt"
	"oris/backend/internal/modules/auth/infra"
	authhttp "oris/backend/internal/modules/auth/transport/http"
	"oris/backend/internal/modules/auth/usecase"
)

type Handler = authhttp.Handler

func NewHandler(db *pgxpool.Pool, tokens *corejwt.Manager) *Handler {
	repo := infra.NewPostgresRepository(db)
	svc := usecase.NewService(repo, tokens)
	return authhttp.NewHandler(svc)
}
