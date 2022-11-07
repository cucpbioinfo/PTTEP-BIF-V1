package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type SummaryService struct {
	SummaryRepository *repositories.SummaryRepository
}

func NewSummaryService(summaryRepository *repositories.SummaryRepository) *SummaryService {
	return &SummaryService{
		SummaryRepository: summaryRepository,
	}
}

func (summaryService *SummaryService) ListSummary(query types.ListSummaryQuery) ([]types.SummaryDto, error) {
	summary, err := summaryService.SummaryRepository.ListSummary(query)
	if err != nil {
		return make([]types.SummaryDto, 0), err
	}
	summaryList := make([]types.SummaryDto, len(summary))
	for i, summary := range summary {
		summaryList[i] = types.SummaryDto{
			SummaryID:            summary.SummaryID,
			Year:                 summary.Year,
			MajorGroupName:       summary.MajorGroup.MajorGroupName,
			IdentificationName:   summary.Identification.IdentificationName,
			AssetName:            summary.Asset.AssetName,
			PlatformName:         summary.Platform.PlatformName,
			StationName:          summary.Station.StationName,
			SurfaceShannon:       summary.SurfaceShannon,
			SurfaceNumber:        summary.SurfaceNumber,
			SurfaceMax:           summary.SurfaceMax,
			SurfaceEvenness:      summary.SurfaceEvenness,
			EuphoticZoneShannon:  summary.EuphoticZoneShannon,
			EuphoticZoneNumber:   summary.EuphoticZoneNumber,
			EuphoticZoneMax:      summary.EuphoticZoneMax,
			EuphoticZoneEvenness: summary.EuphoticZoneEvenness,
			AverageShannon:       summary.AverageShannon,
			AverageNumber:        summary.AverageNumber,
			AverageMax:           summary.AverageMax,
			AverageEvenness:      summary.AverageEvenness,
		}
	}
	return summaryList, nil
}

// all summary
func (summaryService *SummaryService) ListAllSummary(query types.ListSummaryQuery) ([]types.SummaryDto, error) {
	summary, err := summaryService.SummaryRepository.ListAllSummary(query)
	if err != nil {
		return make([]types.SummaryDto, 0), err
	}
	summaryList := make([]types.SummaryDto, len(summary))
	for i, summary := range summary {
		summaryList[i] = types.SummaryDto{
			SummaryID:            summary.SummaryID,
			Year:                 summary.Year,
			MajorGroupName:       summary.MajorGroup.MajorGroupName,
			IdentificationName:   summary.Identification.IdentificationName,
			AssetName:            summary.Asset.AssetName,
			PlatformName:         summary.Platform.PlatformName,
			StationName:          summary.Station.StationName,
			SurfaceShannon:       summary.SurfaceShannon,
			SurfaceNumber:        summary.SurfaceNumber,
			SurfaceMax:           summary.SurfaceMax,
			SurfaceEvenness:      summary.SurfaceEvenness,
			EuphoticZoneShannon:  summary.EuphoticZoneShannon,
			EuphoticZoneNumber:   summary.EuphoticZoneNumber,
			EuphoticZoneMax:      summary.EuphoticZoneMax,
			EuphoticZoneEvenness: summary.EuphoticZoneEvenness,
			AverageShannon:       summary.AverageShannon,
			AverageNumber:        summary.AverageNumber,
			AverageMax:           summary.AverageMax,
			AverageEvenness:      summary.AverageEvenness,
		}
	}
	return summaryList, nil
}

// Platform platform
type PlatformSummaryService struct {
	PlatformSummaryRepository *repositories.PlatformSummaryRepository
}

func NewPlatformSummaryService(platformsummaryRepository *repositories.PlatformSummaryRepository) *PlatformSummaryService {
	return &PlatformSummaryService{
		PlatformSummaryRepository: platformsummaryRepository,
	}
}

