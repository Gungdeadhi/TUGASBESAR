package Regristrasi

import "fmt"

const BanyakAdmin = 10000

type AdminRegristrasi struct {
	Nama     string
	Usser    string
	Password string
}
type Pengguna [BanyakAdmin]AdminRegristrasi

func BuatAkun(T *Pengguna, n *int, r AdminRegristrasi) {
	T[*n] = r
	*n++
}

func Gagal() {
	fmt.Println("------------------------------------------------")
	fmt.Println("|   USSERNAME ATAU PASSWORD TELAH DIGUNAKAN    |")
	fmt.Println("------------------------------------------------")
}

func Berhasil() {
	fmt.Println("-----------------------------")
	fmt.Println("|    Regristasi Berhasil    |")
	fmt.Println("-----------------------------")
}

func Cek(T Pengguna, n int, r AdminRegristrasi) bool {
	for i := 0; i < n; i++ {
		if T[i].Usser == r.Usser || T[i].Password == r.Password {
			return true
		}
	}

	return false
}

func Login(T Pengguna, n int, r *AdminRegristrasi) bool {
	for i := 0; i < n; i++ {
		if T[i].Usser == r.Usser && T[i].Password == r.Password {
			r.Nama = T[i].Nama
			return true
		}
	}

	return false
}
