package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int
	FirstName string
	LastName  string
	Latitude  string
	Longitude string
	IsActive  string
}
