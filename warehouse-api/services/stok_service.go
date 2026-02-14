package services

import (
	"context"
	"warehouse/models"

	"gorm.io/gorm"
)

type StokRepo interface {
	CreateStok(tx *gorm.DB, data models.Mstok) (*models.Mstok, error)
	GetAllStok(ctx context.Context, offset, limit int) ([]models.Mstok, int64, error)
	GetStokByBarangID(tx *gorm.DB, barangID int) (*models.Mstok, error)
	UpdateStok(tx *gorm.DB, barangID, stok int) error
	GetHistoryStok(ctx context.Context, offset, limit int) ([]models.HistoryStok, int64, error)
	GetHistoryStokByBarangID(ctx context.Context, barangID int) ([]models.HistoryStok, error)
	CreateHistoryStok(tx *gorm.DB, data models.HistoryStok) (*models.HistoryStok, error)
	LockStok(tx *gorm.DB, barangID int) (*models.Mstok, error)
	DeleteStok(tx *gorm.DB, barangID int) error
}

type StokService struct {
	DB       *gorm.DB
	stokRepo StokRepo
}

func NewStokService(db *gorm.DB, stokRepo StokRepo) *StokService {
	return &StokService{
		DB:       db,
		stokRepo: stokRepo,
	}
}

func (s *StokService) GetAllStok(ctx context.Context, page, limit int) ([]models.Mstok, int64, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 5
	}

	offset := (page - 1) * limit

	return s.stokRepo.GetAllStok(ctx, offset, limit)
}

func (s *StokService) GetStokByBarangID(ctx context.Context, barangID int) (*models.Mstok, error) {
	var mstok *models.Mstok
	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		stok, err := s.stokRepo.GetStokByBarangID(tx, barangID)
		if err != nil {
			return err
		}

		mstok = stok

		return nil
	})

	if err != nil {
		return nil, err
	}

	return mstok, nil
}

func (s *StokService) GetHistoryStok(ctx context.Context, page, limit int) ([]models.HistoryStok, int64, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 5
	}

	offset := (page - 1) * limit
	return s.stokRepo.GetHistoryStok(ctx, offset, limit)
}
func (s *StokService) GetHistoryStokByBarangID(ctx context.Context, barangID int) ([]models.HistoryStok, error) {
	return s.stokRepo.GetHistoryStokByBarangID(ctx, barangID)
}
