# NIK Checker Premium

[![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)

Aplikasi CLI untuk memvalidasi dan menampilkan informasi lengkap dari **Nomor Induk Kependudukan (NIK)** KTP Indonesia.  
Pengguna cukup memasukkan NIK, dan program langsung menampilkan data kependudukan dalam format tabel yang informatif.

---

## ✨ Fitur

- **Input interaktif** – masukkan NIK 16 digit langsung di terminal.
- **Validasi ketat** – cek panjang, hanya digit, rentang bulan (01–12), serta tanggal lahir (maks 31, aturan 30 hari, Februari maks 29).
- **Koreksi otomatis** – jika kode tanggal > 40, jenis kelamin menjadi perempuan dan tanggal lahir dikurangi 40.
- **Lookup wilayah offline** – kode provinsi, kabupaten/kota, kecamatan, dan kode pos diambil dari `data.json`.
- **Output tabel ringkas** – semua informasi kependudukan ditampilkan dalam format tabel yang mudah dibaca.

---

## 📁 Struktur Proyek

```
nik-checker/
├── cmd/nik-checker/main.go      # Titik masuk program
├── internal/
│   ├── data/loader.go           # Pemuat & pencari data wilayah
│   ├── nik/parser.go            # Validasi & parsing NIK
│   └── output/formatter.go      # Render tabel (dan JSON opsional)
├── data.json                    # Data wilayah lengkap
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

---

## ⚙️ Instalasi

```bash
git clone https://github.com/andikaputradev/nik-checker.git
cd nik-checker
go build -o nik-checker ./cmd/nik-checker/
```

---

## 🚀 Penggunaan

Jalankan program, masukkan NIK saat diminta:

```bash
./nik-checker
Input NIK: 3273014106960002
```

Program akan langsung menampilkan output seperti di bawah ini.

---

## 📋 Contoh Output

```
Input NIK: 3273014106960002
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

---

## 🧠 Teknis Singkat

- **Koreksi gender & tanggal lahir**  
  Digit ke-7 dan ke-8 NIK adalah `DD`. Jika `DD > 40`, maka **perempuan** dan tanggal lahir sebenarnya = `DD - 40`. Contoh: `41` → perempuan, tanggal `1`.

- **Validasi kalender**  
  Bulan harus 01–12. Hari tidak boleh melebihi batas bulan (April, Juni, September, November maks 30; Februari maks 29).

- **Lookup wilayah tanpa alokasi besar**  
  Data wilayah disimpan sebagai byte slice mentah. Pencarian menggunakan `jsonparser.GetString` tanpa membuat map penuh, menjaga memori tetap rendah.

---

## 📦 GitHub Repository

**Deskripsi (Description)**:  
> CLI tool untuk validasi dan penelusuran NIK KTP Indonesia. Masukkan NIK, dapatkan tanggal lahir, jenis kelamin, provinsi, kab/kota, kecamatan, dan kode pos secara instan.

**Tags (Topics)**:  
`go` `golang` `nik-checker` `ktp` `indonesia` `cli` `validation` `data-validation` `json` `open-source` `identity`

---

## 📜 Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).  
Data wilayah (`data.json`) bersumber dari data publik pemerintah Indonesia untuk keperluan edukasi.

---

## 📬 Kontak

<div align="center">

[![Email](https://img.shields.io/badge/Gmail-D14836?style=flat-square&logo=gmail&logoColor=white)](mailto:wahyuandikaputra.co.id@gmail.com)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat-square&logo=linkedin&logoColor=white)](https://linkedin.com/in/wahyu-andika-putra)
[![Portfolio](https://img.shields.io/badge/andikaputra.dev-000000?style=flat-square&logo=safari&logoColor=white)](https://andikaputra.dev)
[![Instagram](https://img.shields.io/badge/Instagram-E4405F?style=flat-square&logo=instagram&logoColor=white)](https://www.instagram.com/w.andikaputraa)
[![WhatsApp](https://img.shields.io/badge/WhatsApp-25D366?style=flat-square&logo=whatsapp&logoColor=white)](https://wa.me/6285728337030)
[![Indonesia](https://img.shields.io/badge/Indonesia-CE1126?style=flat-square&logo=googlemaps&logoColor=white)](#)

</div>

---

*Dipersembahkan oleh*  
**Wahyu Andika Putra**