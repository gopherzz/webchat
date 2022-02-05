package repository

import (
	"database/sql"

	"github.com/gopherzz/webchat/internal/models"
)

type MessagesSqlite struct {
	db *sql.DB
}

func NewMessages(db *sql.DB) *MessagesSqlite {
	return &MessagesSqlite{db: db}
}

func (m *MessagesSqlite) Save(message *models.Message) error {
	_, err := m.db.Exec("INSERT INTO messages (id, sender_id, data) VALUES ($1, $2, $3)", message.Id, message.SenderId, message.Data)
	if err != nil {
		return err
	}

	return nil
}
