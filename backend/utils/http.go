package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/types"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	e, ok := err.(*fiber.Error)
	if ok {
		code = e.Code
	}
	err = ctx.Status(code).JSON(types.BaseResponse{
		Ok:      false,
		Message: e.Message,
	})
	if err != nil {
		return ctx.Status(500).SendString("Internal Server Error")
	}
	return nil
}
