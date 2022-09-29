package models

import (
	"github.com/google/uuid"
	"time"
)

type SuperPhylum struct {
	tableName       struct{}  `pg:"super_phylum"`
	SuperPhylumID   uuid.UUID `pg:"super_phylum_id,pk"`
	SuperPhylumName string    `pg:"super_phylum_name"`

	KingdomID uuid.UUID `pg:"kingdom_id"`
	Kingdom   *Kingdom  `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
