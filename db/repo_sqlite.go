package db

import (
	models "github.com/anovafawzi/socialmedia/models"
	"github.com/jinzhu/gorm"
)

// SQLiteRepository : repo for sqlite db type
type SQLiteRepository struct {
	dbs string
}

// NewSQLiteRepository : create new repository
func NewSQLiteRepository(dbSetting string) *SQLiteRepository {
	return &SQLiteRepository{
		dbs: dbSetting,
	}
}

// InitDb : open gorm DB with sqlite type
func (r *SQLiteRepository) InitDb() *gorm.DB {
	// Opening file
	db, err := gorm.Open("sqlite3", r.dbs)
	//db, err := gorm.Open("sqlite3", "./dbsocmed.db")

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

	return db
}