func (platformsummaryService *PlatformSummaryService) ListPlatformSummary(query types.ListPlatformSummaryQuery) ([]types.PlatformSummaryDto, error) {
	summary, err := platformsummaryService.PlatformSummaryRepository.ListPlatformSummary(query)
	if err != nil {
		return make([]types.PlatformSummaryDto, 0), err
	}
	summaryList := make([]types.PlatformSummaryDto, len(summary))
	for i, summary := range summary {
		summaryList[i] = types.PlatformSummaryDto{
			SummaryID:            summary.SummaryID,
			Year:                 summary.Year,
			MajorGroupName:       summary.MajorGroup.MajorGroupName,
			IdentificationName:   summary.Identification.IdentificationName,
			AssetName:            summary.Asset.AssetName,
			PlatformName:         summary.Platform.PlatformName,
			SurfaceShannon:       summary.SurfaceShannon,
			SurfaceNumber:        summary.SurfaceNumber,
			SurfaceMax:           summary.SurfaceMax,
			SurfaceEvenness:      summary.SurfaceEvenness,
			EuphoticZoneShannon:  summary.EuphoticZoneShannon,
			EuphoticZoneNumber:   summary.EuphoticZoneNumber,
			EuphoticZoneMax:      summary.EuphoticZoneMax,
			EuphoticZoneEvenness: summary.EuphoticZoneEvenness,
			AverageShannon:       summary.AverageShannon,
			AverageNumber:        summary.AverageNumber,
			AverageMax:           summary.AverageMax,
			AverageEvenness:      summary.AverageEvenness,
		}
	}
	return summaryList, nil
}

// Asset asset
type AssetSummaryService struct {
	AssetSummaryRepository *repositories.AssetSummaryRepository
}

func NewAssetSummaryService(assetsummaryRepository *repositories.AssetSummaryRepository) *AssetSummaryService {
	return &AssetSummaryService{
		AssetSummaryRepository: assetsummaryRepository,
	}
}

func (assetsummaryService *AssetSummaryService) ListAssetSummary(query types.ListAssetSummaryQuery) ([]types.AssetSummaryDto, error) {
	summary, err := assetsummaryService.AssetSummaryRepository.ListAssetSummary(query)
	if err != nil {
		return make([]types.AssetSummaryDto, 0), err
	}
	summaryList := make([]types.AssetSummaryDto, len(summary))
	for i, summary := range summary {
		summaryList[i] = types.AssetSummaryDto{
			SummaryID:            summary.SummaryID,
			Year:                 summary.Year,
			MajorGroupName:       summary.MajorGroup.MajorGroupName,
			IdentificationName:   summary.Identification.IdentificationName,
			AssetName:            summary.Asset.AssetName,
			SurfaceShannon:       summary.SurfaceShannon,
			SurfaceNumber:        summary.SurfaceNumber,
			SurfaceMax:           summary.SurfaceMax,
			SurfaceEvenness:      summary.SurfaceEvenness,
			EuphoticZoneShannon:  summary.EuphoticZoneShannon,
			EuphoticZoneNumber:   summary.EuphoticZoneNumber,
			EuphoticZoneMax:      summary.EuphoticZoneMax,
			EuphoticZoneEvenness: summary.EuphoticZoneEvenness,
			AverageShannon:       summary.AverageShannon,
			AverageNumber:        summary.AverageNumber,
			AverageMax:           summary.AverageMax,
			AverageEvenness:      summary.AverageEvenness,
		}
	}
	return summaryList, nil
}

// year Year
type YearSummaryService struct {
	YearSummaryRepository *repositories.YearSummaryRepository
}

func NewYearSummaryService(yearsummaryRepository *repositories.YearSummaryRepository) *YearSummaryService {
	return &YearSummaryService{
		YearSummaryRepository: yearsummaryRepository,
	}
}

func (yearsummaryService *YearSummaryService) ListYearSummary(query types.ListYearSummaryQuery) ([]types.YearSummaryDto, error) {
	yearsummary, err := yearsummaryService.YearSummaryRepository.ListYearSummary(query)
	if err != nil {
		return make([]types.YearSummaryDto, 0), err
	}
	yearsummaryList := make([]types.YearSummaryDto, len(yearsummary))
	for i, yearsummary := range yearsummary {
		yearsummaryList[i] = types.YearSummaryDto{
			Year: yearsummary.Year,
		}
	}
	return yearsummaryList, nil
}
