package models

import "time"

type BeliHeader struct {
	ID        int       `json:"id" gorm:"autoIncrement:true"`
	NoFaktur  string    `json:"no_faktur"`
	Supplier  string    `json:"supplier"`
	Total     float64   `json:"total"`
	UserID    int       `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type BeliDetail struct {
	ID           int     `json:"id"`
	BeliHeaderID int     `json:"jual_header_id"`
	BarangID     int     `json:"barang_id"`
	Qty          int     `json:"qty"`
	Harga        float64 `json:"harga"`
	Subtotal     float64 `json:"subtotal"`
}

type Pembelian struct {
	ID         int                    `json:"id" gorm:"autoIncrement:true"`
	NoFaktur   string                 `json:"no_faktur"`
	Supplier   string                 `json:"supplier"`
	Total      float64                `json:"total"`
	UserID     int                    `json:"user_id"`
	Status     string                 `json:"status"`
	CreatedAt  time.Time              `json:"created_at"`
	User       User                   `json:"user" gorm:"foreignKey:UserID"`
	BeliDetail []BeliDetailWithBarang `json:"beli_detail" gorm:"foreignKey:BeliHeaderID"`
}

type BeliDetailWithBarang struct {
	ID           int     `json:"id"`
	BeliHeaderID int     `json:"beli_header_id"`
	BarangID     int     `json:"barang_id"`
	Qty          int     `json:"qty"`
	Harga        float64 `json:"harga"`
	Subtotal     float64 `json:"subtotal"`
	Barang       Barang  `json:"barang" gorm:"foreignKey:BarangID"`
}

func (BeliHeader) TableName() string {
	return "beli_header"
}

func (BeliDetail) TableName() string {
	return "beli_detail"
}

func (BeliDetailWithBarang) TableName() string {
	return "beli_detail"
}

// type BeliHeader struct {
// 	ID        int       `json:"id" gorm:"autoIncrement:true"`
// 	NoFaktur  string    `json:"no_faktur"`
// 	Supplier  string    `json:"supplier"`
// 	Total     float64   `json:"total"`
// 	UserID    int       `json:"user_id"`
// 	Status    string    `json:"status"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// type BeliDetail struct {
// 	ID           int     `json:"id"`
// 	BeliHeaderID int     `json:"beli_header_id"`
// 	BarangID     int     `json:"barang_id"`
// 	Qty          int     `json:"qty"`
// 	Harga        float64 `json:"harga"`
// 	Subtotal     float64 `json:"subtotal"`
// }

// type Pembelian struct {
// 	ID         int          `json:"id" gorm:"autoIncrement:true"`
// 	NoFaktur   string       `json:"no_faktur"`
// 	Supplier   string       `json:"supplier"`
// 	Total      float64      `json:"total"`
// 	UserID     int          `json:"user_id"`
// 	Status     string       `json:"status"`
// 	CreatedAt  time.Time    `json:"created_at"`
// 	BeliDetail []BeliDetail `json:"beli_detail" gorm:"foreignKey:BeliHeaderID;references:ID"`
// }

// func (BeliHeader) TableName() string {
// 	return "beli_header"
// }

// func (BeliDetail) TableName() string {
// 	return "beli_detail"
// }

// -- Table Pembelian Header
// CREATE TABLE beli_header (
//     id SERIAL PRIMARY KEY,
//     no_faktur VARCHAR(100) UNIQUE NOT NULL,
//     supplier VARCHAR(200) NOT NULL,
//     total DECIMAL(15,2) DEFAULT 0,
//     user_id INTEGER REFERENCES users(id),
//     status VARCHAR(50) DEFAULT 'selesai',
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );

// -- Table Pembelian Detail
// CREATE TABLE beli_detail (
//     id SERIAL PRIMARY KEY,
//     beli_header_id INTEGER REFERENCES beli_header(id),
//     barang_id INTEGER REFERENCES master_barang(id),
//     qty INTEGER NOT NULL,
//     harga DECIMAL(15,2) NOT NULL,
//     subtotal DECIMAL(15,2) NOT NULL
// );
