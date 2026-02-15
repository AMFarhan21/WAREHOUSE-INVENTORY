package repositories

import (
	"context"
	"errors"
	"warehouse/config/response"
	"warehouse/models"

	"gorm.io/gorm"
)

type PenjualanRepo struct {
	DB *gorm.DB
}

func NewPenjualanRepo(db *gorm.DB) *PenjualanRepo {
	return &PenjualanRepo{
		DB: db,
	}
}

func (r *PenjualanRepo) CreateJualHeader(tx *gorm.DB, data models.JualHeader) (*models.JualHeader, error) {
	err := tx.Table("jual_header").Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *PenjualanRepo) CreateJualDetail(tx *gorm.DB, data models.JualDetail) (*models.JualDetail, error) {
	err := tx.Table("jual_detail").Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *PenjualanRepo) UpdateJualHeader(tx *gorm.DB, data models.JualHeader) error {
	row := tx.Table("jual_header").Updates(&data)
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New(response.ItemNotFound)
	}

	return nil
}

func (r *PenjualanRepo) UpdateJualDetail(tx *gorm.DB, data models.JualDetail) error {
	row := tx.Table("jual_detail").Updates(&data)
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New(response.ItemNotFound)
	}

	return nil
}

func (r *PenjualanRepo) GetAllPenjualan(ctx context.Context, offset, limit int) ([]models.JualHeader, int64, error) {
	var jualHeader []models.JualHeader

	db := r.DB.Table("jual_header")

	err := db.WithContext(ctx).Offset(offset).Limit(limit).Order("id DESC").Find(&jualHeader).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	err = db.WithContext(ctx).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return jualHeader, total, nil
}

func (r *PenjualanRepo) GetPenjualan(ctx context.Context, jualID int) (*models.Penjualan, error) {
	var penjualan models.Penjualan

	err := r.DB.Model(&models.JualHeader{}).Preload("User").Preload("JualDetail").Preload("JualDetail.Barang").WithContext(ctx).Where("id=?", jualID).First(&penjualan).Error
	if err != nil {
		return nil, err
	}

	return &penjualan, nil
}

func (r *PenjualanRepo) GetPenjualanByDate(ctx context.Context, startDate, endDate string, offset, limit int) ([]models.Penjualan, int64, error) {
	var penjualan []models.Penjualan
	var total int64

	db := r.DB.Model(&models.JualHeader{}).Preload("User").Preload("JualDetail").Preload("JualDetail.Barang").WithContext(ctx)

	if startDate != "" && endDate != "" {
		fullStart := startDate + " 00:00:00"
		fullEnd := endDate + " 23:59:59"

		db = db.Where("created_at >= ? AND created_at <= ?", fullStart, fullEnd)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&penjualan).Error
	if err != nil {
		return nil, 0, err
	}

	return penjualan, total, nil
}
