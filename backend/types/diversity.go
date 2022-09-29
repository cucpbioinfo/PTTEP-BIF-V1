package types

import (
	"github.com/google/uuid"
)

type DiversityCreateBody struct {
	DiversityName string `json:"diversityName"`
}

type DiversityDto struct {
	DiversityID   uuid.UUID `json:"diversityId"`
	Year          string    `json:"year"`
	AssetName     string    `json:"assetName"`
	PlatformName  string    `json:"platformName"`
	StationName   string    `json:"stationName"`
	Surface       string    `json:"surface"`
	Euphotic_zone string    `json:"euphotic_zone"`
}

type ListDiversityQuery struct {
	AssetID    uuid.UUID `json:"assetId"`
	PlatformID uuid.UUID `json:"platformId"`
	StationID  uuid.UUID `json:"astationId"`
}

type ListDiversityResponse struct {
	BaseResponse
	Data []DiversityDto `json:"data"`
}
