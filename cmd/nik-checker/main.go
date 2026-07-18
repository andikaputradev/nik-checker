package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"nik-checker/internal/data"
	"nik-checker/internal/nik"
	"nik-checker/internal/output"
)

func main() {
	if err := data.Load("data.json"); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input NIK: ")
	nikStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Gagal membaca input: %v", err)
	}
	nikStr = strings.TrimSpace(nikStr)

	parsed, err := nik.Parse(nikStr)
	if err != nil {
		log.Fatalf("NIK tidak valid: %v", err)
	}

	prov, err := data.LookupProvinsi(parsed.KodeProvinsi)
	if err != nil {
		prov = "TIDAK DIKETAHUI"
	}
	kab, err := data.LookupKabKota(parsed.KodeKabKota)
	if err != nil {
		kab = "TIDAK DIKETAHUI"
	}
	kecRaw, err := data.LookupKecamatan(parsed.KodeKecamatan)
	if err != nil {
		kecRaw = "TIDAK DIKETAHUI--"
	}
	kec, kodePos := parseKecamatan(kecRaw)

	output.FormatTable(nikStr, prov, kab, kec, kodePos,
		parsed.TanggalLahirStr, parsed.BulanLahirStr, parsed.TahunLahirStr,
		parsed.JenisKelamin, parsed.UniqCode)
}

func parseKecamatan(raw string) (string, string) {
	parts := strings.SplitN(raw, "--", 2)
	if len(parts) == 2 {
		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	}
	return strings.TrimSpace(parts[0]), "TIDAK DIKETAHUI"
}