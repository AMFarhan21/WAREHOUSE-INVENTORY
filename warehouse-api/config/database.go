package config

import (
	"fmt"
	"log"
	"warehouse/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection(host, port, user, password, db string) *gorm.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, db)

	dbConn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("Error on connecting to the database %v", err.Error())
	}

	err = dbConn.AutoMigrate(
		&models.Users{},
		&models.MasterBarang{},
		&models.Mstok{},
		&models.JualHeader{},
		&models.JualDetail{},
		&models.BeliHeader{},
		&models.BeliDetail{},
		&models.HistoryStok{},
	)

	log.Println("Successfully connected to the server")

	return dbConn
}
