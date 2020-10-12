package models

import "gorm.io/gorm"

// User Model
type User struct {
	gorm.Model
	Name string
	Age  uint8
}
