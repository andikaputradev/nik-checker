package data

import (
	"fmt"
	"os"

	"github.com/buger/jsonparser"
)

var blob []byte

func Load(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("gagal membaca data wilayah: %w", err)
	}
	blob = b
	return nil
}

func LookupProvinsi(kode string) (string, error) {
	return lookup("provinsi", kode)
}

func LookupKabKota(kode string) (string, error) {
	return lookup("kabkot", kode)
}

func LookupKecamatan(kode string) (string, error) {
	return lookup("kecamatan", kode)
}

func lookup(kategori, kode string) (string, error) {
	if blob == nil {
		return "", fmt.Errorf("data wilayah belum dimuat")
	}
	val, err := jsonparser.GetString(blob, kategori, kode)
	if err != nil {
		return "", fmt.Errorf("tidak ditemukan %s untuk kode %s", kategori, kode)
	}
	return val, nil
}