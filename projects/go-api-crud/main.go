package main

import (
	"github.com/gin-gonic/gin"
    "github.com/steebchen/prisma-client-go"
	"log"
)

func main() {
	client := prisma.NewClient()
	err := client.Prisma.Connect()
	if err != nil {
		log.Fatal("Erro ao conectar db", err)
	}
	defer client.Prisma.Disconnect()

	server := gin.Default()

	SetupProductRoutes(r, client)

	server.Run(":8000")
}