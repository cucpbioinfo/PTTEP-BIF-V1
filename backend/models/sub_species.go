package models

import (
	"github.com/google/uuid"
	"time"
)

type SubSpecies struct {
	tableName      struct{}  `pg:"sub_species"`
	SubSpeciesID   uuid.UUID `pg:"sub_species_id,pk"`
	SubSpeciesName string    `pg:"sub_species_name"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
