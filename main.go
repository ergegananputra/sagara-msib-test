package main

import (
	"github.com/ergegananputra/sagara-msib-test/controllers"
	"github.com/ergegananputra/sagara-msib-test/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

/** 
 * Sebelum menjalankan aplikasi, jalankan Script Migration terlebih dahulu
 * go run migrate/migrate.go
 */
func main() {
	r := gin.Default()
	r.POST("/baju/create", controllers.CreateBaju)
	r.GET("/baju", controllers.GetBajus)
	r.GET("/baju/:id", controllers.GetBaju)
	r.PUT("/baju/:id", controllers.UpdateBaju)
	r.DELETE("/baju/:id", controllers.DeleteBaju)
	r.Run()
}