package repository

import (
	"github.com/gopherzz/webchat/internal/models"
)

type MemoryMessages struct {
	db map[string]*models.Message
}

func NewMemoryMessages() *MemoryMessages {
	return &MemoryMessages{
		db: make(map[string]*models.Message),
	}
}

func (m *MemoryMessages) Save(message *models.Message) error {
	m.db[message.Id] = message
	return nil
}

func (m *MemoryMessages) GetAll() ([]*models.Message, error) {
	messages := make([]*models.Message, 0)
	for _, message := range m.db {
		messages = append(messages, message)
	}

	return messages, nil
}
