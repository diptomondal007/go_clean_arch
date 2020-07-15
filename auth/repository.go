package auth

import (
	"context"
	"github.com/diptomondal007/go_clean_arch/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username string)(*models.User, error)
}
