package repository

import (
	"ECCO2K/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (name, password) VALUES ($1, $2) RETURNING id"

	if err := r.db.QueryRow(ctx, query, user.Name, user.Password).Scan(&user.ID); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	query := "SELECT id, name, password FROM users WHERE id = $1"

	var user models.User
	if err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByName(ctx context.Context, name string) (*models.User, error) {
	query := "SELECT id, name, password FROM users WHERE name = $1"

	var user models.User
	if err := r.db.QueryRow(ctx, query, name).Scan(&user.ID, &user.Name, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}
