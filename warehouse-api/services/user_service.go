package services

import (
	"context"
	"time"
	"warehouse/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(ctx context.Context, data models.Users) (*models.Users, error)
	FindUserByEmail(ctx context.Context, email string) (*models.Users, error)
	FindUserByID(tx *gorm.DB, userID int) (*models.Users, error)
}

type UserService struct {
	userRepo  UserRepo
	jwtSecret string
}

func NewUserService(userRepo UserRepo, jwtSecret string) *UserService {
	return &UserService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *UserService) Register(ctx context.Context, data models.Users) (*models.Users, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	data.Password = string(hashedPassword)
	data.Role = "staff"

	return s.userRepo.CreateUser(ctx, data)
}

func (s *UserService) Login(ctx context.Context, email, password string) (*string, error) {
	user, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	type MapClaims struct {
		UserID int    `json:"user_id"`
		Role   string `json:"role"`
		*jwt.RegisteredClaims
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MapClaims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: &jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	signedString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	return &signedString, nil
}
