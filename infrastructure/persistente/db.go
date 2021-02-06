package persistente

import (
	"flutter-store-api/domain/entity"
	"flutter-store-api/infrastructure/persistente/db"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	resources := db.NewDBResources()
	DBUrl := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", resources.User, resources.Pass, resources.DBName)

	db, err := gorm.Open(mysql.Open(DBUrl), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		panic("Failed to create an user entity")
	}

	return db
}
