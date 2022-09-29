package types

import (
	"github.com/google/uuid"
)

type DensityCreateBody struct {
	DensityName string `json:"densityName"`
}

type DensityDto struct {
	DensityID     uuid.UUID `json:"densityId"`
	Year          string    `json:"year"`
	AssetName     string    `json:"assetName"`
	PlatformName  string    `json:"platformName"`
	StationName   string    `json:"stationName"`
	SpeciesName   string    `json:"speciesName"`
	Surface       string    `json:"surface"`
	Euphotic_zone string    `json:"euphotic_zone"`
}

type ListDensityQuery struct {
	AssetID    uuid.UUID `json:"assetId"`
	PlatformID uuid.UUID `json:"platformId"`
	StationID  uuid.UUID `json:"astationId"`
	SpeciesID  uuid.UUID `json:"speciesId"`
}

type ListDensityResponse struct {
	BaseResponse
	Data []DensityDto `json:"data"`
}
