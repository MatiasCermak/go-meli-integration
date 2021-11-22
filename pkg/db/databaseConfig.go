package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Initialize() *gorm.DB {
	dsn := "sql10450724:Kf6AUi3vJU@tcp(sql10.freemysqlhosting.net:3306)/sql10450724?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database ")
	}
	return db
}
