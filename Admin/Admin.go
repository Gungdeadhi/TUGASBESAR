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
		"2. Tambah Buku \n",
		"3. Edit Data Buku \n",
		"4. Hapus Buku\n",
		"5. Informasi Peminjaman \n",
		"6. Daftar Member\n",
		"7. Kembali \n",
		"------------------------------------\n",
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
		"----------------------------------\n",
	)
}

func FungsiAdmin(T *Pusat.TabBuku, r Regristrasi.AdminRegristrasi) {
	var i, N int
	var j string
	var n Pusat.Buku
	var x, Judul string
	// var A Regristrasi.Pengguna

	i = 0
	for i != 1000000 {
		MenuAdmin(r.Nama)
		fmt.Scan(&x)
		switch x {
		case "1":

			Pusat.CetakBuku(*T, i)
			// fmt.Println("Cari Buku? [Y/N]")
			// fmt.Scan(&j)
			// for j == "Y" || j == "y" {
			// 	fmt.Scan(&n)

			// }
		case "2":
			fmt.Println("Masukan Passowrd : ")
			fmt.Scan(&j)
			for j == "y" || j == "Y" || j == "Admin" {
				Pusat.InputBuku(&n)
				for Pusat.CariBuku_Kode(*T, N, n.Kode) {
					fmt.Println("KODE BUKU GANDA")
					Pusat.InputBuku(&n)
				}
				Pusat.TambahBuku(T, n, &i)
				fmt.Println("Tambah Buku Lagi? [Y/N]")
				fmt.Scan(&j)
			}
		case "3":
			FungsiInformasiPeminjaman()
		case "4":
			Pusat.HapusBuku(T, &N, Judul)
		case "5":
			Pusat.CetakBuku(*T, i)

			fmt.Print("Masukan judul buku : ")
			fmt.Scan(&n.Judul)

			Pusat.CariBuku_Nama(*T, N, Judul)

			if Pusat.CariBuku_Nama(*T, N, Judul) == 1 {

			}
		case "6":
			// Pusat.DaftarMember(&A, N)
		case "7":
			i = 1000000
		default:
			fmt.Println("INPUT TIDAK VALID")
		}
	}
}
func FungsiInformasiPeminjaman() {
	// var n Pusat.Member
	// var r Pusat.Buku
	var x string
	var i int

	i = 0
	for i != 1000000 {
		InformasiPeminjaman()
		fmt.Scan(&x)
		switch x {
		case "1":
			// Pusat.PinjamBuku(&n, &r)
		case "2":
			// Pusat.PengembalianBuku(&n, &r)
		case "3":
			i = 1000000
		default:
			fmt.Println("INPUT TIDAK VALID")
		}
	}
}
