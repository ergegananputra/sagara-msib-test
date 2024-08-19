package controllers

import (
	"strings"

	"github.com/ergegananputra/sagara-msib-test/initializers"
	"github.com/ergegananputra/sagara-msib-test/models"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func CreateBaju(c *gin.Context) {
	var body struct {
		Name string
		Warna string
		Ukuran string
		Harga string
		Stok int
	}

	c.Bind(&body)

	if !strings.Contains(body.Harga, ".") {
        body.Harga = body.Harga + ".00"
    }


	harga, err := decimal.NewFromString(body.Harga)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid harga format"})
        return
    }

	baju := models.Baju{
		Name:  body.Name,
		Warna: body.Warna,
		Ukuran: body.Ukuran,
		Harga: harga,
		Stok:  body.Stok,
	}

	result := initializers.DB.Create(&baju)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create Baju",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Create Baju success",
		"data":    baju,
	})
}

func GetBajus(c *gin.Context) {
	var bajus []models.Baju
	initializers.DB.Find(&bajus)

	c.JSON(200, gin.H{
		"data": bajus,
	})
}

func GetBaju(c *gin.Context) {
	id := c.Param("id")

	var baju models.Baju
	initializers.DB.First(&baju, id)

	c.JSON(200, gin.H{
		"data": baju,
	})
}

func UpdateBaju(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name string
		Warna string
		Ukuran string
		Harga string
		Stok int
	}

	c.Bind(&body)

	if !strings.Contains(body.Harga, ".") {
        body.Harga = body.Harga + ".00"
    }

	var baju models.Baju
	initializers.DB.First(&baju, id)

	if body.Harga == ".00" {
		body.Harga = baju.Harga.String()
	}

	initializers.DB.Model(&baju).Updates(models.Baju{
		Name:  body.Name,
		Warna: body.Warna,
		Ukuran: body.Ukuran,
		Harga: decimal.RequireFromString(body.Harga),
		Stok:  body.Stok,
	})

	c.JSON(200, gin.H{
		"data": baju,
	})
}

func DeleteBaju(c *gin.Context) {
	id := c.Param("id")

	var baju models.Baju
	initializers.DB.Delete(&baju, id)

	c.JSON(200, gin.H{
		"message": "Delete Baju success",
		"data": baju,
	})
}