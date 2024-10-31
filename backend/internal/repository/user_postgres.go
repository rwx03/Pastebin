package repository

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v5"
	"github.com/rwx03/Pastebin/backend/internal/models"
)

type UserPostgres struct {
	pool *pgxpool.Pool
}

func NewUserPostgres(pool *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{
		pool: pool,
	}
}

func (u *UserPostgres) Create(ctx context.Context, user models.User) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (email,password,created_at,updated_at) VALUES ($1,$2,$3,$4) RETURNING id`, usersTable)

	if err := u.pool.QueryRow(ctx, query, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserPostgres) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE email = $1`, usersTable)

	if err := pgxscan.Get(ctx, u.pool, &user, query, email); err != nil {
		if err.Error() == fmt.Sprintf("scanning one: %s", pgx.ErrNoRows.Error()) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserPostgres) GetByID(ctx context.Context, id int) (*models.User, error) {
	var user *models.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, usersTable)

	if err := pgxscan.Get(ctx, u.pool, &user, query, id); err != nil {
		return nil, err
	}

	return user, nil
}
