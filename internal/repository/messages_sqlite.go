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

func (m *MessagesSqlite) GetAll() ([]*models.Message, error) {
	rows, err := m.db.Query("SELECT * FROM messages")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := make([]*models.Message, 0)
	for rows.Next() {
		var id, senderId string
		var data []byte
		err = rows.Scan(&id, &senderId, &data)
		if err != nil {
			return nil, err
		}

		messages = append(messages, &models.Message{
			Id:       id,
			SenderId: senderId,
			Data:     data,
		})
	}

	return messages, nil
}
