package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"warehouse/config/response"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type StokService interface {
	GetAllStok(ctx context.Context, page, limit int) ([]models.Mstok, int64, error)
	GetStokByBarangID(ctx context.Context, barangID int) (*models.Mstok, error)
	GetHistoryStok(ctx context.Context, page, limit int) ([]models.HistoryStok, int64, error)
	GetHistoryStokByBarangID(ctx context.Context, page, limit, barangID int) ([]models.HistoryStok, int64, error)
}

type StokHandler struct {
	stokService StokService
	validate    *validator.Validate
}

func NewStokHandler(stokService StokService) *StokHandler {
	return &StokHandler{
		stokService: stokService,
		validate:    validator.New(),
	}
}

func (h *StokHandler) GetAllStok(g *gin.Context) {
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "5"))

	stok, total, err := h.stokService.GetAllStok(g.Request.Context(), page, limit)
	if err != nil {
		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})

		log.Printf("Error on GetAllStok handler: %v", err.Error())

		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data:    stok,
		Meta: &response.Meta{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
}

func (h *StokHandler) GetStokByBarangID(g *gin.Context) {
	barangID, _ := strconv.Atoi(g.Param("barangID"))

	stok, err := h.stokService.GetStokByBarangID(g.Request.Context(), barangID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			g.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
				Success:   false,
				Message:   "Item not found",
				ErrorCode: response.ItemNotFound,
			})
			log.Printf("Error on GetStokByBarangID handler: %v", err.Error())
			return
		}
		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})

		log.Printf("Error on GetStokByBarangID handler internal: %v", err.Error())
		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data:    stok,
	})
}

func (h *StokHandler) GetHistoryStok(g *gin.Context) {
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "5"))

	historyStok, total, err := h.stokService.GetHistoryStok(g.Request.Context(), page, limit)
	if err != nil {
		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})

		log.Printf("Error on GetHistoryStok handler: %v", err.Error())

		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data:    historyStok,
		Meta: &response.Meta{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
}

func (h *StokHandler) GetHistoryStokByBarangID(g *gin.Context) {
	barangID, _ := strconv.Atoi(g.Param("barangID"))
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "5"))

	stok, total, err := h.stokService.GetHistoryStokByBarangID(g.Request.Context(), page, limit, barangID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			g.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
				Success:   false,
				Message:   "Item not found",
				ErrorCode: response.ItemNotFound,
			})
			log.Printf("Error on GetHistoryStokByBarangID handler: %v", err.Error())
			return
		}
		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})

		log.Printf("Error on GetHistoryStokByBarangID handler internal: %v", err.Error())
		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data:    stok,
		Meta: &response.Meta{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
}
