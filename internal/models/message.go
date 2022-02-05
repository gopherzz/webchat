package models

type Message struct {
	// Id is uuid string
	Id string `db:"id"`

	// SenderId is uuid string
	SenderId string `db:"sender_id"`

	// Data is message data
	Data []byte `db:"data"`
}
