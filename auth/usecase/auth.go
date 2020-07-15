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
	user, err := a.userRepo.GetUser(ctx, username)
	if err != nil{
		return "", auth.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil{
		return "", auth.ErrPasswordDoesntMatch
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (a *AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	panic("implement me")
}