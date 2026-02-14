package services

import (
	"context"
	"testing"
	"warehouse/models"
	"warehouse/services"
	mock_services "warehouse/services/mocks"

	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_services.NewMockUserRepo(ctrl)

	service := services.NewUserService(mockUserRepo, "secret")

	input := models.Users{
		Email:    "test@mail.com",
		Password: "123456",
	}

	mockUserRepo.
		EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, data models.Users) (*models.Users, error) {
			// pastikan password sudah ter-hash
			if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte("123456")); err != nil {
				t.Fatal("password not hashed correctly")
			}
			if data.Role != "staff" {
				t.Fatal("role should be staff")
			}

			return &data, nil
		})

	result, err := service.Register(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}

	if result.Email != input.Email {
		t.Fatal("email mismatch")
	}
}

func TestLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_services.NewMockUserRepo(ctrl)

	service := services.NewUserService(mockUserRepo, "secret")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	mockUserRepo.
		EXPECT().
		FindUserByEmail(gomock.Any(), "test@mail.com").
		Return(&models.Users{
			ID:       1,
			Email:    "test@mail.com",
			Password: string(hashedPassword),
			Role:     "staff",
		}, nil)

	token, err := service.Login(context.Background(), "test@mail.com", "123456")
	if err != nil {
		t.Fatal(err)
	}

	if token == nil || *token == "" {
		t.Fatal("token should not be empty")
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_services.NewMockUserRepo(ctrl)

	service := services.NewUserService(mockUserRepo, "secret")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correct"), bcrypt.DefaultCost)

	mockUserRepo.
		EXPECT().
		FindUserByEmail(gomock.Any(), "test@mail.com").
		Return(&models.Users{
			ID:       1,
			Email:    "test@mail.com",
			Password: string(hashedPassword),
			Role:     "staff",
		}, nil)

	_, err := service.Login(context.Background(), "test@mail.com", "wrong")

	if err == nil {
		t.Fatal("expected error for wrong password")
	}
}

func TestLogin_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_services.NewMockUserRepo(ctrl)

	service := services.NewUserService(mockUserRepo, "secret")

	mockUserRepo.
		EXPECT().
		FindUserByEmail(gomock.Any(), "notfound@mail.com").
		Return(nil, gorm.ErrRecordNotFound)

	_, err := service.Login(context.Background(), "notfound@mail.com", "123456")

	if err == nil {
		t.Fatal("expected error when user not found")
	}
}
