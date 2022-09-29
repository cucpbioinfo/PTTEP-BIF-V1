package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type SiteService struct {
	SiteRepository *repositories.SiteRepository
}

func NewSiteService(siteRepository *repositories.SiteRepository) *SiteService {
	return &SiteService{
		SiteRepository: siteRepository,
	}
}

func (siteService *SiteService) ListSite() ([]types.SiteDto, error) {
	sites, err := siteService.SiteRepository.ListSite()
	if err != nil {
		return make([]types.SiteDto, 0), err
	}
	siteList := make([]types.SiteDto, len(sites))
	for i, site := range sites {
		mainStations := make([]types.MainStationDto, len(site.MainStations))
		for j, mainStation := range site.MainStations {
			fieldLocations := make([]types.FieldLocationDto, len(mainStation.FieldLocations))
			for k, fieldLocation := range mainStation.FieldLocations {
				fieldLocations[k] = types.FieldLocationDto{
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
			mainStations[j] = types.MainStationDto{
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
		siteList[i] = types.SiteDto{
			SiteID:       site.SiteID,
			SiteName:     site.SiteName,
			MainStations: mainStations,
			CreatedAt:    site.CreatedAt,
			UpdatedAt:    site.UpdatedAt,
		}
	}
	return siteList, nil
}
