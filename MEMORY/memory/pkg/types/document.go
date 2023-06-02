package types

import (
	"github.com/google/uuid"
	"time"
)

// Document is a struct that represents a document
// in the system
type Document struct {
	ID         uuid.UUID      `json:"id"`
	User       string         `json:"user"`
	Text       string         `json:"text"`
	CreatedAt  time.Time      `json:"created_at"`
	LastReadAt time.Time      `json:"last_read_at"`
	Vector     []float32      `json:"vector"`
	Metadata   map[string]any `json:"metadata"`
}
