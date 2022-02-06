package repository

import (
	"github.com/gopherzz/webchat/internal/models"
)

type MemoryMessages struct {
	db []*models.Message
}

func NewMemoryMessages() *MemoryMessages {
	return &MemoryMessages{
		db: make([]*models.Message, 0),
	}
}

func (m *MemoryMessages) Save(message *models.Message) error {
	m.db = append(m.db, message)
	return nil
}

func (m *MemoryMessages) GetAll() ([]*models.Message, error) {
	return m.db, nil
}
