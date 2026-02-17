package models

import "time"

type MasterBarang struct {
	ID         int        `json:"id" gorm:"autoIncrement:true"`
	KodeBarang string     `json:"kode_barang"`
	NamaBarang string     `json:"nama_barang"`
	Deskripsi  *string    `json:"deskripsi"`
	Satuan     string     `json:"satuan"`
	HargaBeli  float64    `json:"harga_beli"`
	HargaJual  float64    `json:"harga_jual"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type MasterBarangWithStok struct {
	ID         int       `json:"id" gorm:"autoIncrement:true"`
	KodeBarang string    `json:"kode_barang"`
	NamaBarang string    `json:"nama_barang"`
	Deskripsi  *string   `json:"deskripsi"`
	Satuan     string    `json:"satuan"`
	HargaBeli  float64   `json:"harga_beli"`
	HargaJual  float64   `json:"harga_jual"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Stok       Stok      `json:"stok" gorm:"foreignKey:BarangID;references:ID"`
}

type Barang struct {
	ID         int     `json:"id" gorm:"autoIncrement:true"`
	KodeBarang string  `json:"kode_barang"`
	NamaBarang string  `json:"nama_barang"`
	Satuan     string  `json:"satuan"`
	HargaJual  float64 `json:"harga_jual"`
}

func (MasterBarang) TableName() string {
	return "master_barang"
}

func (MasterBarangWithStok) TableName() string {
	return "master_barang"
}

func (Barang) TableName() string {
	return "master_barang"
}

// -- Table Master Barang
// CREATE TABLE master_barang (
//     id SERIAL PRIMARY KEY,
//     kode_barang VARCHAR(50) UNIQUE NOT NULL,
//     nama_barang VARCHAR(200) NOT NULL,
//     deskripsi TEXT,
//     satuan VARCHAR(50) NOT NULL,
//     harga_beli DECIMAL(15,2) DEFAULT 0,
//     harga_jual DECIMAL(15,2) DEFAULT 0,
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );
