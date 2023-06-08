package Regristrasi

import "fmt"

const BanyakAdmin = 10000

type AdminRegristrasi struct {
	Nama     string
	Usser    string
	Password string
}
type Pengguna [BanyakAdmin]AdminRegristrasi

// func bubble_sort_1(A *[5]int, n int){
// 	var i, k int
// 	for k = 1; k <= n-1; k++{
// 		for i = 0;
// 	}
//}
// func selection_sort(A *[10]int, n int){
// 	var i pas
// }

func BuatAkun(T *Pengguna, n *int) {
	fmt.Print("\v",
		"=================================\n",
		"| ***  REGRISTRASI ADMIN  *** |\n",
		"=================================\n",
	)
	fmt.Print("Nama : ")
	fmt.Scan(&T[*n].Nama)
	fmt.Print("Buat Ussername : ")
	fmt.Scan(&T[*n].Usser)
	fmt.Print("Buat Password : ")
	fmt.Scan(&T[*n].Password)
	*n++
	fmt.Println("-----------------------------")
	fmt.Println("|    Regristasi Berhasil    |")
	fmt.Println("-----------------------------")
}

func Login(T Pengguna, n int, r *AdminRegristrasi) bool {
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
