package models

import (
	"time"

	"github.com/google/uuid"
)

type Identification struct {
	tableName          struct{}  `pg:"identification"`
	IdentificationID   uuid.UUID `pg:"identification_id,pk"`
	IdentificationName string    `pg:"identification_name"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
