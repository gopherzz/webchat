package webchat

import "net/http"

func WebSocketHandler(hub *Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		client := newClient(hub, conn)
		client.hub.register <- client

		go client.writePump()
		go client.readPump()
	}
}
