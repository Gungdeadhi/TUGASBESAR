package Admin

import "fmt"

const NMAX = 10000

type Buku struct {
	Judul, Pengarang, Kode string
	T_Terbit, JmlBuku      int
}

type TabBuku [NMAX]Buku

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
	fmt.Println("Masukan Judul Buku : ")
	fmt.Scan(&n.Judul)
	fmt.Println("Masukan Nama Pengarang : ")
	fmt.Scan(&n.Pengarang)
	fmt.Println("Masukan Kode Buku : ")
	fmt.Scan(&n.Kode)
	fmt.Println("Masukan Tahun Terbit Buku : ")
	fmt.Scan(&n.T_Terbit)
	fmt.Println("Masukan Jumlah Buku : ")
	fmt.Scan(&n.JmlBuku)
}
