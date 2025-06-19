package models

import "time"

type User struct {
	UserID       uint      `json:"user_id" gorm:"primaryKey"`
	Username     string    `json:"username" gorm:"unique;not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	FullName     string    `json:"full_name" gorm:"not null"`
	PhoneNumber  string    `json:"phone_number"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}