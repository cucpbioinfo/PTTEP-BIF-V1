package services

import (
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type SpeciesImageService struct {
	SpeciesImageRepository *repositories.SpeciesImageRepository
}

func NewSpeciesImageService(speciesImageRepo *repositories.SpeciesImageRepository) *SpeciesImageService {
	return &SpeciesImageService{
		SpeciesImageRepository: speciesImageRepo,
	}
}

func (speciesImageService *SpeciesImageService) UploadNewFile(speciesImage types.SpeciesImageDto) (uuid.UUID, error) {
	speciesImageId, err := speciesImageService.SpeciesImageRepository.CreateSpeciesImage(&speciesImage)
	if err != nil {
		return uuid.Nil, err
	}
	return speciesImageId, nil
}
