package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type DensityService struct {
	DensityRepository *repositories.DensityRepository
}

func NewDensityService(densityRepository *repositories.DensityRepository) *DensityService {
	return &DensityService{
		DensityRepository: densityRepository,
	}
}

func (densityService *DensityService) ListDensity(query types.ListDensityQuery) ([]types.DensityDto, error) {
	density, err := densityService.DensityRepository.ListDensity(query)
	if err != nil {
		return make([]types.DensityDto, 0), err
	}
	densityList := make([]types.DensityDto, len(density))
	for i, density := range density {
		densityList[i] = types.DensityDto{
			DensityID:     density.DensityID,
			Year:          density.Year,
			AssetName:     density.Asset.AssetName,
			PlatformName:  density.Platform.PlatformName,
			StationName:   density.Station.StationName,
			SpeciesID:     density.SpeciesID,
			SpeciesName:   density.SpeciesName,
			Surface:       density.Surface,
			Euphotic_zone: density.Euphotic_zone,
		}
	}
	return densityList, nil
}

// Platform platform
type PlatformDensityService struct {
	PlatformDensityRepository *repositories.PlatformDensityRepository
}

func NewPlatformDensityService(platformdensityRepository *repositories.PlatformDensityRepository) *PlatformDensityService {
	return &PlatformDensityService{
		PlatformDensityRepository: platformdensityRepository,
	}
}

func (platformdensityService *PlatformDensityService) ListPlatformDensity(query types.ListPlatformDensityQuery) ([]types.PlatformDensityDto, error) {
	density, err := platformdensityService.PlatformDensityRepository.ListPlatformDensity(query)
	if err != nil {
		return make([]types.PlatformDensityDto, 0), err
	}
	densityList := make([]types.PlatformDensityDto, len(density))
	for i, density := range density {
		densityList[i] = types.PlatformDensityDto{
			DensityID:     density.DensityID,
			Year:          density.Year,
			AssetName:     density.Asset.AssetName,
			PlatformName:  density.Platform.PlatformName,
			SpeciesID:     density.SpeciesID,
			SpeciesName:   density.SpeciesName,
			Surface:       density.Surface,
			Euphotic_zone: density.Euphotic_zone,
		}
	}
	return densityList, nil
}

// Asset asset
type AssetDensityService struct {
	AssetDensityRepository *repositories.AssetDensityRepository
}

func NewAssetDensityService(assetdensityRepository *repositories.AssetDensityRepository) *AssetDensityService {
	return &AssetDensityService{
		AssetDensityRepository: assetdensityRepository,
	}
}

func (assetdensityService *AssetDensityService) ListAssetDensity(query types.ListAssetDensityQuery) ([]types.AssetDensityDto, error) {
	density, err := assetdensityService.AssetDensityRepository.ListAssetDensity(query)
	if err != nil {
		return make([]types.AssetDensityDto, 0), err
	}
	densityList := make([]types.AssetDensityDto, len(density))
	for i, density := range density {
		densityList[i] = types.AssetDensityDto{
			DensityID:     density.DensityID,
			Year:          density.Year,
			AssetName:     density.Asset.AssetName,
			SpeciesID:     density.SpeciesID,
			SpeciesName:   density.SpeciesName,
			Surface:       density.Surface,
			Euphotic_zone: density.Euphotic_zone,
		}
	}
	return densityList, nil
}
