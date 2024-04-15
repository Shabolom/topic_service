package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type UsersTopics struct {
	ID          uuid.UUID
	WhenCreated time.Time
	WhenDeleted time.Time
	UserID      uuid.UUID
	TopicID     uuid.UUID
}
