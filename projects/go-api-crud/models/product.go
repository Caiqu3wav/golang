package models


type Product struct {
	ID 			uint `json:"id" gorm:"primary_key"`
	ProductName string `json:"product_name"`
	Price 		float64 `json:"price"`
	Desc    	string `json:"desc"`
}