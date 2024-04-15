package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type Rating struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time
	MessageID uuid.UUID
	UserID    uuid.UUID
}
