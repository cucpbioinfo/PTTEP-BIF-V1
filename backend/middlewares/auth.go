package middlewares

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/types"
)

type AuthMiddlewareCtx struct {
	Redis *redis.Client
}

func NewAuthMiddleware(rdb *redis.Client) *AuthMiddlewareCtx {
	return &AuthMiddlewareCtx{
		Redis: rdb,
	}
}

func (ctx AuthMiddlewareCtx) AuthenticationRequired(c *fiber.Ctx) error {
	authTokenString := string(c.Request().Header.Peek("Authorization"))
	authToken, err := uuid.Parse(authTokenString)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	userSession := types.UserSession{}

	userSessionString, err := ctx.Redis.Get(context.Background(), authToken.String()).Result()
	if err != nil {
		return fiber.ErrUnauthorized
	}

	err = json.Unmarshal([]byte(userSessionString), &userSession)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	c.Locals("user", userSession)
	c.Next()
	return nil
}
