package types

import (
	"github.com/google/uuid"
)

type LocationCreateBody struct {
	LocationName string `json:"locationName"`
}

type LocationDto struct {
	LocationID   uuid.UUID `json:"locationId"`
	AssetName    string    `json:"assetName"`
	PlatformName string    `json:"platformName"`
	StationName  string    `json:"stationName"`
	Latitude     string    `json:"latitude"`
	Longitude    string    `json:"longitude"`
	Type         string    `json:"type"`
}

type LocationAssetDto struct {
	LocationID uuid.UUID `json:"locationId"`
	AssetID    uuid.UUID `json:"assetId"`
	AssetName  string    `json:"assetName"`
	Latitude   string    `json:"latitude"`
	Longitude  string    `json:"longitude"`
	Type       string    `json:"type"`
}

type LocationPlatformDto struct {
	LocationID   uuid.UUID `json:"locationId"`
	AssetName    string    `json:"assetName"`
	PlatformName string    `json:"platformName"`
	Latitude     string    `json:"latitude"`
	Longitude    string    `json:"longitude"`
	Type         string    `json:"type"`
}

type ListLocationQuery struct {
	AssetID    uuid.UUID `json:"assetId"`
	PlatformID uuid.UUID `json:"platformId"`
	StationID  uuid.UUID `json:"astationId"`
}

type ListLocationAssetQuery struct {
	AssetID uuid.UUID `json:"assetId"`
}

type ListLocationResponse struct {
	BaseResponse
	Data []LocationDto `json:"data"`
}

type ListLocationAssetResponse struct {
	BaseResponse
	Data []LocationAssetDto `json:"data"`
}

type ListLocationPlatformResponse struct {
	BaseResponse
	Data []LocationPlatformDto `json:"data"`
}
