package db

import (
	models "github.com/anovafawzi/socialmedia/models"
	"github.com/jinzhu/gorm"
)

// MySQLRepository : repo for sqlite db type
type MySQLRepository struct {
	dbs string
}

// NewMySQLRepository : create new repository
func NewMySQLRepository(dbSetting string) *MySQLRepository {
	return &MySQLRepository{
		dbs: dbSetting,
	}
}

// InitDb : open gorm DB with mysql type
func (r *MySQLRepository) InitDb() *gorm.DB {
	// Opening file
	db, err := gorm.Open("mysql", r.dbs)
	//db, err := gorm.Open("mysql", "root:root@/dbsocmed?charset=utf8&parseTime=True&loc=Local")

	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	db.AutoMigrate(&models.Relations{})

	return db
}
