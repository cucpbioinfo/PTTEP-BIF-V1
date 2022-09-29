package types

import "github.com/google/uuid"

type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	BaseResponse
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthTokenDto struct {
	AuthToken uuid.UUID `json:"authToken"`
}

type LoginResponse struct {
	BaseResponse
	Data AuthTokenDto `json:"data"`
}

type UserSession struct {
	UserID uuid.UUID `json:"userId"`
	Email  string    `json:"email"`
}

type MeResponse struct {
	BaseResponse
	Data UserSession `json:"data"`
}
