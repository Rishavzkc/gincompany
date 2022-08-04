package database

import (
	"final/config"
	"final/models"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDatabase() {
	DbHost := config.GetConfig().DatabaseHost
	DbPort := config.GetConfig().DatabasePort
	DbUser := config.GetConfig().DatabaseUser
	DbName := config.GetConfig().DatabaseName
	DbPass := config.GetConfig().DatabasePass

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", DbUser, DbPass, DbHost, DbPort, DbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the sql Database")
		log.Fatal("Error: ", err)
	}

	db = database

	config, _ := database.DB()

	config.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&models.Company{})

}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
