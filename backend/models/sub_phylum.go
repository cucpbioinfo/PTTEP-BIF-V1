package models

import (
	"github.com/google/uuid"
	"time"
)

type SubPhylum struct {
	tableName     struct{}  `pg:"sub_phylum"`
	SubPhylumID   uuid.UUID `pg:"sub_phylum_id,pk"`
	SubPhylumName string    `pg:"sub_phylum_name"`

	PhylumID uuid.UUID `pg:"phylum_id"`
	Phylum   *Phylum   `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
