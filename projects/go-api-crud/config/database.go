package config

import (
	"golang-prisma/prisma/db"
	"log"
)

func ConnectDB()(*db.PrismaClient, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	log.Info().Msg("Connect to db!")
	return client, nil
}