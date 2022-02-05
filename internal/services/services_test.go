package services

import (
	"testing"

	"github.com/gopherzz/webchat/internal/repository"
)

func TestServices(t *testing.T) {
	repo := repository.NewMemoryRepository()
	services := NewServices(repo)

	data := []struct {
		message  []byte
		clientId string
	}{
		{
			message:  []byte("Hello, World!"),
			clientId: "client-id",
		},
		{
			message:  []byte("Hello, World 2!"),
			clientId: "client2-id",
		},
	}

	for _, d := range data {
		if err := services.Save(d.message, d.clientId); err != nil {
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

	for i, m := range msgs {
		if m.SenderId != data[i].clientId {
			t.Errorf("Expected %s, got %s", data[i].clientId, m.SenderId)
		}
	}
}
