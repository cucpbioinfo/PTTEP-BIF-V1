package models

import (
	"github.com/google/uuid"
	"time"
)

type Genus struct {
	tableName struct{}  `pg:"genus"`
	GenusID   uuid.UUID `pg:"genus_id,pk"`
	GenusName string    `pg:"genus_name"`

	SuperGenusID uuid.UUID   `pg:"super_genus_id"`
	SuperGenus   *SuperGenus `pg:"rel:has-one"`

	FamilyID uuid.UUID `pg:"family_id"`
	Family   *Family   `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
