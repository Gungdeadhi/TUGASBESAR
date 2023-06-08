package Pusat

import "fmt"

const NMAX = 10000

type Buku struct {
	Judul, Pengarang, Kode string
	T_Terbit, JmlBuku      int
}

type TabBuku [NMAX]Buku

const TABEL = 6

var H = [TABEL]string{
	"NO",
	"JUDUL",
	"KODE",
	"PENGARANG",
	"TAHUN",
	"STOK",
}

func TambahBuku(T *TabBuku, n Buku, i *int) {
	T[*i] = n
	*i++
}

func CariBuku(T TabBuku, n int, Judul string) int {

	for i := 0; i < n; i++ {
		if T[i].Judul == Judul {
			return i
		}
	}
	return -1
}

func HapusBuku(T *TabBuku, n *int, Judul string) {
	idx := CariBuku(*T, *n, Judul)
	for i := idx; i < *n; i++ {
		T[i] = T[i+1]
	}
	T[*n-1] = Buku{}
	*n--
}

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
