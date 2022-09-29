package models

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	tableName struct{}  `pg:"class"`
	ClassID   uuid.UUID `pg:"class_id,pk"`
	ClassName string    `pg:"class_name"`

	SuperClassID uuid.UUID   `pg:"super_class_id"`
	SuperClass   *SuperClass `pg:"rel:has-one"`

	PhylumID uuid.UUID `pg:"phylum_id"`
	Phylum   *Phylum   `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
