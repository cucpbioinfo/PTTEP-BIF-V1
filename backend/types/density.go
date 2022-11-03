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
	StationID     uuid.UUID `json:"stationId"`
	StationName   string    `json:"stationName"`
	SpeciesID     uuid.UUID `json:"speciesId"`
	SpeciesName   string    `json:"speciesName"`
	Surface       string    `json:"surface"`
	Euphotic_zone string    `json:"euphotic_zone"`
}

type ListDensityQuery struct {
	Year       string    `json:"year"`
	AssetID    uuid.UUID `json:"assetId"`
	PlatformID uuid.UUID `json:"platformId"`
	StationID  uuid.UUID `json:"stationId"`
	SpeciesID  uuid.UUID `json:"speciesId"`
}

type ListDensityResponse struct {
	BaseResponse
	Data []DensityDto `json:"data"`
}

// Platform
type PlatformDensityCreateBody struct {
	DensityName string `json:"densityName"`
}

type PlatformDensityDto struct {
	DensityID     uuid.UUID `json:"densityId"`
	Year          string    `json:"year"`
	AssetName     string    `json:"assetName"`
	PlatformID    uuid.UUID `json:"platformId"`
	PlatformName  string    `json:"platformName"`
	SpeciesID     uuid.UUID `json:"speciesId"`
	SpeciesName   string    `json:"speciesName"`
	Surface       string    `json:"surface"`
	Euphotic_zone string    `json:"euphotic_zone"`
}

type ListPlatformDensityQuery struct {
	Year       string    `json:"year"`
	AssetID    uuid.UUID `json:"assetId"`
	PlatformID uuid.UUID `json:"platformId"`
	SpeciesID  uuid.UUID `json:"speciesId"`
}

type ListPlatformDensityResponse struct {
	BaseResponse
	Data []PlatformDensityDto `json:"data"`
}

// Asset
type AssetDensityCreateBody struct {
	DensityName string `json:"densityName"`
}

type AssetDensityDto struct {
	DensityID     uuid.UUID `json:"densityId"`
	Year          string    `json:"year"`
	AssetID       uuid.UUID `json:"assetId"`
	AssetName     string    `json:"assetName"`
	SpeciesID     uuid.UUID `json:"speciesId"`
	SpeciesName   string    `json:"speciesName"`
	Surface       string    `json:"surface"`
	Euphotic_zone string    `json:"euphotic_zone"`
}

type ListAssetDensityQuery struct {
	Year      string    `json:"year"`
	AssetID   uuid.UUID `json:"assetId"`
	SpeciesID uuid.UUID `json:"speciesId"`
}

type ListAssetDensityResponse struct {
	BaseResponse
	Data []AssetDensityDto `json:"data"`
}

// YearAsset
type YearAssetDensityCreateBody struct {
	DensityName string `json:"densityName"`
}

type YearAssetDensityDto struct {
	Year string `json:"year"`
}

type ListYearAssetDensityQuery struct {
	Year      string    `json:"year"`
	AssetID   uuid.UUID `json:"assetId"`
	SpeciesID uuid.UUID `json:"speciesId"`
}

type ListYearAssetDensityResponse struct {
	BaseResponse
	Data []YearAssetDensityDto `json:"data"`
}

// YearPlatform
type YearPlatformDensityCreateBody struct {
	DensityName string `json:"densityName"`
}

type YearPlatformDensityDto struct {
	Year string `json:"year"`
}

type ListYearPlatformDensityQuery struct {
	Year       string    `json:"year"`
	PlatformID uuid.UUID `json:"platformId"`
	AssetID    uuid.UUID `json:"assetId"`
	SpeciesID  uuid.UUID `json:"speciesId"`
}

type ListYearPlatformDensityResponse struct {
	BaseResponse
	Data []YearPlatformDensityDto `json:"data"`
}

// YearStation
type YearDensityCreateBody struct {
	DensityName string `json:"densityName"`
}

type YearDensityDto struct {
	Year string `json:"year"`
}

type ListYearDensityQuery struct {
	Year       string    `json:"year"`
	StationID  uuid.UUID `json:"stationId"`
	PlatformID uuid.UUID `json:"platformId"`
	AssetID    uuid.UUID `json:"assetId"`
	SpeciesID  uuid.UUID `json:"speciesId"`
}

type ListYearDensityResponse struct {
	BaseResponse
	Data []YearDensityDto `json:"data"`
}
