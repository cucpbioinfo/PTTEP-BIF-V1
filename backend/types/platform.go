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
}

type ListPlatformQuery struct {
	AssetID uuid.UUID `json:"assetId"`
}

type ListPlatformResponse struct {
	BaseResponse
	Data []PlatformDto `json:"data"`
}
