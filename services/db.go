package services

import (
	"LBD/database"

	"gorm.io/gorm"
)

var connectDB = database.ConnectDB

func setConnectDB(fn func() (*gorm.DB, error)) {
	connectDB = fn
}

func resetConnectDB() {
	connectDB = database.ConnectDB
}
