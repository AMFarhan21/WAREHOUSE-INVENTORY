package handlers

import (
	"context"
	"log"
	"net/http"
	"strings"
	"warehouse/config/response"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserService interface {
	Register(ctx context.Context, data models.Users) (*models.Users, error)
	Login(ctx context.Context, email, password string) (*string, error)
}

type UserHandler struct {
	userService UserService
	validate    *validator.Validate
}

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=4"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required,min=4"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validate:    validator.New(),
	}
}

func (h *UserHandler) RegisterUser(g *gin.Context) {
	var request RegisterUserRequest

	if err := g.Bind(&request); err != nil {
		g.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success:   false,
			Message:   "Data input tidak valid",
			ErrorCode: response.ValidationError,
		})
		log.Printf("Error on RegisterUser input request: %v", err.Error())
		return
	}

	if err := h.validate.Struct(request); err != nil {
		g.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success:   false,
			Message:   "Data input tidak valid",
			ErrorCode: response.ValidationError,
		})
		log.Printf("Error on RegisterUser validation: %v", err.Error())
		return
	}

	user, err := h.userService.Register(g.Request.Context(), models.Users{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		FullName: request.FullName,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			g.JSON(http.StatusBadRequest, response.ErrorResponse{
				Success:   false,
				Message:   "User already exists",
				ErrorCode: response.ValidationError,
			})
			log.Printf("Error on RegisterUser validation: %v", err.Error())
			return
		}

		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})
		log.Printf("Error on RegisterUser internal: %v", err.Error())
		return
	}

	g.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	})
}

func (h *UserHandler) Login(g *gin.Context) {
	var request LoginRequest

	if err := g.Bind(&request); err != nil {
		g.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success:   false,
			Message:   "Data input tidak valid",
			ErrorCode: response.ValidationError,
		})
		log.Printf("Error on Login input request: %v", err.Error())
		return
	}

	if err := h.validate.Struct(request); err != nil {
		g.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success:   false,
			Message:   "Data input tidak valid",
			ErrorCode: response.ValidationError,
		})
		log.Printf("Error on Login validation: %v", err.Error())
		return
	}

	token, err := h.userService.Login(g.Request.Context(), request.Email, request.Password)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "not the hash of the given password") {
			g.JSON(http.StatusBadRequest, response.ErrorResponse{
				Success:   false,
				Message:   "Email or Password is invalid",
				ErrorCode: response.ValidationError,
			})
			log.Printf("Error on Login validation: %v", err.Error())
			return
		}

		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})
		log.Printf("Error on Login internal: %v", err.Error())
		return
	}

	g.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "User login successfully",
		Data:    token,
	})
}
