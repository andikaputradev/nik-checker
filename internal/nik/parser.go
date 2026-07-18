package nik

import (
	"errors"
	"strconv"
)

type Data struct {
	KodeProvinsi    string
	KodeKabKota     string
	KodeKecamatan   string
	TanggalLahirStr string
	BulanLahirStr   string
	TahunLahirStr   string
	UniqCode        string
	JenisKelamin    string
}

func Parse(nik string) (Data, error) {
	if len(nik) != 16 {
		return Data{}, errors.New("panjang NIK harus tepat 16 digit")
	}
	for i := 0; i < len(nik); i++ {
		if nik[i] < '0' || nik[i] > '9' {
			return Data{}, errors.New("NIK hanya boleh berisi digit 0-9")
		}
	}
	tglStr := nik[6:8]
	tglInt, err := strconv.Atoi(tglStr)
	if err != nil {
		return Data{}, errors.New("gagal membaca kode tanggal")
	}
	blnStr := nik[8:10]
	blnInt, err := strconv.Atoi(blnStr)
	if err != nil || blnInt < 1 || blnInt > 12 {
		return Data{}, errors.New("kode bulan tidak valid (harus 01-12)")
	}
	thnStr := nik[10:12]
	uniq := nik[12:16]
	prov := nik[0:2]
	kab := nik[0:4]
	kec := nik[0:6]

	var jk string
	var tglLahir int
	if tglInt > 40 {
		jk = "PEREMPUAN"
		tglLahir = tglInt - 40
	} else {
		jk = "LAKI-LAKI"
		tglLahir = tglInt
	}
	if tglLahir < 1 || tglLahir > 31 {
		return Data{}, errors.New("tanggal lahir hasil koreksi di luar rentang 1-31")
	}
	if (blnInt == 4 || blnInt == 6 || blnInt == 9 || blnInt == 11) && tglLahir > 30 {
		return Data{}, errors.New("tanggal tidak valid untuk bulan dengan 30 hari")
	}
	if blnInt == 2 && tglLahir > 29 {
		return Data{}, errors.New("tanggal tidak valid untuk Februari (maks 29)")
	}
	return Data{
		KodeProvinsi:    prov,
		KodeKabKota:     kab,
		KodeKecamatan:   kec,
		TanggalLahirStr: strconv.Itoa(tglLahir),
		BulanLahirStr:   blnStr,
		TahunLahirStr:   thnStr,
		UniqCode:        uniq,
		JenisKelamin:    jk,
	}, nil
}