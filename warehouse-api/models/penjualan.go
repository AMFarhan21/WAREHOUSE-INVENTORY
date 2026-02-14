package models

import (
	"time"
)

type JualHeader struct {
	ID        int       `json:"id" gorm:"autoIncrement:true"`
	NoFaktur  string    `json:"no_faktur"`
	Customer  string    `json:"customer"`
	Total     float64   `json:"total"`
	UserID    int       `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type JualDetail struct {
	ID           int     `json:"id"`
	JualHeaderID int     `json:"jual_header_id"`
	BarangID     int     `json:"barang_id"`
	Qty          int     `json:"qty"`
	Harga        float64 `json:"harga"`
	Subtotal     float64 `json:"subtotal"`
}

type Penjualan struct {
	ID         int                    `json:"id" gorm:"autoIncrement:true"`
	NoFaktur   string                 `json:"no_faktur"`
	Customer   string                 `json:"customer"`
	Total      float64                `json:"total"`
	UserID     int                    `json:"user_id"`
	Status     string                 `json:"status"`
	CreatedAt  time.Time              `json:"created_at"`
	User       User                   `json:"user" gorm:"foreignKey:UserID"`
	JualDetail []JualDetailWithBarang `json:"jual_detail" gorm:"foreignKey:JualHeaderID"`
}

type JualDetailWithBarang struct {
	ID           int     `json:"id"`
	JualHeaderID int     `json:"jual_header_id"`
	BarangID     int     `json:"barang_id"`
	Qty          int     `json:"qty"`
	Harga        float64 `json:"harga"`
	Subtotal     float64 `json:"subtotal"`
	Barang       Barang  `json:"barang" gorm:"foreignKey:BarangID"`
}

func (JualHeader) TableName() string {
	return "jual_header"
}

func (JualDetail) TableName() string {
	return "jual_detail"
}

func (JualDetailWithBarang) TableName() string {
	return "jual_detail"
}

// -- Table Penjualan Header
// CREATE TABLE jual_header (
//     id SERIAL PRIMARY KEY,
//     no_faktur VARCHAR(100) UNIQUE NOT NULL,
//     customer VARCHAR(200) NOT NULL,
//     total DECIMAL(15,2) DEFAULT 0,
//     user_id INTEGER REFERENCES users(id),
//     status VARCHAR(50) DEFAULT 'selesai',
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );

// -- Table Penjualan Detail
// CREATE TABLE jual_detail (
//     id SERIAL PRIMARY KEY,
//     jual_header_id INTEGER REFERENCES jual_header(id),
//     barang_id INTEGER REFERENCES master_barang(id),
//     qty INTEGER NOT NULL,
//     harga DECIMAL(15,2) NOT NULL,
//     subtotal DECIMAL(15,2) NOT NULL
// );
