package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/diptomondal007/go_clean_arch/auth"
	"github.com/diptomondal007/go_clean_arch/models"
	"time"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type AuthUseCase struct {
	userRepo        auth.Repository
	hashSalt        string
	signingKey      []byte
	expiredDuration time.Duration
}

func NewAuthUseCase(repo auth.Repository, hashSalt string, signingKey []byte, tokenTTLSeconds time.Duration) auth.UseCase {
	return &AuthUseCase{
		userRepo:        repo,
		hashSalt:        hashSalt,
		signingKey:      signingKey,
		expiredDuration: time.Second * tokenTTLSeconds,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, user *models.User) error {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(a.hashSalt))

	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))
	return a.userRepo.CreateUser(ctx, user)
}

func (a *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	panic("implement me")
}

func (a *AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	panic("implement me")
}