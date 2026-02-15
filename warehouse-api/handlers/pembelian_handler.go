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

type PembelianService interface {
	CreatePembelian(ctx context.Context, beliHeaderData models.BeliHeader, beliDetailDataArr []models.BeliDetail) (*models.Pembelian, error)
	GetAllPembelian(ctx context.Context, page, limit int) ([]models.BeliHeader, int64, error)
	GetPembelian(ctx context.Context, beliID int) (*models.Pembelian, error)
	GetPembelianByDate(ctx context.Context, startDate, endDate string, page, limit int) ([]models.Pembelian, int64, error)
}

type PembelianHandler struct {
	pembelianService PembelianService
	validate         *validator.Validate
}

type CreatePembelianRequest struct {
	Supplier   string                         `json:"supplier" validate:"required"`
	BeliDetail []CreatePembelianDetailRequest `json:"beli_detail" validate:"required,dive"`
}

type CreatePembelianDetailRequest struct {
	BarangID int `json:"barang_id" validate:"required"`
	Qty      int `json:"qty" validate:"required,gt=0"`
}

type PembelianReponse struct {
	Header  PembelianHeader               `json:"header"`
	Details []models.BeliDetailWithBarang `json:"details"`
}

type PembelianHeader struct {
	ID        int         `json:"id" gorm:"autoIncrement:true"`
	NoFaktur  string      `json:"no_faktur"`
	Supplier  string      `json:"supplier"`
	Total     float64     `json:"total"`
	UserID    int         `json:"user_id"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	User      models.User `json:"user"`
}

func NewPembelianHandler(pembelianService PembelianService) *PembelianHandler {
	return &PembelianHandler{
		pembelianService: pembelianService,
		validate:         validator.New(),
	}
}

func (h *PembelianHandler) CreatePembelian(g *gin.Context) {
	userIDVal, exists := g.Get("user_id")
	if !exists {
		middleware.Unauthorized(g)
		return
	}
	userID := userIDVal.(int)

	var beliHeaderRequest CreatePembelianRequest
	if err := g.Bind(&beliHeaderRequest); err != nil {
		g.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Success:   false,
			Message:   "Data input tidak valid",
			ErrorCode: response.ValidationError,
		})
		log.Printf("Error on CreatePembelian Handler beliHeaderRequest input: %v", err.Error())
		return
	}
	if err := h.validate.Struct(beliHeaderRequest); err != nil {
		g.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Success:   false,
			Message:   "Data input tidak valid",
			ErrorCode: response.ValidationError,
		})

		log.Printf("Error on CreatePembelian Handler beliHeaderRequest validation: %v", err.Error())
		return
	}

	beliHeader := models.BeliHeader{
		Supplier: beliHeaderRequest.Supplier,
		UserID:   userID,
	}

	var beliDetail []models.BeliDetail
	for _, item := range beliHeaderRequest.BeliDetail {
		beliDetail = append(beliDetail, models.BeliDetail{
			BarangID: item.BarangID,
			Qty:      item.Qty,
		})
	}

	pembelian, err := h.pembelianService.CreatePembelian(g.Request.Context(), beliHeader, beliDetail)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			g.JSON(http.StatusNotFound, response.ErrorResponse{
				Success:   false,
				Message:   "Item not found",
				ErrorCode: response.ItemNotFound,
			})
			log.Printf("Error on CreatePembelian request: %v", err.Error())
			return
		}

		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})
		log.Printf("Error on CreatePembelian Internal: %v", err.Error())
		return
	}

	g.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Berhasil membeli barang",
		Data: response.SuccessResponse{
			Success: true,
			Message: "Data retrieved successfully",
			Data: PembelianReponse{
				Header: PembelianHeader{
					ID:        pembelian.ID,
					NoFaktur:  pembelian.NoFaktur,
					Supplier:  pembelian.Supplier,
					Total:     pembelian.Total,
					UserID:    pembelian.UserID,
					Status:    pembelian.Status,
					CreatedAt: pembelian.CreatedAt,
					User:      pembelian.User,
				},
				Details: pembelian.BeliDetail,
			},
		},
	})
}

func (h *PembelianHandler) GetAllPembelian(g *gin.Context) {
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "5"))

	pembelian, total, err := h.pembelianService.GetAllPembelian(g.Request.Context(), page, limit)
	if err != nil {
		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})
		log.Printf("Error on GetAllPembelian internal server: %v", err.Error())
		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data:    pembelian,
		Meta: &response.Meta{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})

}

func (h *PembelianHandler) GetPembelian(g *gin.Context) {
	beliID, _ := strconv.Atoi(g.Param("beliID"))

	pembelian, err := h.pembelianService.GetPembelian(g.Request.Context(), beliID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			g.JSON(http.StatusNotFound, response.ErrorResponse{
				Success:   false,
				Message:   "Item not found",
				ErrorCode: response.ItemNotFound,
			})
			log.Printf("Error on GetPembelian request: %v", err.Error())
			return
		}

		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})
		log.Printf("Error on GetPembelian internal server: %v", err.Error())
		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data: PembelianReponse{
			Header: PembelianHeader{
				ID:        pembelian.ID,
				NoFaktur:  pembelian.NoFaktur,
				Supplier:  pembelian.Supplier,
				Total:     pembelian.Total,
				UserID:    pembelian.UserID,
				Status:    pembelian.Status,
				CreatedAt: pembelian.CreatedAt,
				User:      pembelian.User,
			},
			Details: pembelian.BeliDetail,
		},
	})

}

func (h *PembelianHandler) GetPembelianByDate(g *gin.Context) {
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "5"))

	startDate := g.Query("start_date")
	endDate := g.Query("end_date")

	pembelian, total, err := h.pembelianService.GetPembelianByDate(g.Request.Context(), startDate, endDate, page, limit)
	if err != nil {
		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})
		log.Printf("Error on GetPembelianByDate internal server: %v", err.Error())
		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data:    pembelian,
		Meta: &response.Meta{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
}
