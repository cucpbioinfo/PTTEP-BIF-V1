package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (userController *UserController) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("userID")
	user := userController.UserService.GetUserByID(userID)

	return c.JSON(&types.UserGetResponse{
		UserID: user.UserID,
	})
}

func (userController *UserController) CreateUser(c *fiber.Ctx) error {
	body := &types.UserCreateBody{}
	c.BodyParser(body)
	userController.UserService.CreateUser(body)
	return c.JSON(&types.UserCreateResponse{
		types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
	})
}
