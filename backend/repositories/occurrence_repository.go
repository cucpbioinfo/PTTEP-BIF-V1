package repositories

import (
	"github.com/wichadak/eDNA/database"
	"github.com/wichadak/eDNA/models"
)

type OccurrenceRepository struct {
	mockDB *database.MockDB
}

func NewOccurrenceRepository(mockDB *database.MockDB) *OccurrenceRepository {
	return &OccurrenceRepository{
		mockDB: mockDB,
	}
}

func (occurrenceRepository *OccurrenceRepository) ListOccurrence() ([]models.Occurrence, int, error) {
	occurrences := occurrenceRepository.mockDB.OccurrenceData

	return occurrences, len(occurrences), nil
}
