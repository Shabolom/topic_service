package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type Message struct {
	ID           uuid.UUID
	WhenCreated  time.Time
	WhenUpdated  time.Time
	WhenDeleted  time.Time
	UserMessage  string
	UserFilePath string
	UserID       uuid.UUID
	TopicID      uuid.UUID
}
