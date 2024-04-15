package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type RespMessage struct {
	UserID      uuid.UUID
	MessageID   uuid.UUID
	UserLogin   string
	Message     string
	WhenCreated time.Time
	WhenUpdate  time.Time
	PathToFiles []string
	Like        int
	DizLike     int
}
