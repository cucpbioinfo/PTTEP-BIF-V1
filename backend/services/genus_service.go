package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type GenusService struct {
	GenusRepository *repositories.GenusRepository
}

func NewGenusService(genusRepository *repositories.GenusRepository) *GenusService {
	return &GenusService{
		GenusRepository: genusRepository,
	}
}

func (genusService *GenusService) ListGenus(query types.ListGenusQuery) ([]types.GenusDto, error) {
	genus, err := genusService.GenusRepository.ListGenus(query)
	if err != nil {
		return make([]types.GenusDto, 0), err
	}
	genusList := make([]types.GenusDto, len(genus))
	for i, genus := range genus {
		genusList[i] = types.GenusDto{
			GenusID:   genus.GenusID,
			GenusName: genus.GenusName,
		}
	}
	return genusList, nil
}
