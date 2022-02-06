package webchat

import (
	"log"

	"github.com/gopherzz/webchat/internal/models"
	"github.com/gopherzz/webchat/internal/services"
)

type Poll struct {
	services *services.Services

	connected  map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *models.Message
}

func NewPoll(services *services.Services) *Poll {
	return &Poll{
		services:   services,
		connected:  make(map[*Client]bool),
		broadcast:  make(chan *models.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (p *Poll) Run() {
	for {
		select {
		case msg := <-p.broadcast:
			err := p.services.Save(msg.Body)
			if err != nil {
				log.Println("error saving message:", err)
			}

			for c := range p.connected {
				c.send <- msg
			}
		case c := <-p.register:
			p.connected[c] = true
		case c := <-p.unregister:
			delete(p.connected, c)
		}
	}
}
