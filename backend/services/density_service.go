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
			StationID:     density.StationID,
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
			PlatformID:    density.PlatformID,
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
			AssetID:       density.AssetID,
			AssetName:     density.Asset.AssetName,
			SpeciesID:     density.SpeciesID,
			SpeciesName:   density.SpeciesName,
			Surface:       density.Surface,
			Euphotic_zone: density.Euphotic_zone,
		}
	}
	return densityList, nil
}

// YearAsset yearasset
type YearAssetDensityService struct {
	YearAssetDensityRepository *repositories.YearAssetDensityRepository
}

func NewYearAssetDensityService(yearassetdensityRepository *repositories.YearAssetDensityRepository) *YearAssetDensityService {
	return &YearAssetDensityService{
		YearAssetDensityRepository: yearassetdensityRepository,
	}
}

func (yearassetdensityService *YearAssetDensityService) ListYearAssetDensity(query types.ListYearAssetDensityQuery) ([]types.YearAssetDensityDto, error) {
	density, err := yearassetdensityService.YearAssetDensityRepository.ListYearAssetDensity(query)
	if err != nil {
		return make([]types.YearAssetDensityDto, 0), err
	}
	densityList := make([]types.YearAssetDensityDto, len(density))
	for i, density := range density {
		densityList[i] = types.YearAssetDensityDto{
			Year: density.Year,
		}
	}
	return densityList, nil
}

// Year Platform year platform
type YearPlatformDensityService struct {
	YearPlatformDensityRepository *repositories.YearPlatformDensityRepository
}

func NewYearPlatformDensityService(yearplatformdensityRepository *repositories.YearPlatformDensityRepository) *YearPlatformDensityService {
	return &YearPlatformDensityService{
		YearPlatformDensityRepository: yearplatformdensityRepository,
	}
}

func (yearplatformdensityService *YearPlatformDensityService) ListYearPlatformDensity(query types.ListYearPlatformDensityQuery) ([]types.YearPlatformDensityDto, error) {
	density, err := yearplatformdensityService.YearPlatformDensityRepository.ListYearPlatformDensity(query)
	if err != nil {
		return make([]types.YearPlatformDensityDto, 0), err
	}
	densityList := make([]types.YearPlatformDensityDto, len(density))
	for i, density := range density {
		densityList[i] = types.YearPlatformDensityDto{
			Year: density.Year,
		}
	}
	return densityList, nil
}

// Year Station year station
type YearDensityService struct {
	YearDensityRepository *repositories.YearDensityRepository
}

func NewYearDensityService(yeardensityRepository *repositories.YearDensityRepository) *YearDensityService {
	return &YearDensityService{
		YearDensityRepository: yeardensityRepository,
	}
}

func (yeardensityService *YearDensityService) ListYearDensity(query types.ListYearDensityQuery) ([]types.YearDensityDto, error) {
	density, err := yeardensityService.YearDensityRepository.ListYearDensity(query)
	if err != nil {
		return make([]types.YearDensityDto, 0), err
	}
	densityList := make([]types.YearDensityDto, len(density))
	for i, density := range density {
		densityList[i] = types.YearDensityDto{
			Year: density.Year,
		}
	}
	return densityList, nil
}
