package main

import (
	"TUGASBESAR/Admin"
	"TUGASBESAR/Pusat"
	"TUGASBESAR/Regristrasi"
	"fmt"
	"os"
)

func tampilanAwal() {
	fmt.Print("\v",
		"=========================================================================\n",
		"||****                  **** SELAMAT DATANG ****                   ****||\n",
		"||         *****        ------------------------         *****         ||\n",
		"=========================================================================\n",
	)
}

func Pilih() {
	fmt.Print("\v",
		"========================================================\n",
		"||***      ***          PILIH MENU         ***      ***||\n",
		"========================================================\n",
		"1. Registrasi\n",
		"2. Login\n",
		"3. Keluar\n",
	)
	fmt.Println(" ")
}
func main() {
	var A Pusat.TabBuku
	var T Regristrasi.Pengguna
	var r Regristrasi.AdminRegristrasi
	var n, percobaan int
	var x string
	var berhasil bool

	tampilanAwal()
	for {
		Pilih()
		fmt.Println("Pilih [1/2/3]")
		fmt.Print("pilihan : ")
		fmt.Scan(&x)

		if x == "1" {
			fmt.Print("\v",
				"=================================\n",
				"|  ***  REGISTRASI  ADMIN  ***  |\n",
				"=================================\n",
			)

			fmt.Print("Nama 		: ")
			fmt.Scan(&r.Nama)
			fmt.Print("Buat Ussername 	: ")
			fmt.Scan(&r.Usser)
			fmt.Print("Buat Password 	: ")
			fmt.Scan(&r.Password)
			for Regristrasi.Cek(T, n, r) {
				Regristrasi.Gagal()
				fmt.Print("Nama 	: ")
				fmt.Scan(&r.Nama)
				fmt.Print("Buat Ussername 	: ")
				fmt.Scan(&r.Usser)
				fmt.Print("Buat Password	 : ")
				fmt.Scan(&r.Password)
			}
			Regristrasi.BuatAkun(&T, &n, r)
			Regristrasi.Berhasil()
		} else if x == "2" {
			percobaan = 3

			for percobaan > 0 && !berhasil {
				fmt.Print("Masukan User 		: ")
				fmt.Scan(&r.Usser)
				fmt.Print("Masukan Password 	: ")
				fmt.Scan(&r.Password)

				berhasil = Regristrasi.Login(T, n, &r)

				if !berhasil {
					fmt.Println("--------------------------------")
					fmt.Println("| Ussername Atau Pasword Salah |")
					fmt.Println("| Silahkan Coba Lagi           |")
					fmt.Println("--------------------------------")
				}

				percobaan--
			}

			if berhasil {
				Admin.FungsiAdmin(&A, r)
			} else {
				fmt.Println("----------------")
				fmt.Println("Login Anda Gagal")
				fmt.Println("----------------")
			}
		} else if x == "3" {
			os.Exit(0)
		} else {
			fmt.Println("INPUT TIDAK VALID")
		}
	}
}
