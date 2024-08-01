package models

import "github.com/jinzhu/gorm"

type Product struct {
	ID 			uint `json:"id" gorm:"primary_key"`
	ProductName string `json:"product_name"`
	Price 		float64 `json:"price"`
	Desc    	string `json:"price"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}