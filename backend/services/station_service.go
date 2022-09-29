package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type StationService struct {
	StationRepository *repositories.StationRepository
}

func NewStationService(stationRepository *repositories.StationRepository) *StationService {
	return &StationService{
		StationRepository: stationRepository,
	}
}

func (stationService *StationService) ListStation(query types.ListStationQuery) ([]types.StationDto, error) {
	station, err := stationService.StationRepository.ListStation(query)
	if err != nil {
		return make([]types.StationDto, 0), err
	}
	stationList := make([]types.StationDto, len(station))
	for i, station := range station {
		stationList[i] = types.StationDto{
			StationID:   station.StationID,
			StationName: station.StationName,
		}
	}
	return stationList, nil
}
