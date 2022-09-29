package models

import (
	"github.com/google/uuid"
	"time"
)

type SuperClass struct {
	tableName      struct{}  `pg:"super_class"`
	SuperClassID   uuid.UUID `pg:"super_class_id,pk"`
	SuperClassName string    `pg:"super_class_name"`

	SubPhylumID uuid.UUID  `pg:"sub_phylum_id"`
	SubPhylum   *SubPhylum `pg:"rel:has-one"`

	PhylumID uuid.UUID `pg:"phylum_id"`
	Phylum   *Phylum   `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
