package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type Topic struct {
	ID          uuid.UUID
	WhenCreated time.Time
	WhenUpdate  time.Time
	WhenDeleted time.Time
	TopicInfo   string
	TopicName   string
	TopicFile   string
}
