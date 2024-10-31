package services

import (
	"context"

	"github.com/rwx03/Pastebin/backend/internal/models"
	"github.com/rwx03/Pastebin/backend/internal/repository"
)

type PasteService struct {
	repo repository.Paste
}

func NewPasteService(repo repository.Paste) *PasteService {
	return &PasteService{repo: repo}
}

func (p *PasteService) Create(paste models.Paste) (int, error) {
	ctx := context.Background()
	return p.repo.Create(ctx, paste)
}

func (p *PasteService) GetPasteByID(id string) (*models.Paste, error) {
	ctx := context.Background()
	return p.repo.GetPasteByID(ctx, id)
}

func (p *PasteService) GetAllPastes() ([]models.Paste, error) {
	ctx := context.Background()
	return p.repo.GetAllPastes(ctx)
}

func (p *PasteService) GetAllPastesByUser(userID int) ([]models.Paste, error) {
	ctx := context.Background()
	return p.repo.GetAllPastesByUser(ctx, userID)
}
