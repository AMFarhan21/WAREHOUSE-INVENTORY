package models

import "time"

type Users struct {
	ID        int       `json:"id" gorm:"autoIncrement:true"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID       int    `json:"id" gorm:"autoIncrement:true"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

// CREATE TABLE users (
//     id SERIAL PRIMARY KEY,
//     username VARCHAR(100) UNIQUE NOT NULL,
//     password VARCHAR(255) NOT NULL,
//     email VARCHAR(150) UNIQUE NOT NULL,
//     full_name VARCHAR(200) NOT NULL,
//     role VARCHAR(50) DEFAULT 'staff',
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );
