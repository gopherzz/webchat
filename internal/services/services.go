package services

import (
	"github.com/gopherzz/webchat/internal/models"
	repo "github.com/gopherzz/webchat/internal/repository"
)

type Messages interface {
	Save(message string) error
	GetAll() ([]*models.Message, error)
}

type Services struct {
	Messages
}

func NewServices(repo *repo.Repository) *Services {
	return &Services{
		Messages: NewMessages(repo),
	}
}
