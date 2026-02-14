package repositories

import (
	"context"
	"warehouse/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, data models.Users) (*models.Users, error) {
	err := r.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *UserRepo) FindUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	var user models.Users
	err := r.DB.WithContext(ctx).Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) FindUserByID(tx *gorm.DB, userID int) (*models.Users, error) {
	var user models.Users
	err := tx.Table("users").Where("id=?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
