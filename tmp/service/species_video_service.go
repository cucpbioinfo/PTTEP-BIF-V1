package services

import (
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type SpeciesVideoService struct {
	SpeciesVideoRepository *repositories.SpeciesVideoRepository
}

func NewSpeciesVideoService(speciesVideoRepo *repositories.SpeciesVideoRepository) *SpeciesVideoService {
	return &SpeciesVideoService{
		SpeciesVideoRepository: speciesVideoRepo,
	}
}

func (speciesVideoService *SpeciesVideoService) UploadNewFile(speciesVideo types.SpeciesVideoDto) (uuid.UUID, error) {
	speciesVideoId, err := speciesVideoService.SpeciesVideoRepository.CreateSpeciesVideo(&speciesVideo)
	if err != nil {
		return uuid.Nil, err
	}
	return speciesVideoId, nil
}
