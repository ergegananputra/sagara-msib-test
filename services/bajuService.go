package services

import (
	"github.com/ergegananputra/sagara-msib-test/helpers"
	"github.com/ergegananputra/sagara-msib-test/models"
)

type BajuServiceInterface interface {
	CreateBaju(model *models.BajuBasicRequest, interceptor helpers.StringToDecimalInterceptor) (models.Baju, []string, error)
	GetBajus() ([]models.Baju, error)
	GetBaju(id *string) (models.Baju, error)
	UpdateBaju(id *string, model models.BajuBasicRequest) (models.Baju, error)
	DeleteBaju(id *string) ( models.Baju, error)
	SearchByWarnaAndUkuranBaju(warna string, ukuran string) ([]models.Baju, error)
	AddStokBaju(id *string, stok int) (models.Baju, error)
	ReduceStokBaju(id *string, stok int) (models.Baju, error)
	StokEmptyBaju() []models.Baju
	StockAlertBaju(limit *int) []models.Baju
}