package database

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var (
		server   = "dev06srvsqlrad2"
		database = "RadisDev"
	)

	dsn := "server=" + server + ";database=" + database
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect to database!" + err.Error())
	}
	dbConfig, _ := db.DB()
	dbConfig.SetMaxOpenConns(10)
	dbConfig.SetMaxIdleConns(5)
	dbConfig.SetConnMaxLifetime(time.Minute * 30)
	db = db.Debug()
	DB = db

	fmt.Println("Connected to database successfully!")
}
