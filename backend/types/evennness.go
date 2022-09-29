package types

import (
	"github.com/google/uuid"
)

type EvennessCreateBody struct {
	EvennessName string `json:"evennessName"`
}

type EvennessDto struct {
	EvennessID    uuid.UUID `json:"evennessId"`
	Year          string    `json:"year"`
	AssetName     string    `json:"assetName"`
	PlatformName  string    `json:"platformName"`
	StationName   string    `json:"stationName"`
	Surface       string    `json:"surface"`
	Euphotic_zone string    `json:"euphotic_zone"`
}

type ListEvennessQuery struct {
	AssetID    uuid.UUID `json:"assetId"`
	PlatformID uuid.UUID `json:"platformId"`
	StationID  uuid.UUID `json:"astationId"`
}

type ListEvennessResponse struct {
	BaseResponse
	Data []EvennessDto `json:"data"`
}
