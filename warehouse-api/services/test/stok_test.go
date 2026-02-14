package services

import (
	"context"
	"testing"
	"warehouse/models"
	"warehouse/services"
	mock_services "warehouse/services/mocks"

	"github.com/golang/mock/gomock"
)

func TestGetAllStok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStokRepo := mock_services.NewMockStokRepo(ctrl)
	db := testDB(t)

	service := services.NewStokService(db, mockStokRepo)

	expected := []models.Mstok{
		{
			ID:        1,
			BarangID:  1,
			StokAkhir: 10,
		},
	}

	mockStokRepo.
		EXPECT().
		GetAllStok(gomock.Any(), 0, 5).
		Return(expected, int64(1), nil)

	data, total, err := service.GetAllStok(context.Background(), 1, 5)

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

func TestGetStokByBarangID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStokRepo := mock_services.NewMockStokRepo(ctrl)
	db := testDB(t)

	service := services.NewStokService(db, mockStokRepo)

	expected := &models.Mstok{
		ID:        1,
		BarangID:  1,
		StokAkhir: 10,
	}

	mockStokRepo.
		EXPECT().
		GetStokByBarangID(gomock.Any(), 1).
		Return(expected, nil)

	result, err := service.GetStokByBarangID(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.StokAkhir != 10 {
		t.Fatal("stok mismatch")
	}
}

func TestGetHistoryStok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStokRepo := mock_services.NewMockStokRepo(ctrl)
	db := testDB(t)

	service := services.NewStokService(db, mockStokRepo)

	expected := []models.HistoryStok{
		{
			ID:             1,
			BarangID:       1,
			JenisTransaksi: "masuk",
			Jumlah:         5,
			StokSebelum:    0,
			StokSesudah:    5,
		},
	}

	mockStokRepo.
		EXPECT().
		GetHistoryStok(gomock.Any(), 0, 5).
		Return(expected, int64(1), nil)

	data, total, err := service.GetHistoryStok(context.Background(), 1, 5)

	if err != nil {
		t.Fatal(err)
	}

	if total != 1 {
		t.Fatal("total mismatch")
	}

	if len(data) != 1 {
		t.Fatal("history mismatch")
	}
}

func TestGetHistoryStokByBarangID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStokRepo := mock_services.NewMockStokRepo(ctrl)
	db := testDB(t)

	service := services.NewStokService(db, mockStokRepo)

	expected := []models.HistoryStok{
		{
			ID:             1,
			BarangID:       1,
			JenisTransaksi: "keluar",
			Jumlah:         2,
		},
	}

	mockStokRepo.
		EXPECT().
		GetHistoryStokByBarangID(gomock.Any(), 1).
		Return(expected, nil)

	result, err := service.GetHistoryStokByBarangID(context.Background(), 1)

	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 1 {
		t.Fatal("barangID mismatch")
	}
}
