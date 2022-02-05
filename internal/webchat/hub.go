package webchat

import "github.com/gopherzz/webchat/internal/services"

type Hub struct {
	services *services.Services

	// All Registered clients
	connections map[*Client]bool

	// Inbound messages from the clients
	broadcast chan []byte

	// Register requests from the clients
	register chan *Client

	// Unregister requests from the clients
	unregister chan *Client
}

func NewHub(services *services.Services) *Hub {
	return &Hub{
		services:    services,
		connections: make(map[*Client]bool),
		broadcast:   make(chan []byte),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.connections[client] = true
		case client := <-h.unregister:
			if _, ok := h.connections[client]; ok {
				delete(h.connections, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.connections {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.connections, client)
				}
			}
		}
	}
}
