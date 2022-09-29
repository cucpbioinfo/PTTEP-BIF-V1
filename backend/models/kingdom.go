package models

import (
	"github.com/google/uuid"
	"time"
)

type Kingdom struct {
	tableName   struct{}  `pg:"kingdom"`
	KingdomID   uuid.UUID `pg:"kingdom_id,pk"`
	KingdomName string    `pg:"kingdom_name"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
