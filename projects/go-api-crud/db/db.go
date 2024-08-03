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

func CreateProduct(product *Product) (*Product, error) {
	product.ID = uuid.New().String()
	res := db.Create(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func GetProduct(id string) (*Product, error) {
	var product Product
	res := db.First(&product, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("Movie of id %s not found", id))
	}
	return &product, nil
}

func GetProducts() ([]*Product, error) {
	var products []*Product
	res := db.Find(&products)
	if res.Error != nil {
		return nil, errors.New("No products found")
	}
	return products, nil
}

func UpdateProduct(product *Product) (*Product, error) {
	var product2Update Product
	result := db.Model(&product2Update).Where("id = ?", product.ID).Updates(product)
	if result.RowsAffected == 0 {
		return &product2Update, errors.New("movie not updated")
	}
	return product, nil
}

func DeleteProduct(id string) error {
	var deletedProduct Product
	result := db.Where("id = ?", id).Delete(&deletedProduct)
	if result.RowsAffected == 0 {
		return errors.New("movie not deleted")
	}
	return nil
}