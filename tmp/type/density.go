package types

import (
	"time"

	"github.com/google/uuid"
)

// "density_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
// 		"year" varchar(500) NOT NULL,
// 		"species_id" UUID,
// 		"species_name" varchar(500),
// 		"asset_id" UUID NOT NULL,
// 		"platform_id" UUID NOT NULL,
// 		"station_id" UUID NOT NULL,
// 		"surface" float NOT NULL,
// 		"euphotic_zone" float NOT NULL,
type DensityListQuery struct {
	Keyword    string    `query:"keyword"`
	AssetID    uuid.UUID `query:"assetId"`
	PlatformID uuid.UUID `query:"platformId"`
	StationID  uuid.UUID `query:"stationId"`
	// IdentificationID uuid.UUID `query:"identificationId"`
	// MethodID         uuid.UUID `query:"methodId"`
	SpeciesID  uuid.UUID `query:"SpeciesId"`
	PageNumber int       `query:"pageNumber"`
	PageSize   int       `query:"pageSize"`
}
type DensityDto struct {
	DensityID uuid.UUID `json:"densityId"`
	Year      string    `json:"density_year"`
	Project   string    `json:"project"`
	Platform  string    `json:"platform"`
	Station   string    `json:"station"`
	// Identification string    `json:"identification"`
	// Method         string    `json:"method"`
	Species   string    `json:"species"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DensityListMeta struct {
	DensityListQuery   `json:"query"`
	PaginationResponse `json:"pagination"`
}

type ListDensityResponse struct {
	BaseResponse
	Data []DensityDto       `json:"data"`
	Meta PaginationResponse `json:"meta"`
}

type DensityDetailsDto struct {
	DensityID          uuid.UUID `json:"densityId"`
	ProjectName        string    `json:"projectName"`
	PlatformName       string    `json:"platformName"`
	StationName        string    `json:"stationName"`
	IdentificationName string    `json:"identificationName"`
	MethodName         string    `json:"methodName"`
	// MajorGroupName     string    `json:"majorGroupName"`
	// KingdomName        string    `json:"kingdomName"`
	// PhylumName         string    `json:"phylumName"`
	// ClassName          string    `json:"className"`
	// OrderName          string    `json:"orderName"`
	// FamilyName         string    `json:"familyName"`
	// GenusName          string    `json:"genusName"`
	SpeciesName string    `json:"speciesName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type DensityDetailsResponse struct {
	BaseResponse
	Data DensityDetailsDto `json:"data"`
}

type DensitySearchQuery struct {
	Keyword string `json:"keyword"`
	Limit   int    `json:"limit"`
}

type DensitySearchResponse struct {
	BaseResponse
	Data []DensityDto `json:"data"`
}

type DensityCountGroupByDensityType struct {
	DensityType string `json:"densityType"`
	Count       int    `json:"count"`
}

type DensityCountGroupByDensityTypeResponse struct {
	BaseResponse
	Data []DensityCountGroupByDensityType `json:"data"`
}
