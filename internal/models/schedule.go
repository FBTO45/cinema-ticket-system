package models

import (
	"time"
)

type Schedule struct {
	ScheduleID uint      `json:"schedule_id" gorm:"primaryKey"`
	MovieID    uint      `json:"movie_id" gorm:"not null"`
	StudioID   uint      `json:"studio_id" gorm:"not null"`
	StartTime  time.Time `json:"start_time" gorm:"not null"`
	EndTime    time.Time `json:"end_time" gorm:"not null"`
	Price      float64   `json:"price" gorm:"not null;type:decimal(10,2)"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}