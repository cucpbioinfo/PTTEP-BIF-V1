package models

import (
	"github.com/google/uuid"
	"time"
)

type SpeciesFieldLocation struct {
	tableName struct{} `pg:"species_field_location"`

	SpeciesID uuid.UUID `pg:"species_id"`
	Species   *Species  `pg:"rel:has-one"`

	FieldLocationID uuid.UUID      `pg:"field_location_id"`
	FieldLocation   *FieldLocation `pg:"rel:has-one"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
