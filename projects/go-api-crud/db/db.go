package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
   "github.com/joho/godotenv"
)

var db *gorm.DB
var err error

type Product struct {
	ID 			string `json:"id" gorm:"primary_key"`
	ProductName string `json:"product_name"`
	Price 		float64 `json:"price"`
	Desc    	string `json:"desc"`
}

func InitPostgredDB() {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	
	dbUrl := os.Getenv("DATABASE_URL")

	var err error
	db, err = gorm.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Failed to connect db:", err)
	}

	if err := db.DB().Ping(); err != nil {
		log.Fatal("Cannot ping db:", err)
	}

	db.AutoMigrate(&Product{})

	fmt.Println("Database connection established successfully!")
}
