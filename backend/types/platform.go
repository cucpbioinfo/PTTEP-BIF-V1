package types

import (
	"github.com/google/uuid"
)

type PlatformCreateBody struct {
	PlatformName string `json:"platformName"`
}

type PlatformDto struct {
	PlatformID   uuid.UUID `json:"platformId"`
	PlatformName string    `json:"platformName"`
	Latitude     string    `json:"latitude"`
	Longitude    string    `json:"longitude"`
	Type         string    `json:"type"`
}

type ListPlatformQuery struct {
	AssetID uuid.UUID `json:"assetId"`
}

type ListPlatformResponse struct {
	BaseResponse
	Data []PlatformDto `json:"data"`
}
