package types

import (
	"time"

	"github.com/google/uuid"
)

type MainStationListQuery struct {
	SiteID uuid.UUID `json:"siteId"`
}

type MainStationDto struct {
	MainStationID   uuid.UUID          `json:"mainStationId"`
	MainStationName string             `json:"mainStationName"`
	RadiusFromSite  string             `json:"radiusFromSite"`
	Easting         int32              `json:"easting"`
	Northing        int32              `json:"northing"`
	Location        *Location          `json:"location"`
	FieldLocations  []FieldLocationDto `json:"fieldLocations"`
	CreatedAt       time.Time          `json:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt"`
}

type MainStationListResponse struct {
	BaseResponse
	Data []MainStationDto `json:"data"`
}
