package main

import (
	"TUGASBESAR/Admin"
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
		"1. Regristrasi\n",
		"2. Login\n",
		"3. Keluar\n",
	)
}
func main() {
	var T Regristrasi.Pengguna
	var r Regristrasi.AdminRegristrasi
	var n int
	var x string
	var berhasil bool

	tampilanAwal()
	for {
		Pilih()
		fmt.Scan(&x)

		if x == "1" {
			Regristrasi.BuatAkun(&T, &n)
		} else if x == "2" {
			fmt.Print("Masukan Usser : ")
			fmt.Scan(&r.Usser)
			fmt.Print("Masukan Password : ")
			fmt.Scan(&r.Password)

			berhasil = Regristrasi.Login(T, n, &r)

			if berhasil {
				Admin.FungsiAdmin(r)
			} else {
				fmt.Println("Login Anda Gagal")
				fmt.Println("Silahkan Coba Lagi")
			}
		} else if x == "3" {
			os.Exit(0)
		} else {
			fmt.Println("INPUT TIDAK VALID")
		}
	}
}
