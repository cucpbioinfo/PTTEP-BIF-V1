package models

import (
	"github.com/google/uuid"
	"time"
)

type SubFamily struct {
	tableName     struct{}  `pg:"sub_family"`
	SubFamilyID   uuid.UUID `pg:"sub_family_id,pk"`
	SubFamilyName string    `pg:"sub_family_name"`

	FamilyID uuid.UUID `pg:"family_id"`
	Family   *Family   `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
