package services

import (
	"context"
	"fmt"
	"warehouse/models"

	"gorm.io/gorm"
)

type BarangRepo interface {
	GetAllBarang(ctx context.Context, search string, offset, limit int) ([]models.MasterBarang, int64, error)
	GetBarang(tx *gorm.DB, id int) (*models.MasterBarang, error)
	CreateBarang(tx *gorm.DB, data models.MasterBarang) (*models.MasterBarang, error)
	UpdateBarang(tx *gorm.DB, data models.MasterBarang) error
	DeleteBarang(tx *gorm.DB, id int) error
	GetAllBarangWithStok(ctx context.Context, offset, limit int) ([]models.MasterBarangWithStok, int64, error)
}

type BarangService struct {
	DB         *gorm.DB
	barangRepo BarangRepo
	stokRepo   StokRepo
}

func NewBarangService(db *gorm.DB, barangrepo BarangRepo, stokRepo StokRepo) *BarangService {
	return &BarangService{
		DB:         db,
		barangRepo: barangrepo,
		stokRepo:   stokRepo,
	}
}

func (s *BarangService) GetAllBarang(ctx context.Context, search string, page, limit int) ([]models.MasterBarang, int64, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 5
	}

	offset := (page - 1) * limit

	return s.barangRepo.GetAllBarang(ctx, search, offset, limit)
}

func (s *BarangService) CreateBarang(ctx context.Context, data models.MasterBarang) (*models.MasterBarang, error) {
	var barangResponse *models.MasterBarang
	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		barang, err := s.barangRepo.CreateBarang(tx, data)
		if err != nil {
			return err
		}

		barang.KodeBarang = fmt.Sprintf("BRG%03d", barang.ID)
		err = s.barangRepo.UpdateBarang(tx, *barang)
		if err != nil {
			return err
		}

		stokData := models.Mstok{
			BarangID:  barang.ID,
			StokAkhir: 0,
		}

		_, err = s.stokRepo.CreateStok(tx, stokData)
		if err != nil {
			return err
		}

		barangResponse = barang

		return nil
	})
	if err != nil {
		return nil, err
	}

	return barangResponse, nil
}

func (s *BarangService) GetBarang(ctx context.Context, id int) (*models.MasterBarang, error) {
	var barang *models.MasterBarang
	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res, err := s.barangRepo.GetBarang(tx, id)
		if err != nil {
			return err
		}

		barang = res

		return nil
	})
	if err != nil {
		return nil, err
	}

	return barang, nil
}
func (s *BarangService) UpdateBarang(ctx context.Context, data models.MasterBarang) error {
	return s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return s.barangRepo.UpdateBarang(tx, data)
	})
}
func (s *BarangService) DeleteBarang(ctx context.Context, id int) error {
	return s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := s.stokRepo.DeleteStok(tx, id)
		if err != nil {
			return err
		}

		return s.barangRepo.DeleteBarang(tx, id)
	})
}
func (s *BarangService) GetAllBarangWithStok(ctx context.Context, page, limit int) ([]models.MasterBarangWithStok, int64, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 5
	}

	offset := (page - 1) * limit

	return s.barangRepo.GetAllBarangWithStok(ctx, offset, limit)
}
