package repositories

import (
	"context"
	"fmt"
	"golang-blueprint-v1/internal/database"
	"golang-blueprint-v1/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Create(ctx context.Context, input *models.RegisterRequest) error
	FindUserByEmail(ctx context.Context, email string) (*models.FindByEmailResponse, error)
}

type UserRepositoryImpl struct {
	db *database.DB
}

func NewUserRepositoryImpl(db *database.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, input *models.RegisterRequest) error {
	sql := `
		insert into users (email, password_hash) values($1, $2)
	`
	if _, err := repo.db.Pool.Exec(ctx, sql, input.Email, input.Password); err != nil {
		return fmt.Errorf("failed to create user : %s", err)
	}
	return nil
}

func (repo *UserRepositoryImpl) FindUserByEmail(ctx context.Context, email string) (*models.FindByEmailResponse, error) {
	sql := `select id, email, created_at from users where email = $1`

	user := &models.FindByEmailResponse{}
	if err := repo.db.Pool.QueryRow(ctx, sql, email).Scan(&user.ID, &user.Email, &user.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return user, nil
		}
		return nil, fmt.Errorf("failed to get user: %s", err.Error())
	}
	return user, nil
}
