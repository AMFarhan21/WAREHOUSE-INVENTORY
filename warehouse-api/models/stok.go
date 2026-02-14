package models

import "time"

type Mstok struct {
	ID        int       `json:"id" gorm:"autoIncrement:true"`
	BarangID  int       `json:"barang_id"`
	StokAkhir int       `json:"stok_akhir"`
	UpdatedAt time.Time `json:"updated_at"`
	Barang    Barang    `json:"barang" gorm:"foreignKey:BarangID"`
}

type Stok struct {
	ID        int       `json:"id" gorm:"autoIncrement:true"`
	BarangID  int       `json:"barang_id"`
	StokAkhir int       `json:"stok_akhir"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Mstok) TableName() string {
	return "mstok"
}

func (Stok) TableName() string {
	return "mstok"
}

// -- Table Stok
// CREATE TABLE mstok (
//     id SERIAL PRIMARY KEY,
//     barang_id INTEGER REFERENCES master_barang(id),
//     stok_akhir INTEGER DEFAULT 0,
//     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );
