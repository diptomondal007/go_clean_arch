package mysql

import (
	"context"
	"github.com/diptomondal007/go_clean_arch/auth"
	"github.com/diptomondal007/go_clean_arch/models"
	"github.com/jinzhu/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB, table string) auth.Repository {
	return &AuthRepository{
		db: db.Table(table),
	}
}

func (a AuthRepository) CreateUser(ctx context.Context, user *models.User) error {
	dbc := a.db.Create(user)
	if dbc.Error != nil{
		return dbc.Error
	}
	return nil
}

func (a AuthRepository) GetUser(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := a.db.Find(&user, "username=?",username).Error
	if err != nil{
		return nil, err
	}
	return &user, nil
}
