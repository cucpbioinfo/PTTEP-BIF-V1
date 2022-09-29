package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type FamilyService struct {
	FamilyRepository *repositories.FamilyRepository
}

func NewFamilyService(familyRepository *repositories.FamilyRepository) *FamilyService {
	return &FamilyService{
		FamilyRepository: familyRepository,
	}
}

func (familyService *FamilyService) ListFamily(query types.ListFamilyQuery) ([]types.FamilyDto, error) {
	familes, err := familyService.FamilyRepository.ListFamily(query)
	if err != nil {
		return make([]types.FamilyDto, 0), err
	}
	familyList := make([]types.FamilyDto, len(familes))
	for i, family := range familes {
		familyList[i] = types.FamilyDto{
			FamilyID:   family.FamilyID,
			FamilyName: family.FamilyName,
		}
	}
	return familyList, nil
}
