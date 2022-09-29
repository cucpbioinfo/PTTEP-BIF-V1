package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type PhylumService struct {
	PhylumRepository *repositories.PhylumRepository
}

func NewPhylumService(phylumRepository *repositories.PhylumRepository) *PhylumService {
	return &PhylumService{
		PhylumRepository: phylumRepository,
	}
}

func (phylumService *PhylumService) ListPhylum(query types.ListPhylumQuery) ([]types.PhylumDto, error) {
	phylums, err := phylumService.PhylumRepository.ListPhylum(query)
	if err != nil {
		return make([]types.PhylumDto, 0), err
	}
	phylumList := make([]types.PhylumDto, len(phylums))
	for i, phylum := range phylums {
		phylumList[i] = types.PhylumDto{
			PhylumID:   phylum.PhylumID,
			PhylumName: phylum.PhylumName,
		}
	}
	return phylumList, nil
}
