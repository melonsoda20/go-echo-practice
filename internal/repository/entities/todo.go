package entities

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	id           uuid.UUID
	name         string
	age          int
	created_by   string
	created_date time.Time
	updated_by   string
	updated_date time.Time
}
