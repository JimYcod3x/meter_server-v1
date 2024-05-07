package database

import (
	"fmt"

	"github.com/JimYcod3x/meter_server/config"
	"github.com/JimYcod3x/meter_server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	if err := db.AutoMigrate(&models.Meter{}); err != nil {
		panic("failed to auto migrate tables: " + err.Error())
	}

	fmt.Println("Connected Successful to the dabase(mysql)")

	
	return db
}