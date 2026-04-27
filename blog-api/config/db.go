package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"go-blog/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	var db *gorm.DB
	var err error

	// 🔁 Retry loop
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Successfully connected to the database")
			db.AutoMigrate(&models.User{}, &models.Post{})
			return db
		}
		fmt.Printf("Failed to connect to the database (attempt %d): %v\n", i+1, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatal("❌ Failed to connect to DB:", err)
	return nil

}
