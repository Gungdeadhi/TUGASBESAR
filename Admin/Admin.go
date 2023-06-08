package Admin

import (
	"TUGASBESAR/Pusat"
	"TUGASBESAR/Regristrasi"
	"fmt"
)

func MenuAdmin(n string) {

	fmt.Print("\v",
		"===================================\n",
		"|-------------- Menu -------------|\n",
		"===================================\n",
	)
	fmt.Printf(" *** Selamat Datang, %s *** ", n)
	fmt.Println(" ")
	fmt.Print("\v",
		"1. Daftar Buku \n",
		"2. Informasi Peminjaman \n",
		"3. Tambah Buku \n",
		"4. Hapus Buku\n",
		"5. Daftar Member\n",
		"6. Kembali \n",
	)
}
func InformasiPeminjaman() {
	fmt.Print("\v",
		"==================================\n",
		"| ***  Informasi Peminjaman  *** |\n",
		"==================================\n",
		"1. Peminjaman Buku \n",
		"2. Pengembalian Buku \n",
		"3. Keluar\n",
	)
}

func FungsiAdmin(r Regristrasi.AdminRegristrasi) {
	var i int
	var n Pusat.Buku
	var x string

	i = 0
	for i != 1000000 {
		MenuAdmin(r.Nama)
		fmt.Scan(&x)
		switch x {
		case "1":
			fmt.Println("1")
		case "2":
			FungsiPeminjaman()
		case "3":
			Pusat.InputBuku(&n)
		case "4":
			fmt.Println("4")
		case "5":
			fmt.Println("5")
		case "6":
			i = 1000000
		default:
			fmt.Println("INPUT TIDAK VALID")
		}
	}
}
func FungsiPeminjaman() {
	var x string
	var i int

	i = 0
	for i != 1000000 {
		InformasiPeminjaman()
		fmt.Scan(&x)
		switch x {
		case "1":
			fmt.Println("1")
		case "2":
			fmt.Println("2")
		case "3":
			i = 1000000
		default:
			fmt.Println("INPUT TIDAK VALID")
		}
	}
}
