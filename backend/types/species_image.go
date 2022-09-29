package types

import (
	"time"

	"github.com/google/uuid"
)

type SpeciesImageDto struct {
	SpeciesImageID  uuid.UUID `json:"speciesImageId"`
	SpeciesID       uuid.UUID `json:"speciesId"`
	FieldLocationID uuid.UUID `json:"fieldLocationId"`
	ImageSource     string    `json:"imageSource"`
	TakenAt         time.Time `json:"taken_at"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
