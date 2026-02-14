package services

import (
	"context"
	"fmt"
	"warehouse/models"

	"gorm.io/gorm"
)

type PembelianRepo interface {
	CreateBeliHeader(tx *gorm.DB, data models.BeliHeader) (*models.BeliHeader, error)
	CreateBeliDetail(tx *gorm.DB, data models.BeliDetail) (*models.BeliDetail, error)
	UpdateBeliHeader(tx *gorm.DB, data models.BeliHeader) error
	UpdateBeliDetail(tx *gorm.DB, data models.BeliDetail) error
	GetAllPembelian(ctx context.Context, offset, limit int) ([]models.BeliHeader, int64, error)
	GetPembelian(ctx context.Context, beliID int) (*models.Pembelian, error)
}

type PembelianService struct {
	DB            *gorm.DB
	pembelianRepo PembelianRepo
	barangRepo    BarangRepo
	stokRepo      StokRepo
	userRepo      UserRepo
}

func NewPembelianService(db *gorm.DB, pembelianRepo PembelianRepo, barangRepo BarangRepo, stokRepo StokRepo, userRepo UserRepo) *PembelianService {
	return &PembelianService{
		DB:            db,
		pembelianRepo: pembelianRepo,
		barangRepo:    barangRepo,
		stokRepo:      stokRepo,
		userRepo:      userRepo,
	}
}

func (s *PembelianService) CreatePembelian(ctx context.Context, beliHeaderData models.BeliHeader, beliDetailDataArr []models.BeliDetail) (*models.Pembelian, error) {
	var pembelian *models.Pembelian
	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		beliHeader, err := s.pembelianRepo.CreateBeliHeader(tx, beliHeaderData)
		if err != nil {
			return err
		}

		nomorFaktur := fmt.Sprintf("BLI%03d", beliHeader.ID)

		var total float64
		var beliDetailResponse []models.BeliDetailWithBarang
		for _, detail := range beliDetailDataArr {
			barang, err := s.barangRepo.GetBarang(tx, detail.BarangID)
			if err != nil {
				return err
			}

			beliDetailData := models.BeliDetail{
				BeliHeaderID: beliHeader.ID,
				BarangID:     detail.BarangID,
				Qty:          detail.Qty,
				Harga:        float64(barang.HargaBeli),
				Subtotal:     float64(detail.Qty) * float64(barang.HargaBeli),
			}
			beliDetail, err := s.pembelianRepo.CreateBeliDetail(tx, beliDetailData)
			if err != nil {
				return err
			}

			beliDetailResponse = append(beliDetailResponse, models.BeliDetailWithBarang{
				ID:           beliDetail.ID,
				BeliHeaderID: beliDetail.BeliHeaderID,
				BarangID:     beliDetail.BarangID,
				Qty:          beliDetail.Qty,
				Harga:        beliDetail.Harga,
				Subtotal:     beliDetail.Subtotal,
				Barang: models.Barang{
					ID:         barang.ID,
					KodeBarang: barang.KodeBarang,
					NamaBarang: barang.NamaBarang,
					Satuan:     barang.Satuan,
					HargaJual:  barang.HargaJual,
				},
			})

			stok, err := s.stokRepo.GetStokByBarangID(tx, barang.ID)
			if err != nil {
				return err
			}

			addStok := stok.StokAkhir + detail.Qty
			err = s.stokRepo.UpdateStok(tx, barang.ID, addStok)
			if err != nil {
				return err
			}

			historyStokData := models.HistoryStok{
				BarangID:       barang.ID,
				UserID:         beliHeaderData.UserID,
				JenisTransaksi: "masuk",
				Jumlah:         detail.Qty,
				StokSebelum:    stok.StokAkhir,
				StokSesudah:    addStok,
				Keterangan:     fmt.Sprintf("Pembelian %s", nomorFaktur),
			}
			_, err = s.stokRepo.CreateHistoryStok(tx, historyStokData)
			if err != nil {
				return err
			}

			total += beliDetailData.Subtotal
		}

		beliHeader.NoFaktur = nomorFaktur
		beliHeader.Total = total
		beliHeader.Status = "selesai"
		err = s.pembelianRepo.UpdateBeliHeader(tx, *beliHeader)
		if err != nil {
			return err
		}

		user, err := s.userRepo.FindUserByID(tx, beliHeader.UserID)
		if err != nil {
			return err
		}

		pembelian = &models.Pembelian{
			ID:        beliHeader.ID,
			NoFaktur:  beliHeader.NoFaktur,
			Supplier:  beliHeader.Supplier,
			Total:     beliHeader.Total,
			UserID:    beliHeader.UserID,
			Status:    beliHeader.Status,
			CreatedAt: beliHeader.CreatedAt,
			User: models.User{
				ID:       user.ID,
				Username: user.Username,
				FullName: user.FullName,
			},
			BeliDetail: beliDetailResponse,
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return pembelian, nil
}

func (s *PembelianService) GetAllPembelian(ctx context.Context, page, limit int) ([]models.BeliHeader, int64, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 5
	}
	offset := (page - 1) * limit
	return s.pembelianRepo.GetAllPembelian(ctx, offset, limit)
}

func (s *PembelianService) GetPembelian(ctx context.Context, beliID int) (*models.Pembelian, error) {
	return s.pembelianRepo.GetPembelian(ctx, beliID)
}
