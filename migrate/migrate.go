package main

import (
	"log"

	"github.com/ergegananputra/sagara-msib-test/initializers"
	"github.com/ergegananputra/sagara-msib-test/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}
func main() {
	err := initializers.DB.AutoMigrate(&models.Baju{})
	if err != nil {
		log.Fatal("Error migrating Baju model:", err)
	}

	log.Println("Baju model migrated successfully")
}
