package repository

import (
	"database/sql"

	"github.com/gopherzz/webchat/internal/models"
)

type Messages interface {
	Save(message *models.Message) error
}

type Repository struct {
	Messages
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Messages: NewMessages(db),
	}
}
