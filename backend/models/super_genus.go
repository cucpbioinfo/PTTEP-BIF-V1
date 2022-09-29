package models

import (
	"github.com/google/uuid"
	"time"
)

type SuperGenus struct {
	tableName      struct{}  `pg:"super_genus"`
	SuperGenusID   uuid.UUID `pg:"super_genus_id,pk"`
	SuperGenusName string    `pg:"super_genus_name"`

	SubFamilyID uuid.UUID  `pg:"sub_family_id"`
	SubFamily   *SubFamily `pg:"rel:has-one"`

	FamilyID uuid.UUID `pg:"family_id"`
	Family   *Family   `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
