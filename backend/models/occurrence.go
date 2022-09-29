package models

import (
	"github.com/google/uuid"
	"time"
)

type Occurrence struct {
	OccurrenceId      uuid.UUID
	ImageSrc          string
	OccurrenceDetails string
	DiscoveredAt      time.Time
}
