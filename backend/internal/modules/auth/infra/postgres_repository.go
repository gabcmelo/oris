package infra

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"oris/backend/internal/modules/auth/domain"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) CreateUser(ctx context.Context, id, email, username, passwordHash string) error {
	_, err := r.db.Exec(ctx, `insert into users(id,email,username,password_hash) values($1,$2,$3,$4)`, id, email, username, passwordHash)
	return err
}

func (r *PostgresRepository) FindCredentialsByUsername(ctx context.Context, username string) (string, string, error) {
	var uid, hash string
	err := r.db.QueryRow(ctx, `select id,password_hash from users where username=$1`, username).Scan(&uid, &hash)
	return uid, hash, err
}

func (r *PostgresRepository) StoreRefreshToken(ctx context.Context, token, userID string, _ time.Time) error {
	_, err := r.db.Exec(ctx, `insert into refresh_tokens(token,user_id,expires_at) values($1,$2,now()+ interval '24 hours')`, token, userID)
	return err
}

func (r *PostgresRepository) FindRefreshTokenOwner(ctx context.Context, token string, _ time.Time) (string, error) {
	var uid string
	err := r.db.QueryRow(ctx, `select user_id from refresh_tokens where token=$1 and expires_at > now()`, token).Scan(&uid)
	return uid, err
}

func (r *PostgresRepository) DeleteRefreshToken(ctx context.Context, token string) error {
	_, err := r.db.Exec(ctx, `delete from refresh_tokens where token=$1`, token)
	return err
}

func (r *PostgresRepository) FindUserByID(ctx context.Context, userID string) (domain.User, error) {
	var u domain.User
	err := r.db.QueryRow(ctx, `select id,email,username,created_at from users where id=$1`, userID).Scan(&u.ID, &u.Email, &u.Username, &u.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}
	return u, nil
}
