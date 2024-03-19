package database

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormInit(sqlDB *sql.DB) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error when connect to database: %v", err)
	}

	return db
}
