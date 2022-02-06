package webchat

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gopherzz/webchat/internal/models"
)

func (poll *Poll) GetMessages(w http.ResponseWriter, r *http.Request) {
	timeout := time.NewTimer(10 * time.Second)

	c := NewClient(poll)
	poll.register <- c

	defer func() {
		timeout.Stop()
		poll.unregister <- c
	}()

	select {
	case msg := <-c.send:

		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			log.Printf("error encoding message: %v", err)
		}
	case <-timeout.C:
		w.WriteHeader(http.StatusGatewayTimeout)
	}
}

func (poll *Poll) SendMessage(w http.ResponseWriter, r *http.Request) {
	msg := &models.Message{}
	json.NewDecoder(r.Body).Decode(msg)

	poll.broadcast <- msg
}
