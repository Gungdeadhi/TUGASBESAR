package Pusat

import (
	"TUGASBESAR/Regristrasi"
	"fmt"
)

const NmaxBuku = 10000

type Buku struct {
	Judul, Pengarang, Kode string
	T_Terbit, JmlBuku      int
}

type TabBuku [NmaxBuku]Buku

const TABEL = 6

const NmaxMember int = 10000

type Member struct {
	Nama string
}

type TabMember [NmaxMember]Member

func InputBuku(n *Buku) {
	fmt.Print("Masukan Judul Buku : ")
	fmt.Scan(&n.Judul)
	fmt.Print("Masukan Nama Pengarang : ")
	fmt.Scan(&n.Pengarang)
	fmt.Print("Masukan Kode Buku : ")
	fmt.Scan(&n.Kode)
	fmt.Print("Masukan Tahun Terbit Buku : ")
	fmt.Scan(&n.T_Terbit)
	fmt.Print("Masukan Jumlah Buku : ")
	fmt.Scan(&n.JmlBuku)
}

func TambahBuku(T *TabBuku, n Buku, i *int) {
	T[*i] = n
	*i++
}

func CetakBuku(T TabBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i].Judul)
		fmt.Println(T[i].Pengarang)
		fmt.Println(T[i].Kode)
		fmt.Println(T[i].T_Terbit)
		fmt.Println(T[i].JmlBuku)
	}
}

func CariBuku_Nama(T TabBuku, n int, Judul string) int {

	for i := 0; i < n; i++ {
		if T[i].Judul == Judul {
			return 1
		}
	}
	return -1
}

func CariBuku_Kode(T TabBuku, n int, Kode string) bool {

	for i := 0; i < n; i++ {
		if T[i].Kode == Kode {
			return true
		}
	}
	return false
}

func HapusBuku(T *TabBuku, n *int, Judul string) {
	idx := CariBuku_Nama(*T, *n, Judul)
	for i := idx; i < *n; i++ {
		T[i] = T[i+1]
	}
	T[*n-1] = Buku{}
	*n--
}

func DaftarMember(A *Regristrasi.Pengguna, N int) {
	var i int
	for i = 1; i < N; i++ {
		fmt.Println(A[i].Nama)
	}

}

func PinjamBuku(n *Member, T *Buku) {
	fmt.Print("Nama : ")
	fmt.Scan(&n.Nama)
	fmt.Print("Judul Buku : ")
	fmt.Scan(&T.Judul)
	fmt.Print("Tanggal Pinjam")
}

func TambahPeminjam(n *TabMember, i *int, r Member) {
	n[*i] = r
	*i++
}
func PengembalianBuku(n *Member, T *Buku) {
	fmt.Print("Nama : ")
	fmt.Scan(&n.Nama)
	fmt.Print("Judul Buku : ")
	fmt.Scan(&T.Judul)
	fmt.Print("Tanggal Kembali : ")
}
