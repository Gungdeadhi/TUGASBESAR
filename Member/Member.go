package Member

import (
	"TUGASBESAR/Pusat"
	"fmt"
)

type Anggota struct {
}

func MenuMember(nama string) {
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("----------------------------- Menu Utama -------------------------------")
	fmt.Printf("----------------------------- Hallo, %s --------------------------------\n", nama)
	fmt.Println("1. Data Buku")
	fmt.Println("2. Pinjam Buku")
	fmt.Println("3. Informasi Peminjaman")
	fmt.Println("4. Kembali")
}

func FungsiMember(r Pusat.Registrasi) {
	var p string

	for {
		MenuMember(r.Nama)
		fmt.Scan(&p)

		if p == "1" {
			fmt.Println("pilihan 1")
		} else if p == "2" {
			fmt.Println("pilihan 2")
		} else if p == "3" {
			fmt.Println("pilihan 3")
		} else if p == "4" {
			break
		} else {
			fmt.Println("input pilihan salah.")
		}
	}
}
