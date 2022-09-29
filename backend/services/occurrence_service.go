package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type OccurrenceService struct {
	OccurrenceRepository *repositories.OccurrenceRepository
}

func NewOccurrenceService(occurrenceRepository *repositories.OccurrenceRepository) *OccurrenceService {
	return &OccurrenceService{
		OccurrenceRepository: occurrenceRepository,
	}
}

func (occurrenceService *OccurrenceService) ListOccurrence() ([]types.OccurrenceDto, int, error) {
	occurrences, total, err := occurrenceService.OccurrenceRepository.ListOccurrence()
	if err != nil {
		return make([]types.OccurrenceDto, 0), 0, err
	}

	occurrenceList := make([]types.OccurrenceDto, len(occurrences))
	for i, occurrence := range occurrences {
		occurrenceList[i] = types.OccurrenceDto{
			OccurrenceId:      occurrence.OccurrenceId,
			ImageSrc:          occurrence.ImageSrc,
			OccurrenceDetails: occurrence.OccurrenceDetails,
			DiscoveredAt:      occurrence.DiscoveredAt,
		}
	}

	return occurrenceList, total, nil
}
