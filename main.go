package main

import (
	"github.com/ergegananputra/sagara-msib-test/configs"
	"github.com/ergegananputra/sagara-msib-test/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectToDatabase()
}

/** 
 * Sebelum menjalankan aplikasi, buatlah file .env contoh seperti .env.example. 
 * Kemudian jalankan Script Migration terlebih dahulu
 * go run migrate/migrate.go
 */
func main() {
	r := gin.Default()
	routers.ApiRoutes(r)
	r.Run()
}