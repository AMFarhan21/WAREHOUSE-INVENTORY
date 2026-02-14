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

func TestCreatePembelian(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPembelianRepo := mock_services.NewMockPembelianRepo(ctrl)
	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockStokRepo := mock_services.NewMockStokRepo(ctrl)
	mockUserRepo := mock_services.NewMockUserRepo(ctrl)

	db := testDB(t)

	service := services.NewPembelianService(
		db,
		mockPembelianRepo,
		mockBarangRepo,
		mockStokRepo,
		mockUserRepo,
	)

	headerInput := models.BeliHeader{
		Supplier: "PT Supplier",
		UserID:   1,
	}

	detailInput := []models.BeliDetail{
		{
			BarangID: 1,
			Qty:      2,
		},
	}

	createdHeader := &models.BeliHeader{
		ID:        1,
		Supplier:  "PT Supplier",
		UserID:    1,
		CreatedAt: time.Now(),
	}

	mockPembelianRepo.
		EXPECT().
		CreateBeliHeader(gomock.Any(), gomock.Any()).
		Return(createdHeader, nil)

	mockBarangRepo.
		EXPECT().
		GetBarang(gomock.Any(), 1).
		Return(&models.MasterBarang{
			ID:         1,
			KodeBarang: "BRG001",
			NamaBarang: "Laptop",
			Satuan:     "unit",
			HargaBeli:  1000,
			HargaJual:  1500,
		}, nil)

	mockPembelianRepo.
		EXPECT().
		CreateBeliDetail(gomock.Any(), gomock.Any()).
		Return(&models.BeliDetail{
			ID:           1,
			BeliHeaderID: 1,
			BarangID:     1,
			Qty:          2,
			Harga:        1000,
			Subtotal:     2000,
		}, nil)

	mockStokRepo.
		EXPECT().
		GetStokByBarangID(gomock.Any(), 1).
		Return(&models.Mstok{
			ID:        1,
			BarangID:  1,
			StokAkhir: 10,
		}, nil)

	mockStokRepo.
		EXPECT().
		UpdateStok(gomock.Any(), 1, 12).
		Return(nil)

	mockStokRepo.
		EXPECT().
		CreateHistoryStok(gomock.Any(), gomock.Any()).
		Return(&models.HistoryStok{}, nil)

	mockPembelianRepo.
		EXPECT().
		UpdateBeliHeader(gomock.Any(), gomock.Any()).
		Return(nil)

	mockUserRepo.
		EXPECT().
		FindUserByID(gomock.Any(), 1).
		Return(&models.Users{
			ID:       1,
			Username: "admin",
			FullName: "Administrator",
		}, nil)

	result, err := service.CreatePembelian(context.Background(), headerInput, detailInput)
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

func TestGetAllPembelian(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPembelianRepo := mock_services.NewMockPembelianRepo(ctrl)
	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockStokRepo := mock_services.NewMockStokRepo(ctrl)
	mockUserRepo := mock_services.NewMockUserRepo(ctrl)

	db := testDB(t)

	service := services.NewPembelianService(
		db,
		mockPembelianRepo,
		mockBarangRepo,
		mockStokRepo,
		mockUserRepo,
	)

	expected := []models.BeliHeader{
		{ID: 1, Supplier: "PT Supplier"},
	}

	mockPembelianRepo.
		EXPECT().
		GetAllPembelian(gomock.Any(), 0, 5).
		Return(expected, int64(1), nil)

	data, total, err := service.GetAllPembelian(context.Background(), 1, 5)
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

func TestGetPembelian(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPembelianRepo := mock_services.NewMockPembelianRepo(ctrl)
	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockStokRepo := mock_services.NewMockStokRepo(ctrl)
	mockUserRepo := mock_services.NewMockUserRepo(ctrl)

	db := testDB(t)

	service := services.NewPembelianService(
		db,
		mockPembelianRepo,
		mockBarangRepo,
		mockStokRepo,
		mockUserRepo,
	)

	expected := &models.Pembelian{
		ID:       1,
		NoFaktur: "BLI001",
	}

	mockPembelianRepo.
		EXPECT().
		GetPembelian(gomock.Any(), 1).
		Return(expected, nil)

	result, err := service.GetPembelian(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != 1 {
		t.Fatal("unexpected result")
	}
}
