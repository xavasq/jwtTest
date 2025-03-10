package service

import (
	"ECCO2K/internal/models"
	"ECCO2K/internal/repository"
	"ECCO2K/internal/security"
	"context"
	"fmt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	if user.Name == "" || user.Password == "" {
		return fmt.Errorf("имя и пароль не могут быть пустыми")
	}

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	return s.repo.GetUserByID(ctx, id)
}
