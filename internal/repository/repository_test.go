package repository

import (
	"testing"

	"github.com/gopherzz/webchat/internal/models"
)

func TestRepository_Save(t *testing.T) {
	repo := NewMemoryRepository()

	data := []*models.Message{
		{
			Id:       "1",
			SenderId: "1",
			Data:     []byte("Hello, World!"),
		},
		{
			Id:       "2",
			SenderId: "2",
			Data:     []byte("Hello, World 2!"),
		},
	}

	for _, m := range data {
		if err := repo.Save(m); err != nil {
			t.Error(err)
		}
	}

	l, err := repo.GetAll()
	if err != nil {
		t.Error(err)
	}

	if len(l) != 2 {
		t.Errorf("Expected 2 messages, got %d", len(l))
	}

	for i, m := range l {
		if m.SenderId != data[i].SenderId {
			t.Errorf("Expected %s, got %s", data[i].SenderId, m.SenderId)
		}
	}
}

func TestRepository_GetAll(t *testing.T) {
	repo := NewMemoryRepository()

	data := []*models.Message{
		{
			Id:       "1",
			SenderId: "1",
			Data:     []byte("Hello, World!"),
		},
		{
			Id:       "2",
			SenderId: "2",
			Data:     []byte("Hello, World 2!"),
		},
	}

	msgs, err := repo.GetAll()
	if err != nil {
		t.Error(err)
	}

	if len(msgs) != 0 {
		t.Errorf("Expected 0 messages, got %d", len(msgs))
	}

	for _, m := range data {
		if err := repo.Save(m); err != nil {
			t.Error(err)
		}
	}

	msgs, err = repo.GetAll()
	if err != nil {
		t.Error(err)
	}

	for i, m := range msgs {
		if m.SenderId != data[i].SenderId {
			t.Errorf("Expected %s, got %s", data[i].SenderId, m.SenderId)
		}
	}
}
