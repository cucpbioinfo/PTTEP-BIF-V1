package types

import (
	"github.com/google/uuid"
	"time"
)

type SiteDto struct {
	SiteID       uuid.UUID        `json:"siteId"`
	SiteName     string           `json:"siteName"`
	MainStations []MainStationDto `json:"mainStations"`
	CreatedAt    time.Time        `json:"createdAt"`
	UpdatedAt    time.Time        `json:"updatedAt"`
}

type SiteListResponse struct {
	BaseResponse
	Data []SiteDto `json:"data"`
}
