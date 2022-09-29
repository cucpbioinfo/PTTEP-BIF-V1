package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/models"
)

type UserRepository struct {
	pg *pg.DB
}

func NewUserRepository(pg *pg.DB) *UserRepository {
	return &UserRepository{
		pg: pg,
	}
}

func (userRepository UserRepository) GetUserByID(userID string) (models.User, error) {
	user := &models.User{UserID: uuid.MustParse(userID)}
	err := userRepository.pg.Model(user).WherePK().Limit(1).Select()
	if err != nil {
		return models.User{}, err
	}
	return *user, err
}

func (userRepository UserRepository) CreateUser(user *models.User) error {
	_, err := userRepository.pg.Model(user).Insert()
	return err
}

func (userRepository UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user = &models.User{}
	err := userRepository.pg.Model(user).
		Where("email = ?", email).
		Limit(1).
		Select()
	if err != nil {
		return models.User{}, err
	}
	return *user, err
}
