package service

import (
	"ECCO2K/internal/models"
	"ECCO2K/internal/repository"
	"ECCO2K/internal/security"
	"context"
	"errors"
)

type AuthService struct {
	repo   *repository.UserRepository
	secret string
}

func NewAuthService(repo *repository.UserRepository, secret string) *AuthService {
	return &AuthService{repo: repo, secret: secret}
}

func (s *AuthService) Register(ctx context.Context, name, password string) (*models.User, error) {
	if name == "" || password == "" {
		return nil, errors.New("имя и пароль не могут быть пустыми")
	}

	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     name,
		Password: hashedPassword,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return user, err

}

func (s *AuthService) Login(ctx context.Context, name, password string) (string, error) {
	user, err := s.repo.GetUserByName(ctx, name)
	if err != nil {
		return "", err
	}

	if !security.CheckPassword(user.Password, password) {
		return "", err
	}

	token, err := security.GenerateToken(user.ID, user.Name)
	if err != nil {
		return "", err
	}
	return token, nil
}
