package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       uuid.UUID `json:"id"`
	CreateAt time.Time `json:"created_at"`
	Body     string    `json:"body"`
}
