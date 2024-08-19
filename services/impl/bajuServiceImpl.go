package services

import (
	"strings"

	"github.com/ergegananputra/sagara-msib-test/configs"
	"github.com/ergegananputra/sagara-msib-test/configs/validations"
	"github.com/ergegananputra/sagara-msib-test/helpers"
	"github.com/ergegananputra/sagara-msib-test/models"
	"github.com/shopspring/decimal"
)

type BajuServiceImpl struct{}

func (e *BajuServiceImpl) CreateBaju(model *models.BajuBasicRequest, interceptor helpers.StringToDecimalInterceptor) (models.Baju, []string, error) {
	
	baju := createBajuModel(model.Name, model.Warna, model.Ukuran, model.Harga, model.Stok, interceptor)

	errs := validations.ValidateBaju(&baju)

	if len(errs) > 0 {
		return baju, errs, nil
	}

	result := configs.DB.Create(&baju)

	err := result.Error

	if err != nil {
		return baju, nil, err
	}

	return baju, nil, nil
}

func (e *BajuServiceImpl) GetBajus() ([]models.Baju, error) {
	var bajus []models.Baju
	result := configs.DB.Find(&bajus)

	err := result.Error

	if err != nil {
		return nil, err
	}

	return bajus, nil
}

func (e *BajuServiceImpl) GetBaju(id *string) (models.Baju, error) {
	var baju models.Baju
	result := configs.DB.First(&baju, id)

	if result.Error != nil {
		conditions := result.Error.Error()
		if strings.Contains(conditions, "record not found") {
			return baju, nil
		} else {
			return baju, result.Error
		}
	}

	return baju, nil
}

func (e *BajuServiceImpl) UpdateBaju(id *string, model models.BajuBasicRequest) (models.Baju, error) {
	var baju models.Baju
	result := configs.DB.First(&baju, id)

	if result.Error != nil {
		return baju, result.Error
	}

	var thisError error
	newBaju := createBajuModel(model.Name, model.Warna, model.Ukuran, model.Harga, model.Stok, 
		func(isEmpty bool, value *decimal.Decimal, err error) {
			if err != nil {
				thisError = err
			}
			if isEmpty {
				*value = baju.Harga
			}
		},
	)

	if thisError != nil {
		return models.Baju{
			Name:  "%CODE%_ERROR",
		}, thisError
	}

	result = configs.DB.Model(&baju).Updates(models.Baju{
		Name:  newBaju.Name,
		Warna: newBaju.Warna,
		Ukuran: newBaju.Ukuran,
		Harga: newBaju.Harga,
		Stok:  newBaju.Stok,
	})

	err := result.Error

	if err != nil {
		return baju, err
	}

	return baju, nil
}
func (e *BajuServiceImpl) DeleteBaju(id *string)( models.Baju, error) {
	var baju models.Baju
	result := configs.DB.Delete(&baju, id)

	if result.Error != nil {
		return baju, result.Error
	}

	return baju, nil
}
func (e *BajuServiceImpl) SearchByWarnaAndUkuranBaju(warna string, ukuran string) ([]models.Baju, error) {
	db := configs.DB
    if warna != "" {
        db = db.Where("LOWER(Warna) = ?", &warna)
    }
    if ukuran != "" {
        db = db.Where("LOWER(Ukuran) = ?", &ukuran)
    }

	var bajus []models.Baju
    result := db.Find(&bajus)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return bajus, nil

}
func (e *BajuServiceImpl) AddStokBaju(id *string, stok int) (models.Baju, error) {
	var baju models.Baju
	result := configs.DB.First(&baju, id)

	if result.Error != nil {
		return baju, result.Error
	}

	newStock := baju.Stok + stok

	result = configs.DB.Model(&baju).Update("Stok", newStock)

	err := result.Error

	if err != nil {
		conditions := result.Error.Error()
		if strings.Contains(conditions, "record not found") {
			return models.Baju{}, nil
		} else {
			return models.Baju{}, result.Error
		}
	}

	return baju, nil
}
func (e *BajuServiceImpl) ReduceStokBaju(id *string, stok int)(models.Baju, error) {
	var baju models.Baju
	result := configs.DB.First(&baju, id)

	if result.Error != nil {
		return baju, result.Error
	}

	newStock := baju.Stok - stok
	if newStock < 0 {
		newStock = 0
	}

	result = configs.DB.Model(&baju).Update("Stok", newStock)

	err := result.Error

	if err != nil {
		conditions := result.Error.Error()
		if strings.Contains(conditions, "record not found") {
			return models.Baju{}, nil
		} else {
			return models.Baju{}, result.Error
		}
	}

	return baju, nil
}
func (e *BajuServiceImpl) StokEmptyBaju() []models.Baju {
	var bajus []models.Baju
	result := configs.DB.Where("Stok = 0").Find(&bajus)

	if result.Error != nil {
		return nil
	}

	return bajus
}
func (e *BajuServiceImpl) StockAlertBaju(limit *int) []models.Baju {
	var bajus []models.Baju
	result := configs.DB.Where("Stok < ?", &limit).Find(&bajus)

	if result.Error != nil {
		return nil
	}

	return bajus
}


func createBajuModel(
	name string, warna string, ukuran string, harga string, stok int, 
	interceptor helpers.StringToDecimalInterceptor,
) models.Baju {
	isEmpty, hargaDecimal, err := helpers.StringToDecimal(harga)
	
	if interceptor != nil {
		interceptor(isEmpty, &hargaDecimal, err)
	} else if err != nil {
		panic(err)
	} else {
		hargaDecimal = decimal.NewFromInt(0)
	}

	return models.Baju{
		Name: helpers.StringTrimAndLower(name),
		Warna: helpers.StringTrimAndLower(warna),
		Ukuran: helpers.StringTrimAndLower(ukuran),
		Harga: hargaDecimal,
		Stok: stok,
	}
}