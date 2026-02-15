package services

import (
	"context"
	"errors"
	"fmt"
	"warehouse/models"

	"gorm.io/gorm"
)

type PenjualanRepo interface {
	CreateJualHeader(tx *gorm.DB, data models.JualHeader) (*models.JualHeader, error)
	CreateJualDetail(tx *gorm.DB, data models.JualDetail) (*models.JualDetail, error)
	UpdateJualHeader(tx *gorm.DB, data models.JualHeader) error
	UpdateJualDetail(tx *gorm.DB, data models.JualDetail) error
	GetAllPenjualan(ctx context.Context, offset, limit int) ([]models.JualHeader, int64, error)
	GetPenjualan(ctx context.Context, jualID int) (*models.Penjualan, error)
	GetPenjualanByDate(ctx context.Context, startDate, endDate string, offset, limit int) ([]models.Penjualan, int64, error)
}

type PenjualanService struct {
	DB            *gorm.DB
	penjualanRepo PenjualanRepo
	stokRepo      StokRepo
	barangRepo    BarangRepo
	userRepo      UserRepo
}

func NewPenjualanService(db *gorm.DB, penjualanRepo PenjualanRepo, stokRepo StokRepo, barangRepo BarangRepo, userRepo UserRepo) *PenjualanService {
	return &PenjualanService{
		DB:            db,
		penjualanRepo: penjualanRepo,
		stokRepo:      stokRepo,
		barangRepo:    barangRepo,
		userRepo:      userRepo,
	}
}

func (s *PenjualanService) CreatePenjualan(ctx context.Context, dataJualHeader models.JualHeader, dataJualDetail []models.JualDetail) (*models.Penjualan, error) {
	var penjualan *models.Penjualan

	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		jualHeader, err := s.penjualanRepo.CreateJualHeader(tx, dataJualHeader)
		if err != nil {
			return err
		}

		nomorFaktur := fmt.Sprintf("JUAL%03d", jualHeader.ID)

		var jualDetailResponse []models.JualDetailWithBarang
		var total float64
		for _, item := range dataJualDetail {
			barang, err := s.barangRepo.GetBarang(tx, item.BarangID)
			if err != nil {
				return err
			}

			mstok, err := s.stokRepo.LockStok(tx, item.BarangID)
			if err != nil {
				return err
			}

			if mstok.StokAkhir < item.Qty {
				return errors.New("Insufficient stock")
			}

			jualDetailData := models.JualDetail{
				JualHeaderID: jualHeader.ID,
				BarangID:     item.BarangID,
				Qty:          item.Qty,
				Harga:        barang.HargaJual,
				Subtotal:     float64(item.Qty) * barang.HargaJual,
			}

			stokSebelum := mstok.StokAkhir
			stokSesudah := stokSebelum - item.Qty
			// err = s.stokRepo.UpdateStok(tx, item.BarangID, substractStok)
			// if err != nil {
			// 	return err
			// }
			mstok.StokAkhir = stokSesudah
			err = tx.Save(mstok).Error
			if err != nil {
				return err
			}

			jualDetail, err := s.penjualanRepo.CreateJualDetail(tx, jualDetailData)
			if err != nil {
				return err
			}

			jualDetailResponse = append(jualDetailResponse, models.JualDetailWithBarang{
				ID:           jualDetail.ID,
				JualHeaderID: jualDetail.JualHeaderID,
				BarangID:     jualDetail.BarangID,
				Qty:          jualDetail.Qty,
				Harga:        jualDetail.Harga,
				Subtotal:     jualDetail.Subtotal,
				Barang: models.Barang{
					ID:         barang.ID,
					KodeBarang: barang.KodeBarang,
					NamaBarang: barang.NamaBarang,
					Satuan:     barang.Satuan,
					HargaJual:  barang.HargaJual,
				},
			})

			historyStokData := models.HistoryStok{
				BarangID:       item.BarangID,
				UserID:         dataJualHeader.UserID,
				JenisTransaksi: "keluar",
				Jumlah:         item.Qty,
				StokSebelum:    stokSebelum,
				StokSesudah:    stokSesudah,
				Keterangan:     fmt.Sprintf("Penjualan %s", nomorFaktur),
			}

			_, err = s.stokRepo.CreateHistoryStok(tx, historyStokData)
			if err != nil {
				return err
			}

			total += jualDetailData.Subtotal
		}

		jualHeader.NoFaktur = nomorFaktur
		jualHeader.Total = total
		jualHeader.Status = "selesai"
		err = s.penjualanRepo.UpdateJualHeader(tx, *jualHeader)
		if err != nil {
			return err
		}

		user, err := s.userRepo.FindUserByID(tx, jualHeader.UserID)
		if err != nil {
			return err
		}

		penjualan = &models.Penjualan{
			ID:        jualHeader.ID,
			NoFaktur:  nomorFaktur,
			Customer:  jualHeader.Customer,
			Total:     total,
			UserID:    jualHeader.UserID,
			Status:    jualHeader.Status,
			CreatedAt: jualHeader.CreatedAt,
			User: models.User{
				ID:       user.ID,
				Username: user.Username,
				FullName: user.FullName,
			},
			JualDetail: jualDetailResponse,
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return penjualan, nil
}

func (s *PenjualanService) GetAllPenjualan(ctx context.Context, page, limit int) ([]models.JualHeader, int64, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 5
	}

	offset := (page - 1) * limit

	return s.penjualanRepo.GetAllPenjualan(ctx, offset, limit)
}

func (s *PenjualanService) GetPenjualan(ctx context.Context, jualID int) (*models.Penjualan, error) {
	return s.penjualanRepo.GetPenjualan(ctx, jualID)
}

func (s *PenjualanService) GetPenjualanByDate(ctx context.Context, startDate, endDate string, page, limit int) ([]models.Penjualan, int64, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 1
	}

	offset := (page - 1) * limit

	return s.penjualanRepo.GetPenjualanByDate(ctx, startDate, endDate, offset, limit)
}
