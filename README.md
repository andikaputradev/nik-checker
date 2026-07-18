# NIK Checker Premium

[![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)

Aplikasi CLI untuk memvalidasi dan menampilkan informasi lengkap dari **Nomor Induk Kependudukan (NIK)** KTP Indonesia. Pengguna memasukkan NIK, dan program menampilkan data kependudukan dalam format tabel.

## Fitur

- **Input interaktif**: NIK 16 digit dimasukkan langsung melalui terminal.
- **Validasi ketat**: pemeriksaan panjang input, format digit, rentang bulan (01 sampai 12), serta validitas tanggal lahir (maksimal 31, aturan bulan 30 hari, Februari maksimal 29).
- **Koreksi otomatis**: apabila kode tanggal lebih besar dari 40, jenis kelamin ditetapkan perempuan dan tanggal lahir dikurangi 40.
- **Lookup wilayah offline**: kode provinsi, kabupaten/kota, kecamatan, dan kode pos diambil dari `data.json`.
- **Output tabel ringkas**: seluruh informasi kependudukan ditampilkan dalam format tabel yang terstruktur.

## Struktur Proyek

```
nik-checker/
├── cmd/nik-checker/main.go      Titik masuk program
├── internal/
│   ├── data/loader.go           Pemuat dan pencari data wilayah
│   ├── nik/parser.go            Validasi dan parsing NIK
│   └── output/formatter.go      Render tabel (dan JSON opsional)
├── data.json                    Data wilayah lengkap
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## Instalasi

```bash
git clone https://github.com/andikaputradev/nik-checker.git
cd nik-checker
go build -o nik-checker ./cmd/nik-checker/
```

## Penggunaan

Jalankan program, kemudian masukkan NIK saat diminta.

```bash
./nik-checker
Input NIK: 3273014106960002
```

Program akan menampilkan output sebagaimana contoh berikut.

## Contoh Output

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

## Detail Teknis

**Koreksi gender dan tanggal lahir**
Digit ke-7 dan ke-8 NIK merepresentasikan tanggal (DD). Apabila nilai DD lebih besar dari 40, jenis kelamin ditetapkan perempuan dan tanggal lahir sebenarnya dihitung sebagai DD dikurangi 40. Contoh: nilai 41 diinterpretasikan sebagai perempuan dengan tanggal lahir 1.

**Validasi kalender**
Bulan harus berada pada rentang 01 sampai 12. Hari tidak boleh melebihi batas bulan terkait (April, Juni, September, dan November maksimal 30 hari; Februari maksimal 29 hari).

**Lookup wilayah tanpa alokasi besar**
Data wilayah disimpan sebagai byte slice mentah. Pencarian menggunakan `jsonparser.GetString` tanpa membentuk map penuh, sehingga penggunaan memori tetap rendah.

## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE). Data wilayah (`data.json`) bersumber dari data publik pemerintah Indonesia untuk keperluan edukasi.

## Kontak

<div align="center">

[![Email](https://img.shields.io/badge/Gmail-D14836?style=flat-square&logo=gmail&logoColor=white)](mailto:wahyuandikaputra.co.id@gmail.com)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat-square&logo=linkedin&logoColor=white)](https://linkedin.com/in/wahyu-andika-putra)
[![Portfolio](https://img.shields.io/badge/andikaputra.dev-000000?style=flat-square&logo=safari&logoColor=white)](https://andikaputra.dev)
[![Instagram](https://img.shields.io/badge/Instagram-E4405F?style=flat-square&logo=instagram&logoColor=white)](https://www.instagram.com/w.andikaputraa)
[![WhatsApp](https://img.shields.io/badge/WhatsApp-25D366?style=flat-square&logo=whatsapp&logoColor=white)](https://wa.me/6285728337030)
[![Indonesia](https://img.shields.io/badge/Indonesia-CE1126?style=flat-square&logo=googlemaps&logoColor=white)](#)

</div>

Dipersembahkan oleh **Wahyu Andika Putra**.
