package main

import (
	"TUGASBESAR/Pusat"
	"fmt"
)

func tampilanAwal() {
	fmt.Print("\v",
		"========================================================================\n",
		"=============================== SELAMAT DATANG =========================\n",
		"========================================================================\n",
		"========================================================================\n",
		"========================================================================\n",
	)
}

func MenuUtama(nama string) {
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("----------------------------- Menu Utama -------------------------------")
	fmt.Printf("----------------------------- Hallo, %s --------------------------------\n", nama)
	fmt.Println("1. Data Buku")
	fmt.Println("2. Tambah Buku")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Informasi Peminjaman ")
	fmt.Println("5. ")
	fmt.Println("6. ")
}

func main() {
	var u, p string
	tampilanAwal()
	if Pusat.Login(&u, &p) {
		MenuUtama(u)
	} else {
		fmt.Println("Login Anda Gagal")
	}
}
