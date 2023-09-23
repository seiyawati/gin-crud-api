package database

import (
	"gorm.io/gorm"
	"sync"

	"app/config"
	"app/models"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDB() {
	once.Do(func() {
		db = config.ConnectDB()
		migrateDB(db)
	})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	config.DisconnectDB(db)
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Todo{})
}
