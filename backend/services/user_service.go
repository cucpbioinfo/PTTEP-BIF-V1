package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
	"github.com/wichadak/eDNA/utils/validators"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repositories.UserRepository
	Redis          *redis.Client
}

func NewUserService(userRepository *repositories.UserRepository, rdb *redis.Client) *UserService {
	return &UserService{
		UserRepository: userRepository,
		Redis:          rdb,
	}
}

func (userService *UserService) GetUserByID(userID string) models.User {
	user, err := userService.UserRepository.GetUserByID(userID)
	if err != nil {
		log.Fatal("error")
	}
	return user
}

func (userService *UserService) CreateUser(user *types.UserCreateBody) error {
	// TODO: add email and password validation
	isValidEmail := validators.IsValidEmail(user.Email)
	if !isValidEmail {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid email",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 13)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid password",
		}
	}

	err = userService.UserRepository.CreateUser(&models.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Cannot create user",
		}
	}

	return nil
}

func (userService *UserService) LoginUser(loginBody *types.LoginBody) (string, error) {
	user, err := userService.UserRepository.GetUserByEmail(loginBody.Email)
	if err != nil {
		return "", &fiber.Error{
			Code:    400,
			Message: "Invalid email or password",
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginBody.Password))
	if err != nil {
		return "", &fiber.Error{
			Code:    400,
			Message: "Invalid email or password",
		}
	}

	// generate token and store in redis
	authToken := uuid.New().String()
	userSession, err := json.Marshal(types.UserSession{
		UserID: user.UserID,
		Email:  user.Email,
	})
	if err != nil {
		return "", &fiber.Error{
			Code:    500,
			Message: "Login failed",
		}
	}
	_, err = userService.Redis.SetEX(context.Background(), authToken, userSession, time.Hour).Result()
	if err != nil {
		return "", &fiber.Error{
			Code:    500,
			Message: "Login failed",
		}
	}
	return authToken, nil
}
