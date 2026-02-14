package services

import (
	"context"
	"testing"
	"warehouse/models"
	"warehouse/services"

	mock_services "warehouse/services/mocks"

	"github.com/glebarez/sqlite"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func testDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	err = db.AutoMigrate(
		&models.Mstok{},
	)
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestGetAllBarang(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockServiceRepo := mock_services.NewMockStokRepo(ctrl)

	db := testDB(t)

	service := services.NewBarangService(db, mockBarangRepo, mockServiceRepo)

	expectedData := []models.MasterBarang{
		{
			ID: 1, NamaBarang: "Laptop",
		},
	}

	mockBarangRepo.EXPECT().GetAllBarang(gomock.Any(), "", 0, 5).Return(expectedData, int64(1), nil)

	data, total, err := service.GetAllBarang(context.Background(), "", 1, 5)

	if err != nil {
		t.Fatal(err)
	}

	if total != 1 {
		t.Fatal("total mismatch")
	}

	if len(data) != 1 {
		t.Fatal("data length mismatch")
	}
}

func TestCreateBarang(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockStokRepo := mock_services.NewMockStokRepo(ctrl)

	db := testDB(t)

	service := services.NewBarangService(db, mockBarangRepo, mockStokRepo)

	inputBarang := models.MasterBarang{
		NamaBarang: "Tes",
	}

	createdBarang := &models.MasterBarang{
		ID:         1,
		NamaBarang: "Tes",
	}

	mockBarangRepo.EXPECT().CreateBarang(gomock.Any(), inputBarang).Return(createdBarang, nil)
	mockBarangRepo.EXPECT().UpdateBarang(gomock.Any(), gomock.Any()).Return(nil)
	mockStokRepo.EXPECT().CreateStok(gomock.Any(), gomock.Any()).Return(&models.Mstok{}, nil)

	result, err := service.CreateBarang(context.Background(), inputBarang)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != 1 {
		t.Fatal("ID mismatch")
	}

	if result.KodeBarang != "BRG001" {
		t.Fatal("KodeBarang tidak sesuai")
	}

}

func TestGetBarang(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockStokRepo := mock_services.NewMockStokRepo(ctrl)

	db := testDB(t)

	service := services.NewBarangService(db, mockBarangRepo, mockStokRepo)

	expected := &models.MasterBarang{
		ID:         1,
		NamaBarang: "Koper",
	}

	mockBarangRepo.EXPECT().GetBarang(gomock.Any(), 1).Return(expected, nil)

	result, err := service.GetBarang(context.Background(), 1)

	if err != nil {
		t.Fatal(err)
	}

	if result.ID != 1 {
		t.Fatal("Unexpected result")
	}
}

func TestUpdateBarang(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockStokRepo := mock_services.NewMockStokRepo(ctrl)

	db := testDB(t)

	service := services.NewBarangService(db, mockBarangRepo, mockStokRepo)

	input := models.MasterBarang{
		ID:         1,
		NamaBarang: "Laptop Update",
	}

	mockBarangRepo.
		EXPECT().
		UpdateBarang(gomock.Any(), input).
		Return(nil)

	err := service.UpdateBarang(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteBarang(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockStokRepo := mock_services.NewMockStokRepo(ctrl)

	db := testDB(t)

	service := services.NewBarangService(db, mockBarangRepo, mockStokRepo)

	mockBarangRepo.
		EXPECT().
		DeleteBarang(gomock.Any(), 1).
		Return(nil)

	err := service.DeleteBarang(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAllBarangWithStok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBarangRepo := mock_services.NewMockBarangRepo(ctrl)
	mockStokRepo := mock_services.NewMockStokRepo(ctrl)

	db := testDB(t)

	service := services.NewBarangService(db, mockBarangRepo, mockStokRepo)

	expected := []models.MasterBarangWithStok{
		{
			ID:         1,
			NamaBarang: "Laptop",
			KodeBarang: "BRG001",
		},
	}

	mockBarangRepo.
		EXPECT().
		GetAllBarangWithStok(gomock.Any(), 0, 5).
		Return(expected, int64(1), nil)

	data, total, err := service.GetAllBarangWithStok(context.Background(), 1, 5)

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
