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

func (diversityService *DiversityService) ListDiversity(query types.ListDiversityQuery) ([]types.DiversityDto, error) {
	diversity, err := diversityService.DiversityRepository.ListDiversity(query)
	if err != nil {
		return make([]types.DiversityDto, 0), err
	}
	diversityList := make([]types.DiversityDto, len(diversity))
	for i, diversity := range diversity {
		diversityList[i] = types.DiversityDto{
			DiversityID:   diversity.DiversityID,
			Year:          diversity.Year,
			AssetName:     diversity.Asset.AssetName,
			PlatformName:  diversity.Platform.PlatformName,
			StationName:   diversity.Station.StationName,
			Surface:       diversity.Surface,
			Euphotic_zone: diversity.Euphotic_zone,
		}
	}
	return diversityList, nil
}
