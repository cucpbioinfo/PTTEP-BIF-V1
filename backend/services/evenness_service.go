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

func (evennessService *EvennessService) ListEvenness(query types.ListEvennessQuery) ([]types.EvennessDto, error) {
	evenness, err := evennessService.EvennessRepository.ListEvenness(query)
	if err != nil {
		return make([]types.EvennessDto, 0), err
	}
	evennessList := make([]types.EvennessDto, len(evenness))
	for i, evenness := range evenness {
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
