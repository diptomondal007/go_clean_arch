package usecase

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/diptomondal007/go_clean_arch/auth"
	"github.com/diptomondal007/go_clean_arch/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type AuthUseCase struct {
	userRepo        auth.Repository
	signingKey      []byte
	expiredDuration time.Duration
}

func NewAuthUseCase(repo auth.Repository, signingKey []byte, tokenTTLSeconds time.Duration) auth.UseCase {
	return &AuthUseCase{
		userRepo:        repo,
		signingKey:      signingKey,
		expiredDuration: time.Second * tokenTTLSeconds,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, user *models.User) error {
	passwordToByte := []byte(user.Password)
	passwordHash, _:= bcrypt.GenerateFromPassword(passwordToByte, bcrypt.MinCost)

	user.Password = string(passwordHash)
	return a.userRepo.CreateUser(ctx, user)
}

func (a *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	panic("implement me")
}

func (a *AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	panic("implement me")
}