package models

import (
	"time"

	"github.com/google/uuid"
)

type SubKingdom struct {
	tableName      struct{}  `pg:"sub_kingdom"`
	SubKingdomID   uuid.UUID `pg:"sub_kingdom_id"`
	SubKingdomName string    `pg:"sub_kingdom_name"`

	KingdomID uuid.UUID `pg:"kingdom_id"`
	// Kingdom   *Kingdom  `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
