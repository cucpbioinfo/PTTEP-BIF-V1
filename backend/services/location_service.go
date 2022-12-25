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

func (locationService *LocationService) ListLocationAsset(query types.ListLocationQuery) ([]types.LocationAssetDto, error) {
	location, err := locationService.LocationRepository.ListLocationAsset(query)
	if err != nil {
		return make([]types.LocationAssetDto, 0), err
	}
	locationList := make([]types.LocationAssetDto, len(location))
	for i, location := range location {
		locationList[i] = types.LocationAssetDto{
			LocationID: location.LocationID,
			AssetName:  location.Asset.AssetName,
			Latitude:   location.Latitude,
			Longitude:  location.Longitude,
		}
	}
	return locationList, nil
}

func (locationService *LocationService) ListLocationPlatform(query types.ListLocationQuery) ([]types.LocationPlatformDto, error) {
	location, err := locationService.LocationRepository.ListLocationSelect(query)
	if err != nil {
		return make([]types.LocationPlatformDto, 0), err
	}
	locationList := make([]types.LocationPlatformDto, len(location))
	for i, location := range location {
		locationList[i] = types.LocationPlatformDto{
			LocationID:   location.LocationID,
			AssetName:    location.Asset.AssetName,
			PlatformName: location.Platform.PlatformName,
			Latitude:     location.Latitude,
			Longitude:    location.Longitude,
		}
	}
	return locationList, nil
}

func (locationService *LocationService) ListAssetLocation(query types.ListLocationAssetQuery) ([]types.LocationAssetDto, error) {
	location, err := locationService.LocationRepository.ListAssetLocation(query)
	if err != nil {
		return make([]types.LocationAssetDto, 0), err
	}
	locationList := make([]types.LocationAssetDto, len(location))
	for i, location := range location {
		locationList[i] = types.LocationAssetDto{
			LocationID: location.LocationID,
			AssetID:    location.AssetID,
			AssetName:  location.Asset.AssetName,
			Latitude:   location.Latitude,
			Longitude:  location.Longitude,
		}
	}
	return locationList, nil
}

func (locationService *LocationService) ListCenterAssetLocation(query types.ListLocationAssetQuery) ([]types.LocationAssetDto, error) {
	location, err := locationService.LocationRepository.ListCenterAssetLocation(query)
	if err != nil {
		return make([]types.LocationAssetDto, 0), err
	}
	locationList := make([]types.LocationAssetDto, len(location))
	for i, location := range location {
		locationList[i] = types.LocationAssetDto{
			LocationID: location.LocationID,
			AssetID:    location.AssetID,
			AssetName:  location.Asset.AssetName,
			Latitude:   location.Latitude,
			Longitude:  location.Longitude,
		}
	}
	return locationList, nil
}
