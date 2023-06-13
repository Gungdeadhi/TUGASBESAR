package Admin

import (
	"TUGASBESAR/PnK"
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
		"6. Informasi Member \n",
		"7. LogOut \n",
	)
}
func InformasiPeminjaman() {
	fmt.Print("\v",
		"==================================\n",
		"| ***  Informasi Peminjaman  *** |\n",
		"==================================\n",
		"1. Peminjaman Buku \n",
		"2. Pengembalian Buku \n",
		"3. Daftar Peminjam Buku \n",
		"4. Daftar Pengembalian Buku \n",
		"5. Keluar\n",
	)
}

func Member() {
	fmt.Print("\v",
		"===============================\n",
		"| ***  Informasi Member   *** |\n",
		"===============================\n",
		"1. Daftar Member\n",
		"2. Tambah Member \n",
		"3. Hapus Member \n",
		"4. Keluar \n",
	)
}

func FungsiAdmin(T *Pusat.TabBuku, r Regristrasi.AdminRegristrasi) {
	var N, i int
	var j string
	var k Pusat.Buku
	var x, Judul string
	var idx int
	// var A Regristrasi.Pengguna

	i = 0
	for i != 1000000 {
		MenuAdmin(r.Nama)
		fmt.Print("Pilih [1/2/3/4/5/6/7] ")
		fmt.Println("Pilihan : ")
		fmt.Scan(&x)
		switch x {
		case "1":
			Pusat.CetakBuku(*T, N)

		case "2":
			for j != "N" || x == "2" {
				Pusat.InputBuku(&k)
				idx = Pusat.CariBuku_Kode(*T, N, k.Kode)
				if idx == -1 {
					fmt.Print("Simpan Buku? [Y/N]: ")
					fmt.Scan(&x)
					if x == "y" || x == "Y" {
						Pusat.TambahBuku(T, &N, k)
					}
					fmt.Print("Tambah Buku Lagi? [Y/N]: ")
					fmt.Scan(&j)
				} else {
					fmt.Println("Kode Buku Ganda")
				}
			}
		case "3":
			Pusat.CetakBuku(*T, N)

			for j != "N" || x == "3" {
				fmt.Print("Masukan Judul Buku Yang Akan Diedit: ")
				fmt.Scan(&Judul)
				fmt.Println(" ")
				idx = Pusat.CariBuku_Nama(*T, N, Judul)
				if idx == i {
					Pusat.InputBuku(&k)
					fmt.Print("Simpan Buku? [Y/N] : ")
					fmt.Scan(&x)
					if x == "y" || x == "Y" {
						Pusat.EditBuku(T, N, Judul, k)
					}
					fmt.Print("Edit Buku lagi? [Y|N]")
					fmt.Scan(&j)
				} else {
					fmt.Println("TIDAK ADA BUKU")
				}

			}
		case "4":
			fmt.Print("Masukan Judul Buku Yang Akan Dihapus : ")
			fmt.Scan(&Judul)

			idx = Pusat.CariBuku_Nama(*T, N, Judul)
			if idx == i {
				fmt.Print("Buku akan dihapus? [Y|N] =  ")
				fmt.Scan(&j)
				if j == "Y" || j == "y" {
					Pusat.HapusBuku(T, N, Judul)
				} else if j == "N" || j == "n" {
					fmt.Println("Buku Tidak Terhapus ")
				}
			} else {
				fmt.Print("Buku Tidak Ditemukan")

			}

		case "5":
			FungsiInformasiPeminjaman()
		case "6":
			i = 1000000
		default:
			fmt.Println("INPUT TIDAK VALID")
		}
	}
}
func FungsiInformasiPeminjaman() {
	var r PnK.Pinjam
	var n PnK.TabPinjam
	var x string
	var i int

	i = 0
	for i != 1000000 {
		InformasiPeminjaman()
		fmt.Println("Pilih [1/2/3/4/5]")
		fmt.Print("Pilihan : ")
		fmt.Scan(&x)
		switch x {
		case "1":
			PnK.InputPinjamBuku(&r)
			PnK.TambahPeminjam(&n, &i, r)
		case "2":
			PnK.TambahPengembalian(&n, &i, r)
		case "3":
			PnK.CetakPeminjam(n, i)
		case "4":

		case "5":
			i = 1000000
		default:
			fmt.Println("INPUT TIDAK VALID")
		}
	}
}
func FungsiMember() {
	var i int
	var x string

	i = 100000
	for i != 100000 {
		Member()
		fmt.Println("Pilih [1/2/3/4]")
		fmt.Print("PIlihan : ")
		fmt.Scan(x)
		switch x {
		case "1":

		case "2":

		case "3":

		case "4":
			i = 100000
		default:
			fmt.Println("INPUT TIDAK VALID")
		}

	}
}
