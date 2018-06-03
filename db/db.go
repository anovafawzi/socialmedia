package db

import (
	models "github.com/anovafawzi/socialmedia/models"
	"github.com/jinzhu/gorm"
)

// InitDb : open gorm DB with sqlite type
func InitDb() *gorm.DB {
	// Opening file
	db, err := gorm.Open("sqlite3", "./dbsocmed.db")
	//db, err := gorm.Open("mysql", "root:root@/dbsocmed?charset=utf8&parseTime=True&loc=Local")

	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&models.Relations{}) {
		db.CreateTable(&models.Relations{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Relations{})
	}
	//db.AutoMigrate(&Relations{})

	return db
}
