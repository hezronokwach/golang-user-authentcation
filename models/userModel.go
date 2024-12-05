package models

import "gorm.io/gorm"

// models/user.go
type User struct {
    gorm.Model
    Email    string `gorm:"uniqueIndex;not null"`
    Password string `gorm:"not null"`
}