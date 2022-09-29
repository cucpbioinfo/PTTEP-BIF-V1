package models

import (
	"github.com/google/uuid"
	"time"
)

type Phylum struct {
	tableName  struct{}  `pg:"phylum"`
	PhylumID   uuid.UUID `pg:"phylum_id,pk"`
	PhylumName string    `pg:"phylum_name"`

	SuperPhylumID uuid.UUID    `pg:"super_phylum_id"`
	SuperPhylum   *SuperPhylum `pg:"rel:has-one"`

	KingdomID uuid.UUID `pg:"kingdom_id"`
	Kingdom   *Kingdom  `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
