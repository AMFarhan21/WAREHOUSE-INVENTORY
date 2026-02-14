package models

import "time"

type HistoryStok struct {
	ID             int       `json:"id" gorm:"autoIncrement:true"`
	BarangID       int       `json:"barang_id"`
	UserID         int       `json:"user_id"`
	JenisTransaksi string    `json:"jenis_transaksi"`
	Jumlah         int       `json:"jumlah"`
	StokSebelum    int       `json:"stok_sebelum"`
	StokSesudah    int       `json:"stok_sesudah"`
	Keterangan     string    `json:"keterangan"`
	CreatedAt      time.Time `json:"created_at"`
	Barang         Barang    `json:"barang" gorm:"foreignKey:BarangID"`
	User           User      `json:"user" gorm:"foreignKey:UserID"`
}

func (HistoryStok) TableName() string {
	return "history_stok"
}

// -- Table History Stok
// CREATE TABLE history_stok (
//     id SERIAL PRIMARY KEY,
//     barang_id INTEGER REFERENCES master_barang(id),
//     user_id INTEGER REFERENCES users(id),
//     jenis_transaksi VARCHAR(50) NOT NULL, -- 'masuk', 'keluar', 'adjustment'
//     jumlah INTEGER NOT NULL,
//     stok_sebelum INTEGER NOT NULL,
//     stok_sesudah INTEGER NOT NULL,
//     keterangan TEXT,
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );
