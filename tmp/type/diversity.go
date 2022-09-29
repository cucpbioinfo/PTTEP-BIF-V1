package types

import (
	"time"

	"github.com/google/uuid"
)

type DiversityListQuery struct {
	Year       string    `query:"year"`
	ProjectID  uuid.UUID `query:"projectId"`
	PlatformID uuid.UUID `query:"platformId"`
	StationID  uuid.UUID `query:"stationId"`
}

type DiversityDto struct {
	DiversityID  uuid.UUID `json:"diversityId"`
	Diversity    string    `json:"diversity"`
	Year         string    `json:"year"`
	ProjectName  string    `json:"projectName"`
	PlatformName string    `json:"PlatformName"`
	StationName  string    `json:"StationName"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DiversityListResponse struct {
	BaseResponse
	Data []DiversityDto `json:"data"`
}
