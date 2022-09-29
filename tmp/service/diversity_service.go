package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type DiversityService struct {
	DiversityRepository *repositories.DiversityRepository
}

func NewDiversityService(diversityRepository *repositories.DiversityRepository) *DiversityService {
	return &DiversityService{
		DiversityRepository: diversityRepository,
	}
}

func (diversityService *DiversityService) ListDiversity(query types.DiversityListQuery) ([]types.DiversityDto, error) {
	Diversities, err := diversityService.DiversityRepository.ListDiversity(query)
	if err != nil {
		return make([]types.DiversityDto, 0), err
	}
	diversityList := make([]types.DiversityDto, len(Diversities))
	for i, diversity := range Diversities {
		diversityList[i] = types.DiversityDto{
			DiversityID:  diversity.DiversityID,
			Diversity:    diversity.Diversity,
			Year:         diversity.Year,
			ProjectName:  diversity.Project.ProjectName,
			PlatformName: diversity.Platform.PlatformName,
			StationName:  diversity.Station.StationName,
		}
	}
	return diversityList, nil
}
