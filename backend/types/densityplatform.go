package types

import (
	"github.com/google/uuid"
)

type PlatformDensityCreateBody struct {
	DensityName string `json:"densityName"`
}

type PlatformDensityDto struct {
	DensityID     uuid.UUID `json:"densityId"`
	Year          string    `json:"year"`
	AssetName     string    `json:"assetName"`
	PlatformName  string    `json:"platformName"`
	SpeciesID     uuid.UUID `json:"speciesId"`
	SpeciesName   string    `json:"speciesName"`
	Surface       string    `json:"surface"`
	Euphotic_zone string    `json:"euphotic_zone"`
}

type ListPlatformDensityQuery struct {
	AssetID    uuid.UUID `json:"assetId"`
	PlatformID uuid.UUID `json:"platformId"`
	SpeciesID  uuid.UUID `json:"speciesId"`
}

type ListPlatformDensityResponse struct {
	BaseResponse
	Data []DensityDto `json:"data"`
}
