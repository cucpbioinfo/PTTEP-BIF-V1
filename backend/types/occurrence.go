package types

import (
	"github.com/google/uuid"
	"time"
)

type OccurrenceDto struct {
	OccurrenceId      uuid.UUID `json:"occurrenceId"`
	ImageSrc          string    `json:"imageSrc"`
	OccurrenceDetails string    `json:"occurrenceDetails"`
	DiscoveredAt      time.Time `json:"discoveredAt"`
}

type ListOccurrenceResponse struct {
	BaseResponse
	Data  []OccurrenceDto `json:"data"`
	Total int             `json:"total"`
}
