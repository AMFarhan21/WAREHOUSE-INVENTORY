package repositories

import (
	"context"
	"errors"
	"time"
	"warehouse/config/response"
	"warehouse/models"

	"gorm.io/gorm"
)

type BarangRepo struct {
	DB *gorm.DB
}

func NewBarangRepo(DB *gorm.DB) *BarangRepo {
	return &BarangRepo{
		DB: DB,
	}
}

func (r *BarangRepo) GetAllBarang(ctx context.Context, search string, offset, limit int) ([]models.MasterBarang, int64, error) {

	db := r.DB.Table("master_barang").WithContext(ctx).Where("deleted_at IS NULL")

	if search != "" {
		searchQuery := "%" + search + "%"
		db = db.Where("nama_barang ILIKE ? or kode_barang ILIKE ?", searchQuery, searchQuery)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var masterBarang []models.MasterBarang
	err = db.Offset(offset).Limit(limit).Order("id DESC").Find(&masterBarang).Error
	if err != nil {
		return nil, 0, err
	}

	return masterBarang, total, nil
}

func (r *BarangRepo) CreateBarang(tx *gorm.DB, data models.MasterBarang) (*models.MasterBarang, error) {
	err := tx.Table("master_barang").Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *BarangRepo) GetBarang(tx *gorm.DB, id int) (*models.MasterBarang, error) {
	var masterBarang models.MasterBarang

	err := tx.Table("master_barang").Where("id=? AND deleted_at IS NULL", id).First(&masterBarang).Error
	if err != nil {
		return nil, err
	}

	return &masterBarang, nil
}

func (r *BarangRepo) UpdateBarang(tx *gorm.DB, data models.MasterBarang) error {
	row := tx.Table("master_barang").Where("id=? AND deleted_at IS NULL", data.ID).Updates(data)
	err := row.Error
	if err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New(response.ItemNotFound)
	}

	return nil
}

func (r *BarangRepo) DeleteBarang(tx *gorm.DB, id int) error {
	// row := tx.Table("master_barang").Where("id=?", id).Delete(models.MasterBarang{})
	row := tx.Table("master_barang").Where("id=? AND deleted_at IS NULL", id).Update("deleted_at", time.Now())
	err := row.Error
	if err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New(response.ItemNotFound)
	}

	return nil
}

func (r *BarangRepo) GetAllBarangWithStok(ctx context.Context, offset, limit int) ([]models.MasterBarangWithStok, int64, error) {
	var masterBarangWithStok []models.MasterBarangWithStok

	db := r.DB.Model(&models.MasterBarang{}).Preload("Stok", "deleted_at IS NULL").WithContext(ctx).Where("deleted_at IS NULL")

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset(offset).Limit(limit).Find(&masterBarangWithStok).Error
	if err != nil {
		return nil, 0, err
	}

	return masterBarangWithStok, total, nil
}
