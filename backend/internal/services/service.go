package services

import (
	"github.com/rwx03/Pastebin/backend/internal/models"
	"github.com/rwx03/Pastebin/backend/internal/repository"
)

type Paste interface {
	Create(paste models.Paste) (int, error)
	GetPasteByID(id string) (*models.Paste, error)
	GetAllPastes() ([]models.Paste, error)
	GetAllPastesByUser(userID int) ([]models.Paste, error)
}

type Auth interface {
	Register(email, password string) (string, string, error)
	Login(email, password string) (string, string, error)
	Refresh(refreshToken string) (string, string, error)
}

type Service struct {
	Paste
	Auth
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Paste: NewPasteService(repo.Paste),
		Auth:  NewAuthService(repo.User, repo.Token),
	}
}
