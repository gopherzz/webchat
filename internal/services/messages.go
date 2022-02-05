package services

import (
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

func (ms *MessagesService) Save(message []byte, clientId string) error {
	msg := createMessage(message, clientId)
	return ms.repo.Save(msg)
}

func createMessage(message []byte, clientId string) *models.Message {
	return &models.Message{
		Id:       uuid.New().String(),
		SenderId: clientId,
		Data:     message,
	}
}
