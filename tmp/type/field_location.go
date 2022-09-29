package types

import (
	"time"

	"github.com/google/uuid"
)

type FieldLocationListQuery struct {
	MainStationID uuid.UUID `json:"mainStationId"`
}

type FieldLocationDto struct {
	FieldLocationID   uuid.UUID `json:"fieldLocationId"`
	FieldLocationName string    `json:"fieldLocationName"`
	Location          *Location `json:"location"`
	Replication       int32     `json:"replication"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type FieldLocationListResponse struct {
	BaseResponse
	Data []FieldLocationDto `json:"data"`
}
