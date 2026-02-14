package repositories

import (
	"context"
	"errors"
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

	db := r.DB.Table("master_barang").WithContext(ctx)

	if search != "" {
		searchQuery := "%" + search + "%"
		db = db.Where("nama_barang ILIKE ? or kode_barang ILIKE ?", searchQuery, searchQuery)
	}

	var masterBarang []models.MasterBarang
	err := db.WithContext(ctx).Offset(offset).Limit(limit).Find(&masterBarang).Error
	if err != nil {
		return nil, 0, nil
	}

	var total int64
	err = db.WithContext(ctx).Count(&total).Error
	if err != nil {
		return nil, 0, nil
	}

	return masterBarang, total, nil
}

func (r *BarangRepo) CreateBarang(tx *gorm.DB, data models.MasterBarang) (*models.MasterBarang, error) {
	err := tx.Table("master_barang").Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

	// return &data, r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
	// 	err := tx.Table("master_barang").WithContext(ctx).Create(&data).Error
	// 	if err != nil {
	// 		return err
	// 	}

	// 	data.KodeBarang = fmt.Sprintf("BRG%03d", data.ID)

	// 	err = tx.Table("master_barang").WithContext(ctx).Where("id=?", data.ID).Update("kode_barang", data.KodeBarang).Error
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil
	// })
}

func (r *BarangRepo) GetBarang(tx *gorm.DB, id int) (*models.MasterBarang, error) {
	var masterBarang models.MasterBarang

	err := tx.Table("master_barang").Where("id=?", id).First(&masterBarang).Error
	if err != nil {
		return nil, err
	}

	return &masterBarang, nil
}

func (r *BarangRepo) UpdateBarang(tx *gorm.DB, data models.MasterBarang) error {
	row := tx.Table("master_barang").Where("id=?", data.ID).Updates(data)
	err := row.Error
	if err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New(response.ItemNotFound)
	}

	return nil
}

func (r *BarangRepo) DeleteBarang(ctx context.Context, id int) error {
	row := r.DB.Table("master_barang").WithContext(ctx).Where("id=?", id).Delete(models.MasterBarang{})
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

	err := r.DB.Model(&models.MasterBarang{}).
		Preload("Stok").
		WithContext(ctx).Offset(offset).Limit(limit).Find(&masterBarangWithStok).Error
	if err != nil {
		return nil, 0, nil
	}

	var total int64
	err = r.DB.Table("master_barang").WithContext(ctx).Count(&total).Error
	if err != nil {
		return nil, 0, nil
	}

	return masterBarangWithStok, total, nil
}
