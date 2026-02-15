package repositories

import (
	"context"
	"errors"
	"warehouse/config/response"
	"warehouse/models"

	"gorm.io/gorm"
)

type PembelianRepo struct {
	DB *gorm.DB
}

func NewPembelianRepo(db *gorm.DB) *PembelianRepo {
	return &PembelianRepo{
		DB: db,
	}
}

func (r *PembelianRepo) CreateBeliHeader(tx *gorm.DB, data models.BeliHeader) (*models.BeliHeader, error) {
	err := tx.Table("beli_header").Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *PembelianRepo) CreateBeliDetail(tx *gorm.DB, data models.BeliDetail) (*models.BeliDetail, error) {
	err := tx.Table("beli_detail").Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *PembelianRepo) UpdateBeliHeader(tx *gorm.DB, data models.BeliHeader) error {
	row := tx.Table("beli_header").Updates(&data)
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New(response.ItemNotFound)
	}

	return nil
}

func (r *PembelianRepo) UpdateBeliDetail(tx *gorm.DB, data models.BeliDetail) error {
	row := tx.Table("beli_detail").Updates(&data)
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New(response.ItemNotFound)
	}

	return nil
}

func (r *PembelianRepo) GetAllPembelian(ctx context.Context, offset, limit int) ([]models.BeliHeader, int64, error) {
	var pembelian []models.BeliHeader
	db := r.DB.Table("beli_header")
	err := db.WithContext(ctx).Offset(offset).Limit(limit).Order("id DESC").Find(&pembelian).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	err = db.WithContext(ctx).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return pembelian, total, nil
}

func (r *PembelianRepo) GetPembelian(ctx context.Context, beliID int) (*models.Pembelian, error) {
	var pembelian *models.Pembelian
	err := r.DB.Model(&models.BeliHeader{}).Preload("User").Preload("BeliDetail").Preload("BeliDetail.Barang").WithContext(ctx).Where("id=?", beliID).First(&pembelian).Error
	if err != nil {
		return nil, err
	}

	return pembelian, nil
}

func (r *PembelianRepo) GetPembelianByDate(ctx context.Context, startDate, endDate string, offset, limit int) ([]models.Pembelian, int64, error) {
	var pembelian []models.Pembelian
	var total int64

	db := r.DB.Model(&models.BeliHeader{}).Preload("User").Preload("BeliDetail").Preload("BeliDetail.Barang").WithContext(ctx)

	if startDate != "" && endDate != "" {
		fullStart := startDate + " 00:00:00"
		fullEnd := endDate + " 23:59:59"

		db = db.Where("created_at >= ? AND created_at <= ?", fullStart, fullEnd)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&pembelian).Error
	if err != nil {
		return nil, 0, err
	}

	return pembelian, total, nil
}
