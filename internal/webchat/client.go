package webchat

import "github.com/gopherzz/webchat/internal/models"

type Client struct {
	poll *Poll
	send chan *models.Message
}

func NewClient(poll *Poll) *Client {
	return &Client{
		poll: poll,
		send: make(chan *models.Message),
	}
}
