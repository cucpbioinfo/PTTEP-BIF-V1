package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
	"github.com/wichadak/eDNA/utils"
)

type DensityController struct {
	DensityService *services.DensityService
}

func NewDensityController(densityService *services.DensityService) *DensityController {
	return &DensityController{
		DensityService: densityService,
	}
}

func (densityController *DensityController) ListDensity(c *fiber.Ctx) error {
	query := &types.DensityListQuery{}
	if err := c.QueryParser(query); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid query",
		}
	}
	pageNumber, pageSize := utils.GetValidPagination(query.PageNumber, query.PageSize)
	// densities, err := densityController.DensityService.ListDensity()
	// if err != nil {
	// 	return &fiber.Error{
	// 		Code:    400,
	// 		Message: "Fail to list density",
	// 	}
	// }
	densities, total, err := densityController.DensityService.ListDensity(types.DensityListQuery{
		Keyword:    query.Keyword,
		AssetID:    query.AssetID,
		PlatformID: query.PlatformID,
		StationID:  query.StationID,
		// IdentificationID: query.IdentificationID,
		// MethodID:         query.MethodID,
		// MajorGroupID:     query.MajorGroupID,
		// KingdomID:        query.KingdomID,
		// PhylumID:         query.PhylumID,
		// ClassID:          query.ClassID,
		// OrderID:          query.OrderID,
		// FamilyID:         query.FamilyID,
		// GenusID:          query.GenusID,
		SpeciesID: query.SpeciesID,
		// PageNumber: pageNumber,
		// PageSize:   pageSize,
	})
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list species",
		}
	}
	return c.JSON(&types.ListDensityResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: densities,
		Meta: types.PaginationResponse{
			// PageNumber: pageNumber,
			// PageSize:   pageSize,
			Total: total,
		},
	})
}

//backup
// func (densityController *DensityController) ListDensity(c *fiber.Ctx) error {
// 	densitys, err := densityController.DensityService.ListDensity()
// 	if err != nil {
// 		return &fiber.Error{
// 			Code:    400,
// 			Message: "Fail to list density",
// 		}
// 	}
// 	return c.JSON(&types.ListDensityResponse{
// 		BaseResponse: types.BaseResponse{
// 			Ok:      true,
// 			Message: "success",
// 		},
// 		Data: densitys,
// 	})
// }
