package models

import (
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
