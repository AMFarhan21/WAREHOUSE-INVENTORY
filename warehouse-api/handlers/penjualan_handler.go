package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"warehouse/config/response"
	"warehouse/middleware"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PenjualanService interface {
	CreatePenjualan(ctx context.Context, dataJualHeader models.JualHeader, dataJualDetail []models.JualDetail) (*models.Penjualan, error)
	GetAllPenjualan(ctx context.Context, page, limit int) ([]models.JualHeader, int64, error)
	GetPenjualan(ctx context.Context, jualID int) (*models.Penjualan, error)
	GetPenjualanByDate(ctx context.Context, startDate, endDate string, page, limit int) ([]models.Penjualan, int64, error)
}

type PenjualanHandler struct {
	penjualanService PenjualanService
	validate         *validator.Validate
}

type CreatePenjualanRequest struct {
	Customer   string                         `json:"customer" validate:"required"`
	JualDetail []CreatePenjualanDetailRequest `json:"jual_detail" validate:"required,dive"`
}

type CreatePenjualanDetailRequest struct {
	BarangID int `json:"barang_id" validate:"required"`
	Qty      int `json:"qty" validate:"required,gt=0"`
}

type PenjualanReponse struct {
	Header  Header                        `json:"header"`
	Details []models.JualDetailWithBarang `json:"details"`
}

type Header struct {
	ID        int         `json:"id" gorm:"autoIncrement:true"`
	NoFaktur  string      `json:"no_faktur"`
	Customer  string      `json:"customer"`
	Total     float64     `json:"total"`
	UserID    int         `json:"user_id"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	User      models.User `json:"user"`
}

func NewPenjualanHandler(penjualanService PenjualanService) *PenjualanHandler {
	return &PenjualanHandler{
		penjualanService: penjualanService,
		validate:         validator.New(),
	}
}

func (h *PenjualanHandler) CreatePenjualan(g *gin.Context) {
	userIDVal, exists := g.Get("user_id")
	if !exists {
		middleware.Unauthorized(g)
		return
	}
	userID := userIDVal.(int)

	var jualHeaderRequest CreatePenjualanRequest
	if err := g.Bind(&jualHeaderRequest); err != nil {
		g.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Success:   false,
			Message:   "Data input tidak valid",
			ErrorCode: response.ValidationError,
		})
		log.Printf("Error on CreatePenjualan Handler jualHeaderRequest input: %v", err.Error())
		return
	}
	if err := h.validate.Struct(jualHeaderRequest); err != nil {
		g.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Success:   false,
			Message:   "Data input tidak valid",
			ErrorCode: response.ValidationError,
		})

		log.Printf("Error on CreatePenjualan Handler jualHeaderRequest validation: %v", err.Error())
		return
	}

	jualHeader := models.JualHeader{
		Customer: jualHeaderRequest.Customer,
		UserID:   userID,
	}

	var jualDetail []models.JualDetail
	for _, item := range jualHeaderRequest.JualDetail {
		jualDetail = append(jualDetail, models.JualDetail{
			BarangID: item.BarangID,
			Qty:      item.Qty,
		})
	}

	penjualan, err := h.penjualanService.CreatePenjualan(g.Request.Context(), jualHeader, jualDetail)
	if err != nil {
		if strings.Contains(err.Error(), "Insufficient") {
			g.JSON(http.StatusBadRequest, response.ErrorResponse{
				Success:   false,
				Message:   err.Error(),
				ErrorCode: response.InsufficientStock,
			})
			log.Printf("Error on CreatePenjualan stok: %v", err.Error())
			return
		}

		if strings.Contains(err.Error(), "not found") {
			g.JSON(http.StatusNotFound, response.ErrorResponse{
				Success:   false,
				Message:   "Item not found",
				ErrorCode: response.ItemNotFound,
			})
			log.Printf("Error on CreatePenjualan request: %v", err.Error())
			return
		}

		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})
		log.Printf("Error on CreatePenjualan Internal: %v", err.Error())
		return
	}

	g.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Berhasil memjual barang",
		Data: PenjualanReponse{
			Header: Header{
				ID:        penjualan.ID,
				NoFaktur:  penjualan.NoFaktur,
				Customer:  penjualan.Customer,
				Total:     penjualan.Total,
				UserID:    penjualan.UserID,
				Status:    penjualan.Status,
				CreatedAt: penjualan.CreatedAt,
				User:      penjualan.User,
			},
			Details: penjualan.JualDetail,
		},
	})
}

func (h *PenjualanHandler) GetAllPenjualan(g *gin.Context) {
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "5"))

	pemjualan, total, err := h.penjualanService.GetAllPenjualan(g.Request.Context(), page, limit)
	if err != nil {
		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})
		log.Printf("Error on GetAllPenjualan internal server: %v", err.Error())
		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data:    pemjualan,
		Meta: &response.Meta{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})

}

func (h *PenjualanHandler) GetPenjualan(g *gin.Context) {
	jualID, _ := strconv.Atoi(g.Param("jualID"))

	penjualan, err := h.penjualanService.GetPenjualan(g.Request.Context(), jualID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			g.JSON(http.StatusNotFound, response.ErrorResponse{
				Success:   false,
				Message:   "Item not found",
				ErrorCode: response.ItemNotFound,
			})
			log.Printf("Error on GetPenjualan request: %v", err.Error())
			return
		}

		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})
		log.Printf("Error on GetPenjualan internal server: %v", err.Error())
		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data: PenjualanReponse{
			Header: Header{
				ID:        penjualan.ID,
				NoFaktur:  penjualan.NoFaktur,
				Customer:  penjualan.Customer,
				Total:     penjualan.Total,
				UserID:    penjualan.UserID,
				Status:    penjualan.Status,
				CreatedAt: penjualan.CreatedAt,
				User:      penjualan.User,
			},
			Details: penjualan.JualDetail,
		},
	})

}

func (h *PenjualanHandler) GetPenjualanByDate(g *gin.Context) {
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "5"))

	startDate := g.Query("start_date")
	endDate := g.Query("end_date")

	penjualan, total, err := h.penjualanService.GetPenjualanByDate(g.Request.Context(), startDate, endDate, page, limit)
	if err != nil {
		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})
		log.Printf("Error on GetPenjualanByDate internal server: %v", err.Error())
		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data:    penjualan,
		Meta: &response.Meta{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
}
