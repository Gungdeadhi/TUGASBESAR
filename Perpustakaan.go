package main

import (
	"TUGASBESAR/Admin"
	"TUGASBESAR/Member"
	"TUGASBESAR/Pusat"
	"fmt"
)

func tampilanAwal() {
	fmt.Print("\v",
		"========================================================================\n",
		"=============================== SELAMAT DATANG =========================\n",
		"========================================================================\n",
		"========================================================================\n",
	)
}

func Pilih() {
	fmt.Print("\v",
		"========================================================================\n",
		"***             ***      PILIH MENU UNTUK LOGIN       ***            ***\n",
		"========================================================================\n",
		"1. Regristrasi\n",
		"2. Login\n",
		"3. Admin\n",
		"4. Keluar\n",
	)
}
func main() {
	var T Pusat.Member
	var r Pusat.Registrasi
	var n int
	var x string
	var berhasil bool

	tampilanAwal()
	for {
		Pilih()
		fmt.Scan(&x)

		if x == "1" {
			Pusat.BuatAkun(&T, &n)
		} else if x == "2" {
			fmt.Print("Masukan Usser : ")
			fmt.Scan(&r.Usser)
			fmt.Print("Masukan Password : ")
			fmt.Scan(&r.Password)

			berhasil = Pusat.Login(T, n, &r)

			if berhasil {
				Member.FungsiMember(r)
			} else {
				fmt.Println("Login Anda Gagal")
				fmt.Println("Silahkan Coba Lagi")
			}
		} else if x == "3" {
			Admin.FungsiAdmin()
		} else if x == "4" {
			break
		}
	}
}
