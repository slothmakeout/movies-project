package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID     uint      `json:"userId"`
	MovieID    uint      `json:"movieId"`
	Rating     float32   `json:"rating"`
	ReviewText string    `json:"reviewText"`
	Timestamp  time.Time `json:"timestamp"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) error {
	r.UserID = 1
	r.Timestamp = time.Now()
	return nil
}