package output

import "fmt"

func FormatTable(nik, prov, kab, kec, kodePos, tgl, bln, thn, jk, uniq string) {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%-20s : %s\n", "NIK", nik)
	fmt.Printf("%-20s : %s/%s/%s\n", "Tanggal Lahir", tgl, bln, thn)
	fmt.Printf("%-20s : %s\n", "Jenis Kelamin", jk)
	fmt.Printf("%-20s : %s\n", "Provinsi", prov)
	fmt.Printf("%-20s : %s\n", "Kab/Kota", kab)
	fmt.Printf("%-20s : %s\n", "Kecamatan", kec)
	fmt.Printf("%-20s : %s\n", "Kode Pos", kodePos)
	fmt.Printf("%-20s : %s\n", "Uniqcode", uniq)
	fmt.Println("--------------------------------------------------")
}