package main

import (
	"log"

	"github.com/ergegananputra/sagara-msib-test/configs"
	"github.com/ergegananputra/sagara-msib-test/models"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectToDatabase()
}

func main() {
	err := configs.DB.AutoMigrate(&models.Baju{})
	if err != nil {
		log.Fatal("Error migrating Baju model:", err)
	}

	log.Println("Baju model migrated successfully")
}
