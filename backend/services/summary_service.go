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
