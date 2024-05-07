package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(localhost:3306)/testgorm?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Creat
}