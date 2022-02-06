package repository

import (
	"errors"
	"sync"

	"github.com/gopherzz/cucumberdb"
	"github.com/gopherzz/webchat/internal/models"
)

type MessagesJdb struct {
	db *cucumberdb.Db
}

func NewMessages(db *cucumberdb.Db) *MessagesJdb {
	(&sync.Once{}).Do(func() { db.LCreate("messages"); db.Dump() })
	return &MessagesJdb{db: db}
}

func (m *MessagesJdb) Save(message *models.Message) error {
	if ok := m.db.LAdd("messages", message); !ok {
		return errors.New("failed to add message to list")
	}
	return m.db.Dump()
}

func (m *MessagesJdb) GetAll() ([]*models.Message, error) {
	msgs := m.db.LGetAll("messages")

	messages := make([]*models.Message, 0)
	for _, msg := range msgs {
		message := msg.(*models.Message)
		messages = append(messages, message)
	}

	return messages, nil
}
