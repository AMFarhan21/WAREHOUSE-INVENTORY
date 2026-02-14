package repositories

import (
	"context"
	"errors"
	"warehouse/config/response"
	"warehouse/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StokRepo struct {
	DB *gorm.DB
}

func NewStokRepo(DB *gorm.DB) *StokRepo {
	return &StokRepo{
		DB: DB,
	}
}

func (r *StokRepo) CreateStok(tx *gorm.DB, data models.Mstok) (*models.Mstok, error) {
	err := tx.Table("mstok").Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *StokRepo) GetAllStok(ctx context.Context, offset, limit int) ([]models.Mstok, int64, error) {
	var Stok []models.Mstok

	db := r.DB.Model(&models.Mstok{}).Preload("Barang").WithContext(ctx)
	err := db.WithContext(ctx).Offset(offset).Limit(limit).Order("id ASC").Find(&Stok).Error
	if err != nil {
		return nil, 0, nil
	}

	var total int64
	err = db.WithContext(ctx).Count(&total).Error
	if err != nil {
		return nil, 0, nil
	}

	return Stok, total, nil
}

func (r *StokRepo) GetStokByBarangID(tx *gorm.DB, barangID int) (*models.Mstok, error) {
	var Stok models.Mstok

	err := tx.Model(&models.Mstok{}).Preload("Barang").Where("barang_id=?", barangID).First(&Stok).Error
	if err != nil {
		return nil, err
	}

	return &Stok, nil
}

func (r *StokRepo) UpdateStok(tx *gorm.DB, barangID, stok int) error {
	row := tx.Table("mstok").Where("barang_id=?", barangID).Update("stok_akhir", stok)
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New(response.ItemNotFound)
	}

	return nil
}

func (r *StokRepo) GetHistoryStok(ctx context.Context, offset, limit int) ([]models.HistoryStok, int64, error) {
	var HistoryStok []models.HistoryStok

	db := r.DB.Model(&models.HistoryStok{}).Preload("User").Preload("Barang")

	err := db.WithContext(ctx).Offset(offset).Limit(limit).Order("id DESC").Find(&HistoryStok).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	err = db.WithContext(ctx).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return HistoryStok, total, nil
}

func (r *StokRepo) GetHistoryStokByBarangID(ctx context.Context, barangID int) ([]models.HistoryStok, error) {
	var HistoryStok []models.HistoryStok

	err := r.DB.Model(&models.HistoryStok{}).Preload("User").Preload("Barang").WithContext(ctx).Where("barang_id=?", barangID).Find(&HistoryStok).Error
	if err != nil {
		return nil, err
	}

	return HistoryStok, nil
}

func (r *StokRepo) CreateHistoryStok(tx *gorm.DB, data models.HistoryStok) (*models.HistoryStok, error) {
	err := tx.Table("history_stok").Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *StokRepo) LockStok(tx *gorm.DB, barangID int) (*models.Mstok, error) {
	var stok models.Mstok
	err := tx.Table("mstok").Clauses(clause.Locking{Strength: "UPDATE"}).Where("barang_id=?", barangID).First(&stok).Error
	if err != nil {
		return nil, err
	}

	return &stok, nil
}
