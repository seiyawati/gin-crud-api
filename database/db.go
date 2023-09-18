package database

import (
	"sync"
	"gorm.io/gorm"

	"app/config"
)

var (
	db  *gorm.DB
	err error
	once sync.Once
)

func InitDB() {
	once.Do(func() {
		db = config.ConnectDB()
	})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	config.DisconnectDB(db)
}
