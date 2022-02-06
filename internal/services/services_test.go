package services

import (
	"testing"

	"github.com/gopherzz/webchat/internal/repository"
)

func TestServices(t *testing.T) {
	repo := repository.NewMemoryRepository()
	services := NewServices(repo)

	data := []struct {
		message  string
		clientId string
	}{
		{
			message:  "Hello, World",
			clientId: "client-id",
		},
		{
			message:  "Hello, World 2!",
			clientId: "client2-id",
		},
	}

	for _, d := range data {
		if err := services.Save(d.message); err != nil {
			t.Error(err)
		}
	}

	msgs, err := services.GetAll()
	if err != nil {
		t.Error(err)
	}

	if len(msgs) != 2 {
		t.Errorf("Expected 2 messages, got %d", len(msgs))
	}
}
