package types

import (
	"github.com/google/uuid"
)

type PlatformDto struct {
	PlatformID   uuid.UUID `json:"platformId"`
	PlatformName string    `json:"platformName"`
	ProjectID    uuid.UUID `json:"projectId"`
	ProjectName  string    `json:"projectName"`
	Latitude     string    `json:"Latitude"`
	Longitude    string    `json:"Longitude"`
}

type ListPlatformResponse struct {
	BaseResponse
	Data []PlatformDto `json:"data"`
}
