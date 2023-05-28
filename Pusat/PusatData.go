package Pusat

import "fmt"

const NMAX = 10000

type Registrasi struct {
	Nama     string
	Usser    string
	Password string
}
type Member [NMAX]Registrasi

// func bubble_sort_1(A *[5]int, n int){
// 	var i, k int
// 	for k = 1; k <= n-1; k++{
// 		for i = 0;
// 	}
//}
// func selection_sort(A *[10]int, n int){
// 	var i pas
// }

func BuatAkun(T *Member, n *int) {
	fmt.Print("Nama : ")
	fmt.Scan(&T[*n].Nama)
	fmt.Print("Buat Ussername : ")
	fmt.Scan(&T[*n].Usser)
	fmt.Print("Buat Password : ")
	fmt.Scan(&T[*n].Password)
	*n++
	fmt.Println("Registrasi Berhasil")
	for i := 1; i < *n-1; i++ {
		if T[i].Usser == T[i-1].Usser && T[i].Password == T[i-1].Password {
			fmt.Print("Username dan password sudah ada")
		}
	}
}

func Login(T Member, n int, r *Registrasi) bool {
	var percobaan int = 3

	for percobaan > 0 {
		for i := 0; i < n; i++ {
			if T[i].Usser == r.Usser && T[i].Password == r.Password {
				r.Nama = T[i].Nama
				return true
			}
		}

		percobaan--
	}

	return false
}
