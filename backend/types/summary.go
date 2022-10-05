package types

import (
	"github.com/google/uuid"
)

type SummaryCreateBody struct {
	SummaryYear string `json:"summaryYear"`
}

type SummaryDto struct {
	SummaryID            uuid.UUID `json:"summaryId"`
	Year                 string    `json:"year"`
	MajorGroupName       string    `json:"majorGroupName"`
	IdentificationName   string    `json:"identificationName"`
	AssetName            string    `json:"assetName"`
	PlatformName         string    `json:"platformName"`
	StationName          string    `json:"stationName"`
	SurfaceShannon       string    `json:"surfaceShannon"`
	SurfaceNumber        string    `json:"surfaceNumber"`
	SurfaceMax           string    `json:"surfaceMax"`
	SurfaceEvenness      string    `json:"surfaceEvenness"`
	EuphoticZoneShannon  string    `json:"euphoticzoneShannon"`
	EuphoticZoneNumber   string    `json:"euphoticzoneNumber"`
	EuphoticZoneMax      string    `json:"euphoticzoneMax"`
	EuphoticZoneEvenness string    `json:"euphoticzoneEvenness"`
	AverageShannon       string    `json:"averageShannon"`
	AverageNumber        string    `json:"averageNumber"`
	AverageMax           string    `json:"averageMax"`
	AverageEvenness      string    `json:"averageEvenness"`
}

type ListSummaryQuery struct {
	MajorGroupID     uuid.UUID `json:"majorGroupId"`
	IdentificationID uuid.UUID `json:"identificationId"`
	AssetID          uuid.UUID `json:"assetId"`
	PlatformID       uuid.UUID `json:"platformId"`
	StationID        uuid.UUID `json:"astationId"`
}

type ListSummaryResponse struct {
	BaseResponse
	Data []SummaryDto `json:"data"`
}
