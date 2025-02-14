package models

import "gorm.io/gorm"

// users model
type User struct {
    gorm.Model
    Username string `json:"username"`
    Password string `json:"password"`
}