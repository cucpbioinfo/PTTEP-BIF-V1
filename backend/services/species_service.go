package services

import (
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type SpeciesService struct {
	SpeciesRepository *repositories.SpeciesRepository
}

func NewSpeciesService(speciesRepository *repositories.SpeciesRepository) *SpeciesService {
	return &SpeciesService{
		SpeciesRepository: speciesRepository,
	}
}

func (speciesService *SpeciesService) ListSpecies(query types.SpeciesListQuery) ([]types.SpeciesListDto, int, error) {
	species, total, err := speciesService.SpeciesRepository.ListSpecies(query)
	if err != nil {
		return make([]types.SpeciesListDto, 0), 0, err
	}
	speciesList := make([]types.SpeciesListDto, len(species))
	for i, species := range species {
		// speciesImage := ""
		// if species.SpeciesImages != nil {
		// 	if len(species.SpeciesImages) > 0 {
		// 		speciesImage = species.SpeciesImages[0].ImageSource
		// 	}
		// }
		speciesList[i] = types.SpeciesListDto{
			SpeciesID:   species.SpeciesID,
			SpeciesName: species.SpeciesName,
			// KingdomName: species.Kingdom.KingdomName,
			// PhylumName:  species.Phylum.PhylumName,
			//SpeciesImage: speciesImage,
			FamilyName: species.Family.FamilyName,
			GenusName:  species.Genus.GenusName,
		}
	}
	return speciesList, total, nil
}

// func (speciesService *SpeciesService) ListSpecies(query types.SpeciesListQuery) ([]types.SpeciesListDto, int, error) {
// 	species, total, err := speciesService.SpeciesRepository.ListSpecies()
// 	if err != nil {
// 		return make([]types.SpeciesListDto, 0), 0, err
// 	}
// 	speciesList := make([]types.SpeciesListDto, len(species))
// 	for i, species := range species {
// 		speciesList[i] = types.SpeciesListDto{
// 			SpeciesID:   species.SpeciesID,
// 			SpeciesName: species.SpeciesName,
// 			KingdomName: species.Kingdom.KingdomName,
// 			PhylumName:  species.Phylum.PhylumName,
// 		}
// 	}
// 	return speciesList, total, nil
// }

// func (speciesService *SpeciesService) GetSpeciesCountGroupBySpeciesType() ([]types.SpeciesCountGroupBySpeciesType, error) {
// 	result, err := speciesService.SpeciesRepository.GetSpeciesCountGroupBySpeciesType()
// 	if err != nil {
// 		return make([]types.SpeciesCountGroupBySpeciesType, 0), nil
// 	}
// 	return result, nil
// }

func (speciesService *SpeciesService) GetSpeciesDetails(speciesId uuid.UUID) (types.SpeciesDetailsDto, error) {
	species, err := speciesService.SpeciesRepository.GetSpeciesById(speciesId)
	if err != nil {
		return types.SpeciesDetailsDto{}, err
	}

	// familyName := ""
	// if species.Family != nil {
	// 	familyName = species.Family.FamilyName
	// }

	speciesDetails := types.SpeciesDetailsDto{
		SpeciesID:   species.SpeciesID,
		SpeciesName: species.SpeciesName,
		// ScientificName: species.ScientificName,
		IdentificationID:   species.Identification.IdentificationID,
		IdentificationName: species.Identification.IdentificationName,
		MethodID:           species.Method.MethodID,
		MethodName:         species.Method.MethodName,
		MajorGroupID:       species.MajorGroup.MajorGroupID,
		MajorGroupName:     species.MajorGroup.MajorGroupName,
		KingdomID:          species.Kingdom.KingdomID,
		KingdomName:        species.Kingdom.KingdomName,
		PhylumID:           species.Phylum.PhylumID,
		PhylumName:         species.Phylum.PhylumName,
		ClassID:            species.Class.ClassID,
		ClassName:          species.Class.ClassName,
		OrderID:            species.Order.OrderID,
		OrderName:          species.Order.OrderName,
		FamilyID:           species.Family.FamilyID,
		FamilyName:         species.Family.FamilyName,
		GenusID:            species.Genus.GenusID,
		GenusName:          species.Genus.GenusName,
	}
	return speciesDetails, nil
}
