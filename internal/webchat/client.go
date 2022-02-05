package webchat

import (
	"bytes"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	readDeadline = time.Second * 60
	pingPeriod   = time.Second * 10

	writeWait = time.Second * 10

	maxMessageSize = 512
)

var (
	space   = []byte(" ")
	newline = []byte("\n")
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub  *Hub
	id   string
	ws   *websocket.Conn
	send chan []byte
}

func newClient(hub *Hub, ws *websocket.Conn) *Client {
	return &Client{
		hub:  hub,
		id:   uuid.NewString(),
		ws:   ws,
		send: make(chan []byte, 256),
	}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.ws.Close()
	}()

	c.ws.SetReadDeadline(time.Now().Add(readDeadline))
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetPongHandler(func(appData string) error { time.Now().Add(readDeadline); return nil })

	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message = bytes.Trim(bytes.ReplaceAll(message, newline, space), "\x00")
		c.hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.sendMessage(message); err != nil {
				log.Printf("error: %v", err)
				return
			}
		case <-ticker.C:
			// Ping server every 10 seconds
			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("error while ping to server: %v", err)
				return
			}
		}
	}
}

func (c *Client) sendMessage(message []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	w, err := c.ws.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	w.Write(message)

	// Add queued chat messages to the current websocket message, and save to database
	n := len(c.send)
	for i := 0; i < n; i++ {
		w.Write(newline)

		m := <-c.send
		w.Write(m)

		if err := c.hub.services.Messages.Save(m, c.id); err != nil {
			log.Printf("warning: %v", err)
		}
	}

	return w.Close()
}
