package auth

import (
	"context"
	"github.com/diptomondal007/go_clean_arch/models"
)

//UseCase interface is to hold the use cases for auth
type UseCase interface {
	SignUp(ctx context.Context, user *models.User) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
