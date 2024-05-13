package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/JimYcod3x/meter_server/config"
)

func ConnectionDB(config *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBName)
	fmt.Println(dsn)
	db, err := sql.Open("mysql",dsn)
	if err != nil {
		fmt.Println("can not connected to db", err)
		return nil
	}


	fmt.Println("Connected Successful to the dabase(mysql)")

	
	return db
}
