package repository

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rwx03/Pastebin/backend/internal/models"
)

type TokenPostgres struct {
	pool *pgxpool.Pool
}

func NewTokenPostgres(pool *pgxpool.Pool) *TokenPostgres {
	return &TokenPostgres{
		pool: pool,
	}
}

func (t *TokenPostgres) Create(ctx context.Context, token models.Token) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (user_id,refresh_token) VALUES ($1,$2) RETURNING id`, tokensTable)

	if err := t.pool.QueryRow(ctx, query, token.UserID, token.RefreshToken).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TokenPostgres) GetToken(ctx context.Context, token string) (*models.Token, error) {
	var tokenModel *models.Token

	query := fmt.Sprintf(`SELECT * FROM %s WHERE refresh_token = $1`, tokensTable)

	if err := pgxscan.Get(ctx, t.pool, &tokenModel, query, token); err != nil {
		return &models.Token{}, err
	}

	return tokenModel, nil
}

func (t *TokenPostgres) Delete(ctx context.Context, token string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE refresh_token = $1`, tokensTable)

	if _, err := t.pool.Exec(ctx, query, token); err != nil {
		return err
	}

	return nil
}

func (t *TokenPostgres) UpdateByToken(ctx context.Context, oldToken, newToken string) error {
	query := fmt.Sprintf(`UPDATE %s SET refresh_token = $1 WHERE refresh_token = $2`, tokensTable)

	if _, err := t.pool.Exec(ctx, query, newToken, oldToken); err != nil {
		return err
	}

	return nil
}

func (t *TokenPostgres) UpdateByID(ctx context.Context, id int, token string) error {
	query := fmt.Sprintf(`UPDATE %s SET refresh_token = $1 WHERE id = $2`, tokensTable)

	if _, err := t.pool.Exec(ctx, query, token, id); err != nil {
		return err
	}

	return nil
}
