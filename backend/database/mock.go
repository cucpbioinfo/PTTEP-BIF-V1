package database

import (
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"time"
)

// For mock data
type MockDB struct {
	OccurrenceData []models.Occurrence
}

func newOccurrenceData() []models.Occurrence {

	occurrences := []models.Occurrence{
		models.Occurrence{
			OccurrenceId:      uuid.MustParse("b36b99e6-f400-46c6-ab98-639792e29d2c"),
			ImageSrc:          "/blue-whale.jpeg",
			OccurrenceDetails: "The blue whale (Balaenoptera musculus) is a marine mammal. Reaching a maximum confirmed length of 29.9 metres and weighing up to 199 tonnes, it is the largest animal known to have ever existed.",
			DiscoveredAt:      time.Now(),
		},
		models.Occurrence{
			OccurrenceId:      uuid.MustParse("e26beac4-f29c-418b-8281-e751a3fc1992"),
			ImageSrc:          "/chaetoceros.jpeg",
			OccurrenceDetails: "Chaetoceros is probably the largest genus of marine planktonic diatoms with approximately 400 species described, although many of these descriptions are no longer valid. It is often very difficult to distinguish between different Chaetoceros species.",
			DiscoveredAt:      time.Now(),
		},
	}

	return occurrences
}

func NewMockDB() *MockDB {
	return &MockDB{
		OccurrenceData: newOccurrenceData(),
	}
}
