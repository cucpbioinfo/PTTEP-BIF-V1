package types

import (
	"time"

	"github.com/google/uuid"
)

type SpeciesVideoDto struct {
	SpeciesVideoID  uuid.UUID `json:"speciesVideoId"`
	SpeciesID       uuid.UUID `json:"speciesId"`
	FieldLocationID uuid.UUID `json:"fieldLocationId"`
	VideoSource     string    `json:"videoSource"`
	TakenAt         time.Time `json:"takenAt"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
