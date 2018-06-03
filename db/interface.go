package db

import "github.com/jinzhu/gorm"

// Repository interface
type Repository interface {
	InitDb() *gorm.DB
}
