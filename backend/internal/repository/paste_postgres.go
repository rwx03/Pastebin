package repository

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rwx03/Pastebin/backend/internal/models"
)

type PastePostgres struct {
	pool *pgxpool.Pool
}

func NewPastePostgres(pool *pgxpool.Pool) *PastePostgres {
	return &PastePostgres{
		pool: pool,
	}
}

func (p *PastePostgres) Create(ctx context.Context, paste models.Paste) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (paste_id,title,content,creator_id,created_at,views) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`, pastesTable)
	if err := p.pool.QueryRow(ctx, query, paste.PasteID, paste.Title, paste.Content, paste.CreatorID, paste.CreatedAt, paste.Views).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PastePostgres) GetPasteByID(ctx context.Context, id string) (*models.Paste, error) {
	var paste models.Paste

	query := fmt.Sprintf(`SELECT * FROM %s WHERE paste_id = $1`, pastesTable)
	if err := pgxscan.Get(ctx, p.pool, &paste, query, id); err != nil {
		return nil, err
	}

	return &paste, nil
}

func (p *PastePostgres) GetAllPastes(ctx context.Context) ([]models.Paste, error) {
	var pastes []models.Paste

	query := fmt.Sprintf(`SELECT * FROM %s`, pastesTable)
	if err := pgxscan.Select(ctx, p.pool, &pastes, query); err != nil {
		return nil, err
	}

	return pastes, nil
}

func (p *PastePostgres) GetAllPastesByUser(ctx context.Context, userID int) ([]models.Paste, error) {
	var pastes []models.Paste

	query := fmt.Sprintf(`SELECT * FROM %s WHERE creator_id = $1`, pastesTable)
	if err := pgxscan.Select(ctx, p.pool, &pastes, query, userID); err != nil {
		return nil, err
	}

	return pastes, nil
}
