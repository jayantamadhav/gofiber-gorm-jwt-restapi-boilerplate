package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	Title     string         `json:"title"`
	SubTitle  string         `json:"sub_title"`
	Text      string         `json:"text"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
