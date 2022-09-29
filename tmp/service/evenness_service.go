package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type EvennessService struct {
	EvennessRepository *repositories.EvennessRepository
}

func NewEvennessService(evennessRepository *repositories.EvennessRepository) *EvennessService {
	return &EvennessService{
		EvennessRepository: evennessRepository,
	}
}

func (evennessService *EvennessService) ListEvenness() ([]types.EvennessDto, error) {
	Evennesses, err := evennessService.EvennessRepository.ListEvenness()
	if err != nil {
		return make([]types.EvennessDto, 0), err
	}
	evennessList := make([]types.EvennessDto, len(Evennesses))
	for i, evenness := range Evennesses {
		evennessList[i] = types.EvennessDto{
			EvennessID:    evenness.EvennessID,
			Year:          evenness.Year,
			AssetName:     evenness.Asset.AssetName,
			PlatformName:  evenness.Platform.PlatformName,
			StationName:   evenness.Station.StationName,
			Surface:       evenness.Surface,
			Euphotic_zone: evenness.Euphotic_zone,
		}
	}
	return evennessList, nil
}

// func (evennessService *EvennessService) ListEvennessYear(query types.EvennessListQuery) ([]types.EvennessYearDto, error) {
// 	Evennesses, err := evennessService.EvennessRepository.ListEvennessYear(query)
// 	if err != nil {
// 		return make([]types.EvennessYearDto, 0), err
// 	}
// 	evennessList := make([]types.EvennessYearDto, len(Evennesses))
// 	for i, evenness := range Evennesses {
// 		evennessList[i] = types.EvennessYearDto{
// 			Year: evenness.Year,
// 		}
// 	}
// 	return evennessList, nil
// }

// func (evennessService *EvennessService) ListEvennessProject(query types.EvennessListQuery) ([]types.EvennessProjectDto, error) {
// 	Evennesses, err := evennessService.EvennessRepository.ListEvennessProject(query)
// 	if err != nil {
// 		return make([]types.EvennessProjectDto, 0), err
// 	}
// 	evennessList := make([]types.EvennessProjectDto, len(Evennesses))
// 	for i, evenness := range Evennesses {
// 		evennessList[i] = types.EvennessProjectDto{
// 			Year:        evenness.Year,
// 			ProjectID:   evenness.ProjectID,
// 			ProjectName: evenness.Project.ProjectName,
// 		}
// 	}
// 	return evennessList, nil
// }
