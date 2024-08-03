package main

import (
	"go-api-crud/db"
	"go-api-crud/router"
)

func main() {
	db.InitPostgredDB()
	router.InitRouter().Run()
}