package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type KingdomController struct {
	KingdomService *services.KingdomService
}

func NewKingdomController(kingdomService *services.KingdomService) *KingdomController {
	return &KingdomController{
		KingdomService: kingdomService,
	}
}

func (kingdomController *KingdomController) ListKingdom(c *fiber.Ctx) error {
	kingdoms, err := kingdomController.KingdomService.ListKingdom()
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Fail to list kingdom",
		}
	}
	return c.JSON(&types.ListKingdomResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: kingdoms,
	})
}
