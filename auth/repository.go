package auth

import (
	"context"
	"myProject/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (uint16, error)
	GetUser(ctx context.Context, username, password string) (*models.User, error)
}
