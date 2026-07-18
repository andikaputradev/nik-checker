package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"nik-checker/internal/data"
	"nik-checker/internal/nik"
	"nik-checker/internal/output"
)

func main() {
	var dataPath string
	var inputFile string
	var outputFormat string
	flag.StringVar(&dataPath, "data", "data.json", "path ke data.json")
	flag.StringVar(&inputFile, "input", "", "file batch berisi daftar NIK, satu per baris")
	flag.StringVar(&outputFormat, "format", "table", "format output: table atau json")
	flag.Parse()

	if err := data.Load(dataPath); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	if inputFile != "" {
		processBatch(inputFile, outputFormat)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input NIK: ")
	nikStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Gagal membaca input: %v", err)
	}
	nikStr = strings.TrimSpace(nikStr)
	processSingle(nikStr, outputFormat)
}

func processSingle(nikStr string, format string) {
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

	switch format {
	case "json":
		output.FormatJSON(nikStr, prov, kab, kec, kodePos, parsed.TanggalLahirStr, parsed.BulanLahirStr, parsed.TahunLahirStr, parsed.JenisKelamin, parsed.UniqCode)
	default:
		output.FormatTable(nikStr, prov, kab, kec, kodePos, parsed.TanggalLahirStr, parsed.BulanLahirStr, parsed.TahunLahirStr, parsed.JenisKelamin, parsed.UniqCode)
	}
}

func parseKecamatan(raw string) (string, string) {
	parts := strings.SplitN(raw, "--", 2)
	if len(parts) == 2 {
		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	}
	return strings.TrimSpace(parts[0]), "TIDAK DIKETAHUI"
}

func processBatch(filePath string, format string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Gagal membuka file batch: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		processSingle(line, format)
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error membaca file batch: %v", err)
	}
}