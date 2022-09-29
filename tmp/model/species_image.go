package models

import (
	"time"

	"github.com/google/uuid"
)

type SpeciesImage struct {
	tableName      struct{}  `pg:"species_image"`
	SpeciesImageID uuid.UUID `pg:"species_image_id"`

	SpeciesID uuid.UUID `pg:"species_id"`
	Species   *Species  `pg:"rel:has-one"`

	FieldLocationID uuid.UUID      `pg:"field_location_id"`
	FieldLocation   *FieldLocation `pg:"rel:has-one"`

	ImageSource string    `pg:"image_source"`
	TakenAt     time.Time `pg:"taken_at"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
