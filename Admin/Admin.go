package Admin

import (
	"TUGASBESAR/Member"
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
		"6. Data Member\n",
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

func MenuMember() {
	fmt.Print("\v",
		"===============================\n",
		"| ***  Informasi Member   *** |\n",
		"===============================\n",
		"1. Daftar Member\n",
		"2. Tambah Member \n",
		"3. Edit Member \n",
		"4. Hapus Member \n",
		"5. Keluar \n",
	)
}

func FungsiAdmin(T *Pusat.TabBuku, r Regristrasi.AdminRegristrasi, P1 *PnK.TabPinjam, P2 *PnK.TabPengembali, H *Member.TabMember, N1, N2, NPinjam, NKembali *int) {
	var i int
	var j string
	var k Pusat.Buku
	var x, Judul string
	var idx int

	i = 0
	for i != 1000000 {
		MenuAdmin(r.Nama)
		fmt.Println("Pilih [1/2/3/4/5/6/7] ")
		fmt.Print("Pilihan : ")
		fmt.Scan(&x)
		switch x {
		case "1":
			if *N1 > 0 {
				Pusat.CetakBuku(*T, *N1)
			} else {
				fmt.Println("----------------------------------------------")
				fmt.Println("Data Buku Kosong, Silahkan Tambahkan Data Buku")
				fmt.Println("----------------------------------------------")
			}

		case "2":
			for j != "N" || x == "2" {
				Pusat.InputBuku(&k)
				idx = Pusat.CariBuku_Kode(*T, *N1, k.Kode)
				if idx == -1 {
					fmt.Print("Simpan Buku? [Y/N]: ")
					fmt.Scan(&x)
					if x == "y" || x == "Y" {
						Pusat.TambahBuku(T, N1, k)
						Pusat.SelSort(T, *N1)
					}
					fmt.Print("Tambah Buku Lagi? [Y/N]: ")
					fmt.Scan(&j)
				} else {
					fmt.Println("Kode Buku Ganda")
				}
			}
		case "3":
			Pusat.CetakBuku(*T, *N1)

			for j != "N" || x == "3" {
				fmt.Print("Masukan Judul Buku Yang Akan Diedit: ")
				fmt.Scan(&Judul)
				fmt.Println(" ")
				idx = Pusat.CariBuku_Nama(*T, *N1, Judul)
				if idx != -1 {
					Pusat.InputBuku(&k)
					fmt.Print("Simpan Buku? [Y/N] : ")
					fmt.Scan(&x)
					if x == "y" || x == "Y" {
						Pusat.EditBuku(T, *N1, Judul, k)
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

			idx = Pusat.CariBuku_Nama(*T, *N1, Judul)
			if idx != -1 {
				fmt.Print("Buku akan dihapus? [Y|N] : ")
				fmt.Scan(&j)
				if j == "Y" || j == "y" {
					Pusat.HapusBuku(T, N1, Judul)
				} else if j == "N" || j == "n" {
					fmt.Println("Buku Tidak Terhapus ")
				}
			} else {
				fmt.Print("Buku Tidak Ditemukan")

			}

		case "5":
			FungsiInformasiPeminjaman(T, P1, P2, H, N1, N2, NPinjam, NKembali)

		case "6":
			FungsiMember(H, N2)
		case "7":
			i = 1000000
		default:
			fmt.Println("INPUT TIDAK VALID")
		}
	}
}
func FungsiInformasiPeminjaman(T *Pusat.TabBuku, P1 *PnK.TabPinjam, P2 *PnK.TabPengembali, H *Member.TabMember, N1, N2, NPinjam, NKembali *int) {
	var z Member.Member
	var r PnK.Pinjam
	var y PnK.Pengembali
	var i, idx1, idx2 int
	var x string

	i = 0
	for i != 1000000 {
		InformasiPeminjaman()
		fmt.Println("Pilih [1/2/3/4/5]")
		fmt.Print("Pilihan : ")
		fmt.Scan(&x)
		switch x {
		case "1":
			Pusat.CetakBuku(*T, *N1)
			Member.CetakMember(*H, *N2)
			fmt.Print("Masukan Nama Member : ")
			fmt.Scan(&z.Nama)
			idx1 = Member.CariMember(*H, *N2, z.Nama)
			if idx1 != -1 {
				PnK.InputPinjamBuku(&r)
				idx2 = Pusat.CariBuku_Nama(*T, *N1, r.Judul)
				if idx2 != -1 && T[idx2].JmlBuku > 0 {
					PnK.TambahPeminjam(P1, NPinjam, r)
					T[idx2].JmlBuku--

				} else {
					fmt.Println("--------------------------------------------------------------------")
					fmt.Println("Stok Buku Habis / Buku Belum Tersedia, Silahkan Pilih Buku Yang Lain")
					fmt.Println("--------------------------------------------------------------------")
				}
			} else {
				fmt.Println("-----------------------")
				fmt.Println("Nama Member Tidak Valid")
				fmt.Println("-----------------------")
			}
		case "2":
			PnK.CetakPeminjam(*P1, *NPinjam)
			PnK.InputPengembalianBuku(&y)
			idx2 = PnK.CariPeminjamBuku(*P1, *NPinjam, y.Nama)
			if idx2 != -1 {
				PnK.TambahPengembalian(P2, NKembali, y)
				idx1 = Pusat.CariBuku_Nama(*T, *N1, r.Judul)
				T[idx1].JmlBuku++
			} else {
				fmt.Println("----------------------------------------------------")
				fmt.Println("Daftar peminjam dengan nama tersebut tidak tersedia")
				fmt.Println("----------------------------------------------------")
			}
		case "3":
			if *NPinjam > 0 {
				PnK.CetakPeminjam(*P1, *NPinjam)
			} else {
				fmt.Println("-----------------------")
				fmt.Println("Tidak Ada Peminjam Buku")
				fmt.Println("-----------------------")
			}
		case "4":
			if *NKembali > 0 {
				PnK.CetakPengembalian(*P1, *P2, *NKembali)
			} else {
				fmt.Println("---------------------------")
				fmt.Println("Tidak Ada Pengembalian Buku")
				fmt.Println("---------------------------")
			}
		case "5":
			i = 1000000
		default:
			fmt.Println("INPUT TIDAK VALID")
		}
	}
}

func FungsiMember(T *Member.TabMember, N *int) {
	var r Member.Member
	var x, y, Nama string
	var i int

	i = 0
	for i != 100000 {
		MenuMember()
		fmt.Println("Pilih [1/2/3/4/5]")
		fmt.Print("Pilihan : ")
		fmt.Scan(&x)
		switch x {
		case "1":
			if *N > 0 {
				Member.CetakMember(*T, *N)
			} else {
				fmt.Println("------------------")
				fmt.Println("Data Member Kosong")
				fmt.Println("------------------")
			}
		case "2":
			for y != "N" || x == "2" {
				Member.InputMember(&r)
				idx := Member.CariMember(*T, *N, r.Nama)
				if idx == -1 {
					fmt.Print("Simpan Member? [Y|N] ")
					fmt.Scan(&x)
					if x == "Y" || x == "y" {
						Member.TambahMember(T, N, r)
					} else {
						fmt.Println("Member Tidak Tersimpan")
					}
					fmt.Print("Tambah member Lagi? [Y|N] ")
					fmt.Scan(&y)
				} else {
					fmt.Println("Member telah Terdaftar")
				}
			}
		case "3":
			for y != "N" || x == "3" {
				Member.CetakMember(*T, *N)
				fmt.Print("Masukan Nama Member Yang Akan Diedit : ")
				fmt.Scan(&Nama)
				idx := Member.CariMember(*T, *N, Nama)
				if idx != -1 {
					Member.InputMember(&r)
					fmt.Print("Simpan? [Y|N]")
					fmt.Scan(&x)
					if x == "Y" || x == "y" {
						Member.EditMember(T, *N, Nama, r)
						fmt.Print("Edit Member Lagi? [Y|N]")
						fmt.Scan(&y)
					} else {
						fmt.Print("Pembaruan Tidak Tersimpan")
					}
				} else {
					fmt.Println("---------------------")
					fmt.Println("Member Tidak Tersedia")
					fmt.Println("---------------------")
				}

			}
		case "4":
			for y != "N" || x == "4" {
				Member.CetakMember(*T, *N)
				fmt.Print("Masukan Member Yang Akan Dihapus : ")
				fmt.Scan(&r.Nama)
				idx := Member.CariMember(*T, *N, r.Nama)
				if idx != -1 {
					fmt.Print("Member Akan Dihapus? [Y|N] ")
					fmt.Scan(&x)
					if x == "Y" || x == "y" {
						Member.HapusMember(T, N, r.Nama)
						fmt.Print("Hapus Member Lagi ? [Y|N] ")
						fmt.Scan(&y)
					} else {
						fmt.Print("Member Tidak Terhapus")
					}
				} else {
					fmt.Println("--------------------------")
					fmt.Println("Data Member Tidak Tersedia")
					fmt.Println("--------------------------")
				}
			}
		case "5":
			i = 100000
		default:
			fmt.Println("INPUT TIDAK VALID")
		}

	}

}
