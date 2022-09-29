package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type LocationService struct {
	LocationRepository *repositories.LocationRepository
}

func NewLocationService(locationRepository *repositories.LocationRepository) *LocationService {
	return &LocationService{
		LocationRepository: locationRepository,
	}
}

func (locationService *LocationService) ListLocation(query types.ListLocationQuery) ([]types.LocationDto, error) {
	location, err := locationService.LocationRepository.ListLocation(query)
	if err != nil {
		return make([]types.LocationDto, 0), err
	}
	locationList := make([]types.LocationDto, len(location))
	for i, location := range location {
		locationList[i] = types.LocationDto{
			LocationID:   location.LocationID,
			AssetName:    location.Asset.AssetName,
			PlatformName: location.Platform.PlatformName,
			StationName:  location.Station.StationName,
			Latitude:     location.Latitude,
			Longitude:    location.Longitude,
		}
	}
	return locationList, nil
}
