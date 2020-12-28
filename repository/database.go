package repository

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"moviesearch/config"
	"moviesearch/models"
)

func InitDatabase() {

	db := Connect()
	defer Disconnect(db)

	if err := db.AutoMigrate(&models.Movie{}); err != nil {
		panic("Could not automigrate models into db")
	}
	fmt.Println("Database migrated")

}

func Connect() *gorm.DB {
	dbPath := config.Config("DB-PATH")
	db, err := gorm.Open(sqlite.Open(dbPath+"/cache.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Db connection opened successfully")
	return db
}

func Disconnect(db *gorm.DB) {
	dbC, _ := db.DB()
	if err := dbC.Close(); err != nil {
		panic("Cannot close db connection")
	}
	fmt.Println("DB connection closed successfully")
}
