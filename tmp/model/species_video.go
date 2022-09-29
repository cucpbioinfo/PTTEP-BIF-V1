package models

import (
	"time"

	"github.com/google/uuid"
)

type SpeciesVideo struct {
	tableName      struct{}  `pg:"species_video"`
	SpeciesVideoID uuid.UUID `pg:"species_video_id"`

	SpeciesID uuid.UUID `pg:"species_id"`
	Species   *Species  `pg:"rel:has-one"`

	FieldLocationID uuid.UUID      `pg:"field_location_id"`
	FieldLocation   *FieldLocation `pg:"rel:has-one"`

	VideoSource string    `pg:"video_source"`
	TakenAt     time.Time `pg:"taken_at"`

	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt time.Time `pg:"deleted_at"`
}
