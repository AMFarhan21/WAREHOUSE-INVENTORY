package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"warehouse/config/response"
	"warehouse/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BarangService interface {
	GetAllBarang(ctx context.Context, search string, page, limit int) ([]models.MasterBarang, int64, error)
	CreateBarang(ctx context.Context, data models.MasterBarang) (*models.MasterBarang, error)
	GetBarang(ctx context.Context, id int) (*models.MasterBarang, error)
	UpdateBarang(ctx context.Context, data models.MasterBarang) error
	DeleteBarang(ctx context.Context, id int) error
	GetAllBarangWithStok(ctx context.Context, page, limit int) ([]models.MasterBarangWithStok, int64, error)
}

type BarangHandler struct {
	barangService BarangService
	validate      *validator.Validate
}

type inputBarang struct {
	NamaBarang string  `json:"nama_barang" validate:"required"`
	Deskripsi  *string `json:"deskripsi"`
	Satuan     string  `json:"satuan" validate:"required"`
	HargaBeli  float64 `json:"harga_beli"`
	HargaJual  float64 `json:"harga_jual"`
}

type MasterBarangResponse struct {
	ID         int       `json:"id" gorm:"autoIncrement:true"`
	KodeBarang string    `json:"kode_barang"`
	NamaBarang string    `json:"nama_barang"`
	Deskripsi  *string   `json:"deskripsi"`
	Satuan     string    `json:"satuan"`
	HargaBeli  float64   `json:"harga_beli"`
	HargaJual  float64   `json:"harga_jual"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewBarangHandler(barangService BarangService) *BarangHandler {
	return &BarangHandler{
		barangService: barangService,
		validate:      validator.New(),
	}
}

func (h *BarangHandler) GetAllBarang(g *gin.Context) {
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "5"))
	search := g.Query("search")

	barang, total, err := h.barangService.GetAllBarang(g.Request.Context(), search, page, limit)
	if err != nil {
		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})

		log.Printf("Error on getallbarang handler: %v", err.Error())

		return
	}

	var masterBarangResponse []MasterBarangResponse
	for _, item := range barang {
		masterBarangResponse = append(masterBarangResponse, MasterBarangResponse{
			ID:         item.ID,
			KodeBarang: item.KodeBarang,
			NamaBarang: item.NamaBarang,
			Deskripsi:  item.Deskripsi,
			Satuan:     item.Satuan,
			HargaBeli:  item.HargaBeli,
			HargaJual:  item.HargaJual,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		})
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data:    masterBarangResponse,
		Meta: &response.Meta{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
}

func (h *BarangHandler) CreateBarang(g *gin.Context) {
	var request inputBarang

	if err := g.Bind(&request); err != nil {
		g.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Success:   false,
			Message:   "Data input tidak valid",
			ErrorCode: response.ValidationError,
		})

		log.Printf("Error on CreateBarang Handler request input: %v", err.Error())
		return
	}

	if err := h.validate.Struct(request); err != nil {
		g.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Success:   false,
			Message:   "Data input tidak valid",
			ErrorCode: response.ValidationError,
		})

		log.Printf("Error on CreateBarang Handler validation: %v", err.Error())
		return
	}

	barang, err := h.barangService.CreateBarang(g.Request.Context(), models.MasterBarang{
		NamaBarang: request.NamaBarang,
		Deskripsi:  request.Deskripsi,
		Satuan:     request.Satuan,
		HargaBeli:  request.HargaBeli,
		HargaJual:  request.HargaJual,
	})
	if err != nil {
		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})

		log.Printf("Error on CreateBarang Handler validation: %v", err.Error())
		return
	}

	g.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Data created successfully",
		Data: MasterBarangResponse{
			ID:         barang.ID,
			KodeBarang: barang.KodeBarang,
			NamaBarang: barang.NamaBarang,
			Deskripsi:  barang.Deskripsi,
			Satuan:     barang.Satuan,
			HargaBeli:  barang.HargaBeli,
			HargaJual:  barang.HargaJual,
			CreatedAt:  barang.CreatedAt,
			UpdatedAt:  barang.UpdatedAt,
		},
	})
}

func (h *BarangHandler) GetBarang(g *gin.Context) {
	barangID, _ := strconv.Atoi(g.Param("barangID"))
	barang, err := h.barangService.GetBarang(g.Request.Context(), barangID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			g.JSON(http.StatusNotFound, response.ErrorResponse{
				Success:   false,
				Message:   "Item not found",
				ErrorCode: response.ItemNotFound,
			})

			log.Printf("Error on GetBarang Handler param: %v", err.Error())
			return
		}

		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})

		log.Printf("Error on GetBarang handler internal: %v", err.Error())

		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data: MasterBarangResponse{
			ID:         barang.ID,
			KodeBarang: barang.KodeBarang,
			NamaBarang: barang.NamaBarang,
			Deskripsi:  barang.Deskripsi,
			Satuan:     barang.Satuan,
			HargaBeli:  barang.HargaBeli,
			HargaJual:  barang.HargaJual,
			CreatedAt:  barang.CreatedAt,
			UpdatedAt:  barang.UpdatedAt,
		},
	})
}

func (h *BarangHandler) DeleteBarang(g *gin.Context) {
	barangID, _ := strconv.Atoi(g.Param("barangID"))
	err := h.barangService.DeleteBarang(g.Request.Context(), barangID)
	if err != nil {
		if strings.Contains(err.Error(), "NOT_FOUND") {
			g.JSON(http.StatusNotFound, response.ErrorResponse{
				Success:   false,
				Message:   "Item not found",
				ErrorCode: response.ItemNotFound,
			})

			log.Printf("Error on DeleteBarang Handler param: %v", err.Error())
			return
		}

		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})

		log.Printf("Error on DeleteBarang handler: %v", err.Error())
		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data deleted successfully",
	})
}

func (h *BarangHandler) UpdateBarang(g *gin.Context) {
	barangID, _ := strconv.Atoi(g.Param("barangID"))

	var request inputBarang

	if err := g.Bind(&request); err != nil {
		g.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Success:   false,
			Message:   "Data input tidak valid",
			ErrorCode: response.ValidationError,
		})

		log.Printf("Error on UpdateBarang handler request input: %v", err.Error())
		return
	}

	if err := h.validate.Struct(request); err != nil {
		g.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Success:   false,
			Message:   "Data input tidak valid",
			ErrorCode: response.ValidationError,
		})

		log.Printf("Error on UpdateBarang handler validation: %v", err.Error())
		return
	}

	err := h.barangService.UpdateBarang(g.Request.Context(), models.MasterBarang{
		ID:         barangID,
		NamaBarang: request.NamaBarang,
		Deskripsi:  request.Deskripsi,
		Satuan:     request.Satuan,
		HargaBeli:  request.HargaBeli,
		HargaJual:  request.HargaJual,
	})
	if err != nil {
		if strings.Contains(err.Error(), "NOT_FOUND") {
			g.JSON(http.StatusNotFound, response.ErrorResponse{
				Success:   false,
				Message:   "Item not found",
				ErrorCode: response.ItemNotFound,
			})

			log.Printf("Error on UpdateBarang Handler param: %v", err.Error())
			return
		}

		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})

		log.Printf("Error on UpdateBarang Handler internal: %v", err.Error())
		return
	}
	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data updated successfully",
	})
}

func (h *BarangHandler) GetAllBarangWithStok(g *gin.Context) {
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "5"))

	barang, total, err := h.barangService.GetAllBarangWithStok(g.Request.Context(), page, limit)
	if err != nil {
		g.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "Server Error",
			ErrorCode: response.InternalError,
		})

		log.Printf("Error on GetAllBarangWithStock handler: %v", err.Error())

		return
	}

	g.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data:    barang,
		Meta: &response.Meta{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
}
