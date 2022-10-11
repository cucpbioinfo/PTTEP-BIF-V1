package types

import (
	"github.com/google/uuid"
)

type AssetDensityCreateBody struct {
	DensityName string `json:"densityName"`
}

type AssetDensityDto struct {
	DensityID     uuid.UUID `json:"densityId"`
	Year          string    `json:"year"`
	AssetName     string    `json:"assetName"`
	SpeciesID     uuid.UUID `json:"speciesId"`
	SpeciesName   string    `json:"speciesName"`
	Surface       string    `json:"surface"`
	Euphotic_zone string    `json:"euphotic_zone"`
}

type ListAssetDensityQuery struct {
	AssetID   uuid.UUID `json:"assetId"`
	SpeciesID uuid.UUID `json:"speciesId"`
}

type ListAssetDensityResponse struct {
	BaseResponse
	Data []DensityDto `json:"data"`
}
