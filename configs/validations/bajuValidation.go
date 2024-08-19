package validations

import "github.com/ergegananputra/sagara-msib-test/models"

func ValidateBaju(baju *models.Baju) []string {
	var validationErrors []string

	if baju.Name == "" {
		validationErrors = append(validationErrors, "Nama tidak boleh kosong")
	}
	if baju.Warna == "" {
		validationErrors = append(validationErrors, "Warna tidak boleh kosong")
	}
	if baju.Ukuran == "" {
		validationErrors = append(validationErrors, "Ukuran tidak boleh kosong")
	}
	if baju.Harga.IsZero() {
		validationErrors = append(validationErrors, "Harga tidak boleh kosong")
	}
	if baju.Stok == 0 {
		validationErrors = append(validationErrors, "Stok tidak boleh kosong")
	}

	return validationErrors
}