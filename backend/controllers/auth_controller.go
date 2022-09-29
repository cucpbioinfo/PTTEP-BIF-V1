package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type AuthController struct {
	UserService *services.UserService
	AuthService *services.AuthService
}

func NewAuthController(
	userService *services.UserService,
	authService *services.AuthService,
) *AuthController {
	return &AuthController{
		UserService: userService,
		AuthService: authService,
	}
}

func (authController *AuthController) Register(c *fiber.Ctx) error {
	body := &types.RegisterBody{}
	err := c.BodyParser(body)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid body",
		}
	}
	err = authController.UserService.CreateUser(&types.UserCreateBody{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		return err
	}
	return c.JSON(&types.RegisterResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
	})
}

func (authController *AuthController) Login(c *fiber.Ctx) error {
	body := &types.LoginBody{}
	err := c.BodyParser(body)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid body",
		}
	}
	authToken, err := authController.UserService.LoginUser(body)
	if err != nil {
		return err
	}
	return c.JSON(&types.LoginResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: types.AuthTokenDto{
			AuthToken: uuid.MustParse(authToken),
		},
	})
}

func (authController *AuthController) GetMe(c *fiber.Ctx) error {
	user := c.Locals("user")
	userSession, ok := user.(types.UserSession)
	if !ok {
		return fiber.ErrInternalServerError
	}
	return c.JSON(types.MeResponse{
		BaseResponse: types.BaseResponse{
			Ok:      true,
			Message: "success",
		},
		Data: userSession,
	})
}
