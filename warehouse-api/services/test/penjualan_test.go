package services

import (
	"context"
	"testing"
	"time"
	"warehouse/models"
	"warehouse/services"
	mock_services "warehouse/services/mocks"

	"github.com/golang/mock/gomock"
)

func TestCreatePenjualan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPenjualanRepo := mock_services.NewMockPenjualanRepo(ctrl)
	mockStokRepo := mock_services.NewMockStokRepo(ctrl)
	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockUserRepo := mock_services.NewMockUserRepo(ctrl)

	db := testDB(t)

	service := services.NewPenjualanService(
		db,
		mockPenjualanRepo,
		mockStokRepo,
		mockBarangRepo,
		mockUserRepo,
	)

	headerInput := models.JualHeader{
		Customer: "PT Customer",
		UserID:   1,
	}

	detailInput := []models.JualDetail{
		{
			BarangID: 1,
			Qty:      2,
		},
	}

	createdHeader := &models.JualHeader{
		ID:        1,
		Customer:  "PT Customer",
		UserID:    1,
		CreatedAt: time.Now(),
	}

	mockPenjualanRepo.
		EXPECT().
		CreateJualHeader(gomock.Any(), headerInput).
		Return(createdHeader, nil)

	mockBarangRepo.
		EXPECT().
		GetBarang(gomock.Any(), 1).
		Return(&models.MasterBarang{
			ID:         1,
			KodeBarang: "BRG001",
			NamaBarang: "Laptop",
			Satuan:     "unit",
			HargaJual:  1000,
		}, nil)

	mockStokRepo.
		EXPECT().
		LockStok(gomock.Any(), 1).
		Return(&models.Mstok{
			ID:        1,
			BarangID:  1,
			StokAkhir: 10,
		}, nil)

	mockPenjualanRepo.
		EXPECT().
		CreateJualDetail(gomock.Any(), gomock.Any()).
		Return(&models.JualDetail{
			ID:           1,
			JualHeaderID: 1,
			BarangID:     1,
			Qty:          2,
			Harga:        1000,
			Subtotal:     2000,
		}, nil)

	mockStokRepo.
		EXPECT().
		CreateHistoryStok(gomock.Any(), gomock.Any()).
		Return(&models.HistoryStok{}, nil)

	mockPenjualanRepo.
		EXPECT().
		UpdateJualHeader(gomock.Any(), gomock.Any()).
		Return(nil)

	mockUserRepo.
		EXPECT().
		FindUserByID(gomock.Any(), 1).
		Return(&models.Users{
			ID:       1,
			Username: "admin",
			FullName: "Administrator",
		}, nil)

	result, err := service.CreatePenjualan(context.Background(), headerInput, detailInput)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != 1 {
		t.Fatal("ID mismatch")
	}

	if result.Total != 2000 {
		t.Fatal("Total mismatch")
	}
}

func TestCreatePenjualan_InsufficientStock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPenjualanRepo := mock_services.NewMockPenjualanRepo(ctrl)
	mockStokRepo := mock_services.NewMockStokRepo(ctrl)
	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockUserRepo := mock_services.NewMockUserRepo(ctrl)

	db := testDB(t)

	service := services.NewPenjualanService(
		db,
		mockPenjualanRepo,
		mockStokRepo,
		mockBarangRepo,
		mockUserRepo,
	)

	headerInput := models.JualHeader{
		Customer: "PT Customer",
		UserID:   1,
	}

	detailInput := []models.JualDetail{
		{
			BarangID: 1,
			Qty:      20,
		},
	}

	mockPenjualanRepo.
		EXPECT().
		CreateJualHeader(gomock.Any(), headerInput).
		Return(&models.JualHeader{ID: 1, UserID: 1}, nil)

	mockBarangRepo.
		EXPECT().
		GetBarang(gomock.Any(), 1).
		Return(&models.MasterBarang{
			ID:        1,
			HargaJual: 1000,
		}, nil)

	mockStokRepo.
		EXPECT().
		LockStok(gomock.Any(), 1).
		Return(&models.Mstok{
			ID:        1,
			BarangID:  1,
			StokAkhir: 5,
		}, nil)

	_, err := service.CreatePenjualan(context.Background(), headerInput, detailInput)

	if err == nil {
		t.Fatal("expected insufficient stock error")
	}
}

func TestGetAllPenjualan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPenjualanRepo := mock_services.NewMockPenjualanRepo(ctrl)
	mockStokRepo := mock_services.NewMockStokRepo(ctrl)
	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockUserRepo := mock_services.NewMockUserRepo(ctrl)

	db := testDB(t)

	service := services.NewPenjualanService(
		db,
		mockPenjualanRepo,
		mockStokRepo,
		mockBarangRepo,
		mockUserRepo,
	)

	expected := []models.JualHeader{
		{ID: 1, Customer: "PT Customer"},
	}

	mockPenjualanRepo.
		EXPECT().
		GetAllPenjualan(gomock.Any(), 0, 5).
		Return(expected, int64(1), nil)

	data, total, err := service.GetAllPenjualan(context.Background(), 1, 5)
	if err != nil {
		t.Fatal(err)
	}

	if total != 1 {
		t.Fatal("total mismatch")
	}

	if len(data) != 1 {
		t.Fatal("data mismatch")
	}
}

func TestGetPenjualan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPenjualanRepo := mock_services.NewMockPenjualanRepo(ctrl)
	mockStokRepo := mock_services.NewMockStokRepo(ctrl)
	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockUserRepo := mock_services.NewMockUserRepo(ctrl)

	db := testDB(t)

	service := services.NewPenjualanService(
		db,
		mockPenjualanRepo,
		mockStokRepo,
		mockBarangRepo,
		mockUserRepo,
	)

	expected := &models.Penjualan{
		ID:       1,
		NoFaktur: "JUAL001",
	}

	mockPenjualanRepo.
		EXPECT().
		GetPenjualan(gomock.Any(), 1).
		Return(expected, nil)

	result, err := service.GetPenjualan(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != 1 {
		t.Fatal("unexpected result")
	}
}
