package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rwx03/Pastebin/backend/internal/models"
)

type User interface {
	Create(ctx context.Context, user models.User) (int, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByID(ctx context.Context, id int) (*models.User, error)
}

type Paste interface {
	Create(ctx context.Context, paste models.Paste) (int, error)
	GetPasteByID(ctx context.Context, id string) (*models.Paste, error)
	GetAllPastes(ctx context.Context) ([]models.Paste, error)
	GetAllPastesByUser(ctx context.Context, userID int) ([]models.Paste, error)
}

type Token interface {
	Create(ctx context.Context, token models.Token) (int, error)
	GetToken(ctx context.Context, token string) (*models.Token, error)
	Delete(ctx context.Context, token string) error
	UpdateByToken(ctx context.Context, oldToken, newToken string) error
	UpdateByID(ctx context.Context, id int, token string) error
}

type Repository struct {
	User
	Token
	Paste
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Paste: NewPastePostgres(pool),
		Token: NewTokenPostgres(pool),
		User:  NewUserPostgres(pool),
	}
}
