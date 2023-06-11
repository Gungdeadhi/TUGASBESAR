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
		"3. Hapus Buku\n",
		"4. Informasi Peminjaman \n",
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

func FungsiAdmin(T *Pusat.TabBuku, r Regristrasi.AdminRegristrasi) {
	var i, N int
	var j string
	var n Pusat.Buku
	var x, Judul string
	var A Regristrasi.Pengguna

	i = 0
	for i != 1000000 {
		MenuAdmin(r.Nama)
		fmt.Scan(&x)
		switch x {
		case "1":
			Pusat.CetakBuku(*T, i)
		case "2":
			fmt.Println("Masukan Passowrd : ")
			fmt.Scan(&j)
			for j == "y" || j == "Y" || j == "Admin" {
				Pusat.InputBuku(&n)
				if Pusat.CariBuku_Kode(*T, N, n.Kode) {
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
			Pusat.DaftarMember(&A, N)
		case "6":
			i = 1000000
		default:
			fmt.Println("INPUT TIDAK VALID")
		}
	}
}
func FungsiInformasiPeminjaman() {
	var n Pusat.Member
	var r Pusat.Buku
	var x string
	var i int

	i = 0
	for i != 1000000 {
		InformasiPeminjaman()
		fmt.Scan(&x)
		switch x {
		case "1":
			Pusat.PinjamBuku(&n, &r)
		case "2":
			Pusat.PengembalianBuku(&n, &r)
		case "3":
			i = 1000000
		default:
			fmt.Println("INPUT TIDAK VALID")
		}
	}
}
