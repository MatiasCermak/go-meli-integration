package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Initialize() *gorm.DB {
	dsn := "matrox32:Matrox23@tcp(www.db4free.net:3306)/gomeli?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database ")
	}
	return db
}
