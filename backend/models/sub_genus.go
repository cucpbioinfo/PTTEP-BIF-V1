package models

import (
	"github.com/google/uuid"
	"time"
)

type SubGenus struct {
	tableName    struct{}  `pg:"sub_genus"`
	SubGenusID   uuid.UUID `pg:"sub_genus_id,pk"`
	SubGenusName string    `pg:"sub_genus_name"`

	GenusID uuid.UUID `pg:"genus_id"`
	Genus   *Genus    `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
