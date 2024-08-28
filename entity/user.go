package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
