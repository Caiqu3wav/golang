package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/google/uuid"
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
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect db:", err)
	}

	sqlDB, err := db.DB() // Aqui pegamos a conexão *sql.DB
	if err != nil {
		log.Fatal("Failed to get DB connection:", err)
	}

	if err := sqlDB.Ping(); err != nil { // Agora usamos o método Ping corretamente
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
		return nil, errors.New(fmt.Sprintf("Product of id %s not found", id))
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
		return &product2Update, errors.New("product not updated")
	}
	return product, nil
}

func DeleteProduct(id string) error {
	var deletedProduct Product
	result := db.Where("id = ?", id).Delete(&deletedProduct)
	if result.RowsAffected == 0 {
		return errors.New("product not deleted")
	}
	return nil
}