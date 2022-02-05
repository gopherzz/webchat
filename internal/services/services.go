package services

import repo "github.com/gopherzz/webchat/internal/repository"

type Messages interface {
	Save(message []byte, clientId string) error
}

type Services struct {
	Messages
}

func NewServices(repo *repo.Repository) *Services {
	return &Services{
		Messages: NewMessages(repo),
	}
}
