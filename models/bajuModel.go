package models

import (
	"github.com/ergegananputra/sagara-msib-test/helpers"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Baju struct {
	gorm.Model
	Name string
	Warna string
	Ukuran string
	Harga decimal.Decimal `gorm:"type:numeric"`
	Stok int
}

// Pembuatan model Baju wajib menggunakan fungsi CreateBajuModel sebelum menyimpan ke database.
// Fungsi ini akan menyeragamkan pattern data yang akan disimpan ke database.
func CreateBajuModel(
	name string, warna string, ukuran string, harga string, stok int, 
	interceptor helpers.StringToDecimalInterceptor,
) Baju {
	isEmpty, hargaDecimal, err := helpers.StringToDecimal(harga)
	
	if interceptor != nil {
		interceptor(isEmpty, &hargaDecimal, err)
	} else if err != nil {
		panic(err)
	} else {
		hargaDecimal = decimal.NewFromInt(0)
	}

	return Baju{
		Name: helpers.StringTrimAndLower(name),
		Warna: helpers.StringTrimAndLower(warna),
		Ukuran: helpers.StringTrimAndLower(ukuran),
		Harga: hargaDecimal,
		Stok: stok,
	}
}