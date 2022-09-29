package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type MainStationService struct {
	MainStationRepository *repositories.MainStationRepository
}

func NewMainStationService(mainStationRepository *repositories.MainStationRepository) *MainStationService {
	return &MainStationService{
		MainStationRepository: mainStationRepository,
	}
}

func (mainStationService *MainStationService) ListMainStation(query *types.MainStationListQuery) ([]types.MainStationDto, error) {
	mainStations, err := mainStationService.MainStationRepository.ListMainStation(query)
	if err != nil {
		return make([]types.MainStationDto, 0), err
	}
	mainStationList := make([]types.MainStationDto, len(mainStations))
	for i, mainStation := range mainStations {
		fieldLocations := make([]types.FieldLocationDto, len(mainStation.FieldLocations))
		for j, fieldLocation := range mainStation.FieldLocations {
			fieldLocations[j] = types.FieldLocationDto{
				FieldLocationID:   fieldLocation.FieldLocationID,
				FieldLocationName: fieldLocation.FieldLocationName,
				Location: &types.Location{
					Latitude:  fieldLocation.FieldLocationLat,
					Longitude: fieldLocation.FieldLocationLong,
				},
				Replication: fieldLocation.Replication,
				CreatedAt:   fieldLocation.CreatedAt,
				UpdatedAt:   fieldLocation.UpdatedAt,
			}
		}
		mainStationList[i] = types.MainStationDto{
			MainStationID:   mainStation.MainStationID,
			MainStationName: mainStation.MainStationName,
			RadiusFromSite:  mainStation.RadiusFromSite,
			Easting:         mainStation.Easting,
			Northing:        mainStation.Northing,
			Location: &types.Location{
				Latitude:  mainStation.MainStationLat,
				Longitude: mainStation.MainStationLong,
			},
			FieldLocations: fieldLocations,
			CreatedAt:      mainStation.CreatedAt,
			UpdatedAt:      mainStation.UpdatedAt,
		}
	}
	return mainStationList, nil
}
