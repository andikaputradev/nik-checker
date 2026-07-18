# NIK Checker Premium

[![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)
[![Release](https://img.shields.io/badge/Release-v1.0.0-2088FF?style=flat-square&logo=github)](https://github.com/wahyuandikaputra/nik-checker/releases)

Aplikasi **CLI** berbahasa Indonesia untuk validasi, parsing, dan penelusuran informasi lengkap **Nomor Induk Kependudukan (NIK)** KTP Indonesia.  
Dirancang dengan arsitektur **modular**, **zero-allocation parsing**, dan **lookup data wilayah** berbasis `jsonparser` tanpa *unmarshaling* penuh.

## ✨ Fitur

- Validasi ketat 16 digit (panjang, numerik, rentang bulan & tanggal lahir).
- Koreksi otomatis tanggal lahir dan jenis kelamin berdasarkan aturan kode NIK (DD>40).
- Pencarian wilayah offline: provinsi, kab/kota, kecamatan, kode pos.
- Dukungan output **tabel** ringkas atau **JSON**.
- Mode **batch** untuk memproses ribuan NIK dari berkas teks.
- Memory-safe, *production-ready*.

## 📁 Struktur Proyek

```
nik-checker/
├── cmd/nik-checker/main.go      # Orkestrator: mode interaktif, batch, flag
├── internal/
│   ├── data/loader.go           # Singleton loader & lookup wilayah
│   ├── nik/parser.go            # Pure function validasi & parsing NIK
│   └── output/formatter.go      # Render tabel & JSON
├── data.json                    # Data wilayah (provinsi, kab/kota, kecamatan)
├── go.mod / go.sum
├── LICENSE
└── README.md
```

## ⚙️ Instalasi

```bash
git clone https://github.com/andikaputradev/nik-checker.git
cd nik-checker
go build -o nik-checker ./cmd/nik-checker/
```

## 🚀 Penggunaan

**Mode Interaktif:**
```bash
./nik-checker -data data.json
```

**Mode Batch:**
```bash
./nik-checker -data data.json -input daftar_nik.txt -format json > hasil.json
```

**Flag:**
| Flag | Default | Keterangan |
|------|---------|------------|
| `-data` | `data.json` | Path berkas data wilayah |
| `-input` | - | Path berkas daftar NIK (mode batch) |
| `-format` | `table` | Format output: `table` atau `json` |


## 📋 Contoh Output

**Tabel**:
```
--------------------------------------------------
NIK                  : 3273014106960002
Tanggal Lahir        : 1/06/96
Jenis Kelamin        : PEREMPUAN
Provinsi             : JAWA BARAT
Kab/Kota             : KOTA BANDUNG
Kecamatan            : SUKASARI
Kode Pos             : 40151
Uniqcode             : 0002
--------------------------------------------------
```

**JSON**:
```json
{"nik":"3273014106960002","tanggal_lahir":"1/06/96","jenis_kelamin":"PEREMPUAN","provinsi":"JAWA BARAT","kab_kota":"KOTA BANDUNG","kecamatan":"SUKASARI","kode_pos":"40151","uniqcode":"0002"}
```

## 🧠 Teknis Singkat

- **Koreksi Gender & Tanggal**: `DD = digit ke-7&8`, jika >40 → perempuan, lahir = DD-40. Validasi bulan (01-12) dan hari sesuai bulan.
- **Lookup Wilayah**: `jsonparser.GetString` langsung ke byte slice, tanpa alokasi map penuh.
- **Keamanan**: slicing aman, validasi panjang, nol `unsafe`, error tertangani eksplisit.

## 📦 GitHub Repo

**Deskripsi**:  
> CLI tool untuk validasi, parsing, dan pelacakan NIK KTP Indonesia dengan lookup wilayah terintegrasi. Dibangun dengan Go, arsitektur modular, siap produksi.

**Tags**: `go` `golang` `nik-checker` `ktp` `indonesia` `cli` `validation` `data-validation` `json` `security` `open-source` `identity` `tools`

## 📜 Lisensi

MIT License. Lihat [LICENSE](LICENSE). Data wilayah (`data.json`) berasal dari data publik pemerintah Indonesia untuk keperluan edukasi.

## 📬 Kontak

<div align="center">

[![Email](https://img.shields.io/badge/Gmail-D14836?style=flat-square&logo=gmail&logoColor=white)](mailto:wahyuandikaputra.co.id@gmail.com)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat-square&logo=linkedin&logoColor=white)](https://linkedin.com/in/wahyu-andika-putra)
[![Portfolio](https://img.shields.io/badge/andikaputra.dev-000000?style=flat-square&logo=safari&logoColor=white)](https://andikaputra.dev)
[![Instagram](https://img.shields.io/badge/Instagram-E4405F?style=flat-square&logo=instagram&logoColor=white)](https://www.instagram.com/w.andikaputraa)
[![WhatsApp](https://img.shields.io/badge/WhatsApp-25D366?style=flat-square&logo=whatsapp&logoColor=white)](https://wa.me/6285728337030)
[![Indonesia](https://img.shields.io/badge/Indonesia-CE1126?style=flat-square&logo=googlemaps&logoColor=white)](#)

</div>

*Dipersembahkan oleh*  
**Wahyu Andika Putra**
