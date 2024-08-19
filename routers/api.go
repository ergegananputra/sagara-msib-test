package routers

import (
	"github.com/ergegananputra/sagara-msib-test/controllers"
	"github.com/gin-gonic/gin"
)


func ApiRoutes(r *gin.Engine) {
	r.POST("/baju/create", controllers.CreateBaju)
	r.GET("/baju", controllers.GetBajus)
	r.GET("/baju/:id", controllers.GetBaju)
	r.PUT("/baju/:id", controllers.UpdateBaju)
	r.DELETE("/baju/:id", controllers.DeleteBaju)

	r.GET("/baju/search", controllers.SearchByWarnaAndUkuranBaju)
	r.PATCH("/baju/:id/add-stok", controllers.AddStokBaju)
	r.PATCH("/baju/:id/reduce-stok", controllers.ReduceStokBaju)

	r.GET("/baju/empty", controllers.StokEmptyBaju)
	r.GET("/baju/alert", controllers.StockAlertBaju)
}