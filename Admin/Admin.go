package Admin

import "fmt"

func MenuAdmin() {
	fmt.Print("\v",
		"============================================================================\n",
		"============================================================================\n",
		"--------------------------------- Menu Admin  ------------------------------\n",
		"----------------------------------------------------------------------------\n",
		"1. Tambah Buku \n",
		"2. Hapus Buku \n",
		"3. Lihat Member \n",
		"4. Kembali \n",
	)
}

func FungsiAdmin() {
	var n Buku
	var x string
	for {
		MenuAdmin()
		fmt.Scan(&x)
		if x == "1" {
			InputBuku(&n)
		} else if x == "2" {
			fmt.Println("Pilihan 2")
		} else if x == "3" {
			fmt.Println("Pilihan 3")
		} else if x == "4" {
			break
		} else {
			fmt.Println("tes salah input.")
		}
	}
}
