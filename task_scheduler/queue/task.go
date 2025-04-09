package queue

import (
	"time"
)

type Task struct {
	ID		string    `json:"id"`
	Payload	string    `json:"payload"`
	CreatedAt	time.Time `json:"created_at"`
	AvailableAt	time.Time `json:"available_at"`
}