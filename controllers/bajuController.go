package controllers

import (
	"fmt"
	"strings"

	"github.com/ergegananputra/sagara-msib-test/models"
	servicesI "github.com/ergegananputra/sagara-msib-test/services"
	services "github.com/ergegananputra/sagara-msib-test/services/impl"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

/**
 * Utilitas Baju Controller
 */

func bindRequest(c *gin.Context) (models.BajuBasicRequest, error) {
	var body models.BajuBasicRequest
	err := c.Bind(&body)
	if err != nil {
		return body, err
	}
	return body, nil
}

/**
 * Controller Baju
 */

var bajuService servicesI.BajuServiceInterface = &services.BajuServiceImpl{}

func CreateBaju(c *gin.Context) {
	body, err := bindRequest(c)
	if err != nil {
		c.JSON(400, models.Response{
			Message: "Permintaan tidak valid",
		})
		return
	}

	baju, validationErrors, err := bajuService.CreateBaju(
		&body,
		func(isEmpty bool, value *decimal.Decimal, err error) {
			if isEmpty {
				c.JSON(400, models.Response{
					Message: "Harga tidak boleh kosong",
				})
				return
			}
			if err != nil {
				c.JSON(400, models.Response{
					Message: "Format harga tidak valid",
				})
				return
			}
		},
	)

	if len(validationErrors) > 0 {
		c.JSON(400, models.Response{
			Message: "Validation errors",
			Data:    validationErrors,
		})
		return
	}

	if err != nil {
		c.JSON(500, models.Response{
			Message: "Gagal membuat baju",
		})
		return
	}

	c.JSON(200, models.Response{
		Message: "Berhasil membuat baju",
		Data:    baju,
	})
}

func GetBajus(c *gin.Context) {
	bajus, err := bajuService.GetBajus()

	if err != nil {
		c.JSON(500, models.Response{
			Message: "Gagal menemukan baju",
		})
		return
	}

	c.JSON(200, models.Response{
		Message: "Berhasil menemukan baju",
		Data:    bajus,
	})
}

func GetBaju(c *gin.Context) {
	id := c.Param("id")

	baju, err := bajuService.GetBaju(&id)

	if err != nil {
		c.JSON(500, models.Response{
			Message: "Gagal menemukan baju dengan id " + id,
			Data:    nil,
		})
		return
	}

	if baju.ID == 0 {
		c.JSON(404, models.Response{
			Message: "Baju dengan id " + id + " tidak ditemukan",
			Data:    nil,
		})
		return
	}

	c.JSON(200, models.Response{
		Message: "Berhasil menemukan baju",
		Data:    baju,
	})
}

func UpdateBaju(c *gin.Context) {
	id := c.Param("id")

	body, err := bindRequest(c)

	if err != nil {
		c.JSON(400, models.Response{
			Message: "Permintaan tidak valid",
		})
		return
	}

	baju, err := bajuService.UpdateBaju(&id, body)

	if baju.ID != 0 && err != nil {
		c.JSON(400, models.Response{
			Message: "Permintaan tidak valid",
		})

		return
	}

	if baju.ID == 0 && err != nil {
		c.JSON(400, models.Response{
			Message: "Format harga tidak valid",
		})
		return
	}

	c.JSON(200, models.Response{
		Message: "Baju dengan id " + id + " berhasil diupdate",
		Data:    baju,
	})
}

func DeleteBaju(c *gin.Context) {
	id := c.Param("id")

	baju, err := bajuService.DeleteBaju(&id)

	if err != nil {
		c.JSON(500, models.Response{
			Message: "Gagal menghapus baju dengan id " + id + " " + err.Error(),
		})
		return
	}

	c.JSON(200, models.Response{
		Message: "Berhasil menghapus baju dengan id " + id,
		Data:    baju,
	})
}

func SearchByWarnaAndUkuranBaju(c *gin.Context) {
	warna := strings.ToLower(c.Query("warna"))
	ukuran := strings.ToLower(c.Query("ukuran"))

	bajus, err := bajuService.SearchByWarnaAndUkuranBaju(warna, ukuran)

	if err != nil {
		c.JSON(500, models.Response{
			Message: "Gagal menemukan baju",
		})
		return
	}

	if bajus == nil {
		c.JSON(404, models.Response{
			Message: "Tidak ada baju yang cocok",
		})
		return
	}

	c.JSON(200, models.Response{
		Message: "Berhasil menemukan baju",
		Data:    bajus,
	})
}

func AddStokBaju(c *gin.Context) {
	id := c.Param("id")

	var requestBody struct {
		Stok int `json:"stok"`
	}

	err := c.Bind(&requestBody)

	if err != nil {
		c.JSON(400, models.Response{
			Message: "Stok tidak boleh kosong",
		})
		return
	}

	baju, err := bajuService.AddStokBaju(&id, requestBody.Stok)
	if err == nil && baju.ID == 0 {
		c.JSON(404, models.Response{
			Message: "Baju dengan id " + id + " tidak ditemukan",
			Data:    nil,
		})
		return
	}
	if err == nil && baju.ID != 0 {
		c.JSON(500, models.Response{
			Message: "Gagal menemukan baju dengan id " + id,
			Data:    nil,
		})
		return
	}

	c.JSON(200, models.Response{
		Message: "Stok baju dengan id " + id + " berhasil diupdate",
		Data:    baju,
	})
}

func ReduceStokBaju(c *gin.Context) {
	id := c.Param("id")

	var requestBody struct {
		Stok int `json:"stok"`
	}

	err := c.Bind(&requestBody)

	if err != nil {
		c.JSON(400, models.Response{
			Message: "Stok tidak boleh kosong",
		})
		return
	}

	baju, err := bajuService.ReduceStokBaju(&id, requestBody.Stok)
	if err == nil && baju.ID == 0 {
		c.JSON(404, models.Response{
			Message: "Baju dengan id " + id + " tidak ditemukan",
			Data:    nil,
		})
		return
	}
	if err == nil && baju.ID == 0 {
		c.JSON(500, models.Response{
			Message: "Gagal menemukan baju dengan id " + id,
			Data:    nil,
		})
		return
	}

	c.JSON(200, models.Response{
		Message: "Stok baju dengan id " + id + " berhasil diupdate",
		Data:    baju,
	})
}

func StokEmptyBaju(c *gin.Context) {
	bajus := bajuService.StokEmptyBaju()

	if bajus == nil {
		c.JSON(404, models.Response{
			Message: "Tidak ada baju yang stoknya kosong",
		})
		return
	}

	c.JSON(200, models.Response{
		Message: "Berhasil menemukan baju yang stoknya kosong",
		Data:    bajus,
	})
}

func StockAlertBaju(c *gin.Context) {
	limit := 5
	bajus := bajuService.StockAlertBaju(&limit)

	if bajus == nil {
		c.JSON(404, models.Response{
			Message: fmt.Sprintf("Tidak ada baju yang stoknya kurang dari %d", limit),
		})
		return
	}

	c.JSON(200, models.Response{
		Message: fmt.Sprintf("Berhasil menemukan baju yang stoknya kurang dari %d", limit),
		Data:    bajus,
	})
}
