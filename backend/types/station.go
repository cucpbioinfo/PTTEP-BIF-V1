package types

import (
	"github.com/google/uuid"
)

type StationCreateBody struct {
	StationName string `json:"stationName"`
}

type StationDto struct {
	StationID   uuid.UUID `json:"stationId"`
	StationName string    `json:"stationName"`
}

type ListStationQuery struct {
	PlatformID uuid.UUID `json:"platformId"`
}

type ListStationResponse struct {
	BaseResponse
	Data []StationDto `json:"data"`
}
