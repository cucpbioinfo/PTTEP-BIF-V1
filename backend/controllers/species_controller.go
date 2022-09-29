package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
	"github.com/wichadak/eDNA/utils"
)

type SpeciesController struct {
	SpeciesService *services.SpeciesService
}

func NewSpeciesController(speciesService *services.SpeciesService) *SpeciesController {
	return &SpeciesController{
		SpeciesService: speciesService,
	}
}

func (speciesController *SpeciesController) ListSpecies(c *fiber.Ctx) error {
	query := &types.SpeciesListQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}

	pageNumber, pageSize := utils.GetValidPagination(query.PageNumber, query.PageSize)

	species, total, err := speciesController.SpeciesService.ListSpecies(types.SpeciesListQuery{
		Keyword:      query.Keyword,
		MajorGroupID: query.MajorGroupID,
		KingdomID:    query.KingdomID,
		PhylumID:     query.PhylumID,
		ClassID:      query.ClassID,
		OrderID:      query.OrderID,
		FamilyID:     query.FamilyID,
		GenusID:      query.GenusID,
		PageNumber:   pageNumber,
		PageSize:     pageSize,
	})
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list species",
		}
	}
	return c.JSON(&types.SpeciesListResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: species,
		Meta: types.PaginationResponse{
			PageNumber: pageNumber,
			PageSize:   pageSize,
			Total:      total,
		},
	})
}

// func (speciesController *SpeciesController) GetSpeciesCountGroupBySpeciesType(c *fiber.Ctx) error {
// 	result, err := speciesController.SpeciesService.GetSpeciesCountGroupBySpeciesType()
// 	if err != nil {
// 		return &fiber.Error{
// 			Code:    400,
// 			Message: err.Error(),
// 		}
// 	}

// 	return c.JSON(&types.SpeciesCountGroupBySpeciesTypeResponse{
// 		BaseResponse: types.BaseResponse{
// 			Ok:      true,
// 			Message: "success",
// 		},
// 		Data: result,
// 	})
// }

func (speciesController *SpeciesController) GetSpeciesDetails(c *fiber.Ctx) error {
	speciesId, err := uuid.Parse(c.Params("speciesId"))
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid speciesId",
		}
	}

	result, err := speciesController.SpeciesService.GetSpeciesDetails(speciesId)
	if err != nil {
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return c.JSON(&types.SpeciesDetailsResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: result,
	})
}
