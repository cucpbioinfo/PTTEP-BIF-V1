package services

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type AuthService struct {
	UserRepository *repositories.UserRepository
	Redis          *redis.Client
}

func NewAuthService(userRepository *repositories.UserRepository, rdb *redis.Client) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
		Redis:          rdb,
	}
}

func (authService *AuthService) GetUserSession(authToken uuid.UUID) (types.UserSession, error) {
	userSession := types.UserSession{}
	userSessionString, err := authService.Redis.Get(context.Background(), authToken.String()).Result()
	if err != nil {
		return types.UserSession{}, err
	}

	err = json.Unmarshal([]byte(userSessionString), &userSession)

	if err != nil {
		return types.UserSession{}, err
	}
	return userSession, nil
}
