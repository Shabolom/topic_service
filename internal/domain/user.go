package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID
	WhenCreated time.Time
	WhenUpdate  time.Time
	WhenDeleted time.Time
	UserName    string
}
