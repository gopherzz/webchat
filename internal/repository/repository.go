package repository

import (
	"github.com/gopherzz/cucumberdb"
	"github.com/gopherzz/webchat/internal/models"
)

type Messages interface {
	Save(message *models.Message) error
	GetAll() ([]*models.Message, error)
}

type Repository struct {
	Messages
}

func NewRepository(db *cucumberdb.Db) *Repository {
	return &Repository{
		Messages: NewMessages(db),
	}
}
