package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/gopherzz/webchat/internal/models"
	repo "github.com/gopherzz/webchat/internal/repository"
)

type MessagesService struct {
	repo *repo.Repository
}

func NewMessages(repo *repo.Repository) *MessagesService {
	return &MessagesService{
		repo: repo,
	}
}

func (ms *MessagesService) Save(message string) error {
	msg := createMessage(message)
	return ms.repo.Save(msg)
}

func createMessage(message string) *models.Message {
	return &models.Message{
		ID:       uuid.New(),
		CreateAt: time.Now(),
		Body:     message,
	}
}

func (ms *MessagesService) GetAll() ([]*models.Message, error) {
	return ms.repo.GetAll()
}
