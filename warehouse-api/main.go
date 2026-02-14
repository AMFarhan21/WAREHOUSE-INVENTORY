package main

import (
	"log"
	"warehouse/config"
	"warehouse/handlers"
	"warehouse/middleware"
	"warehouse/repositories"
	"warehouse/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db := config.GetDatabaseConnection(cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDb)

	// repositories
	barangRepo := repositories.NewBarangRepo(db)
	stokRepo := repositories.NewStokRepo(db)
	pembelianRepo := repositories.NewPembelianRepo(db)
	penjualanRepo := repositories.NewPenjualanRepo(db)
	userRepo := repositories.NewUserRepo(db)

	// services
	stokService := services.NewStokService(db, stokRepo)
	barangService := services.NewBarangService(db, barangRepo, stokRepo)
	pembelianService := services.NewPembelianService(db, pembelianRepo, barangRepo, stokRepo, userRepo)
	penjualanService := services.NewPenjualanService(db, penjualanRepo, stokRepo, barangRepo, userRepo)
	userService := services.NewUserService(userRepo, cfg.JwtSecret)

	// handlers
	barangHandler := handlers.NewBarangHandler(barangService)
	stokHandler := handlers.NewStokHandler(stokService)
	pembelianHandler := handlers.NewPembelianHandler(pembelianService)
	penjualanHandler := handlers.NewPenjualanHandler(penjualanService)
	userHandler := handlers.NewUserHandler(userService)

	app := gin.Default()

	// app.Use(cors.Default())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router(app, cfg.JwtSecret, barangHandler, stokHandler, pembelianHandler, penjualanHandler, userHandler)

	log.Println("Server running on port :" + cfg.Server)
	app.Run(":" + cfg.Server)
}

func router(
	app *gin.Engine,
	jwtSecret string,
	barangHandler *handlers.BarangHandler,
	stokHandler *handlers.StokHandler,
	pembelianHandler *handlers.PembelianHandler,
	penjualanHandler *handlers.PenjualanHandler,
	userHandler *handlers.UserHandler,
) {
	api := app.Group("/api")
	api.Use(middleware.JWTMiddleware(jwtSecret))

	// adminAccess := middleware.ACLMiddleware(map[string]bool{
	// 	"admin": true,
	// })
	adminAndStaffAccess := middleware.ACLMiddleware(map[string]bool{
		"admin": true,
		"staff": true,
	})

	barang := api.Group("/barang")
	barang.GET("", adminAndStaffAccess, barangHandler.GetAllBarang)
	barang.POST("", adminAndStaffAccess, barangHandler.CreateBarang)
	barang.GET("/:barangID", adminAndStaffAccess, barangHandler.GetBarang)
	barang.DELETE("/:barangID", adminAndStaffAccess, barangHandler.DeleteBarang)
	barang.PUT("/:barangID", adminAndStaffAccess, barangHandler.UpdateBarang)
	barang.GET("/stok", adminAndStaffAccess, barangHandler.GetAllBarangWithStok)

	stok := api.Group("/stok")
	stok.GET("", adminAndStaffAccess, stokHandler.GetAllStok)
	stok.GET("/:barangID", adminAndStaffAccess, stokHandler.GetStokByBarangID)

	historyStok := api.Group("/history-stok")
	historyStok.GET("", adminAndStaffAccess, stokHandler.GetHistoryStok)
	historyStok.GET("/:barangID", adminAndStaffAccess, stokHandler.GetHistoryStokByBarangID)

	pembelian := api.Group("/pembelian")
	pembelian.POST("", adminAndStaffAccess, pembelianHandler.CreatePembelian)
	pembelian.GET("", adminAndStaffAccess, pembelianHandler.GetAllPembelian)
	pembelian.GET("/:beliID", adminAndStaffAccess, pembelianHandler.GetPembelian)

	penjualan := api.Group("/penjualan")
	penjualan.POST("", adminAndStaffAccess, penjualanHandler.CreatePenjualan)
	penjualan.GET("", adminAndStaffAccess, penjualanHandler.GetAllPenjualan)
	penjualan.GET("/:jualID", adminAndStaffAccess, penjualanHandler.GetPenjualan)

	user := app.Group("/user")
	user.POST("/register", userHandler.RegisterUser)
	user.POST("/login", userHandler.Login)
}
