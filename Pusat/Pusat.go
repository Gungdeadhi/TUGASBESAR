package Pusat

import (
	"fmt"
	"strconv"
	"strings"
)

const NmaxBuku = 10000

type Buku struct {
	Judul, Pengarang, Kode string
	T_Terbit, JmlBuku      int
}

type TabBuku [NmaxBuku]Buku

const TabCetak = 6

var Tabel = [TabCetak]string{
	"NO",
	"JUDUL",
	"KODE",
	"PENGARANG",
	"TAHUN TERBIT",
	"STOK",
}

const TabCari = 5

var TabelCari = [TabCari]string{
	"JUDUL",
	"KODE",
	"PENGARANG",
	"TAHUN TERBIT",
	"STOK",
}

func InputBuku(n *Buku) {
	fmt.Print("Masukan Judul Buku		: ")
	fmt.Scan(&n.Judul)
	fmt.Print("Masukan Nama Pengarang 		: ")
	fmt.Scan(&n.Pengarang)
	fmt.Print("Masukan Kode Buku 		: ")
	fmt.Scan(&n.Kode)
	fmt.Print("Masukan Tahun Terbit Buku	: ")
	fmt.Scan(&n.T_Terbit)
	fmt.Print("Masukan Jumlah Buku 		: ")
	fmt.Scan(&n.JmlBuku)
}

func TambahBuku(T *TabBuku, N *int, k Buku) {
	T[*N] = k
	*N++
}

func ToStr(n int) string {
	return strconv.Itoa(n)
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func LJust(s string, lebar int) string {
	var padding string

	if lebar <= len(s) {
		return s
	}

	padding = strings.Repeat(" ", lebar-len(s))
	return s + padding
}

func CetakBuku(T TabBuku, n int) {
	var i int
	var MaxTab [TabCetak]int
	var Pemisah string

	for i = 0; i < TabCetak; i++ {
		MaxTab[i] = len(Tabel[i])
	}

	for i = 0; i < n; i++ {
		MaxTab[0] = Max(MaxTab[0], len(ToStr(i+1)))
		MaxTab[1] = Max(MaxTab[1], len(T[i].Judul))
		MaxTab[2] = Max(MaxTab[2], len(T[i].Pengarang))
		MaxTab[3] = Max(MaxTab[3], len(T[i].Kode))
		MaxTab[4] = Max(MaxTab[4], len(ToStr(T[i].T_Terbit)))
		MaxTab[5] = Max(MaxTab[5], len(ToStr(T[i].JmlBuku)))
	}

	Pemisah = "+"

	for i = 0; i < TabCetak; i++ {
		Pemisah += strings.Repeat("-", MaxTab[i]+2)
		Pemisah += "+"
	}

	fmt.Println(Pemisah)

	for i = 0; i < TabCetak; i++ {
		fmt.Printf("| %s ", LJust(Tabel[i], MaxTab[i]))
	}

	fmt.Println("|")
	fmt.Println(Pemisah)

	for i = 0; i < n; i++ {
		fmt.Printf(
			"| %s | %s | %s | %s | %s | %s |\n",
			LJust(ToStr(i+1), MaxTab[0]),
			LJust(T[i].Judul, MaxTab[1]),
			LJust(T[i].Kode, MaxTab[2]),
			LJust(T[i].Pengarang, MaxTab[3]),
			LJust(ToStr(T[i].T_Terbit), MaxTab[4]),
			LJust(ToStr(T[i].JmlBuku), MaxTab[5]),
		)
	}
	fmt.Println(Pemisah)
}

func CariBuku_Nama(T TabBuku, n int, judul string) int {
	for i := 0; i < n; i++ {
		if T[i].Judul == judul {
			return i
		}
	}
	return -1
}

func CariBuku_Kode(T TabBuku, n int, kode string) int {
	for i := 0; i < n; i++ {
		if T[i].Kode == kode {
			return i
		}
	}
	return -1
}

func EditBuku(T *TabBuku, n int, judul string, k Buku) {
	idx := CariBuku_Nama(*T, n, judul)
	T[idx] = k
}

func HapusBuku(T *TabBuku, n int, Judul string) {
	idx := CariBuku_Nama(*T, n, Judul)
	for i := idx; i < n; i++ {
		T[i] = T[i+1]
	}
	T[n-1] = Buku{}
	n--
}

func SelSort(T *TabBuku, n int) {
	var pass, i, idx int
	var temp Buku
	pass = 1
	for pass < n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if T[i].JmlBuku > T[idx].JmlBuku {
				idx = i
			}
			i++
		}
		temp = T[idx]
		T[idx] = T[pass-1]
		T[pass-1] = temp
		pass++
	}
}

func CetakCariBuku(T TabBuku, n int) {
	var i int
	var MaxTab [TabCari]int
	var Pemisah string

	for i = 0; i < TabCari; i++ {
		MaxTab[i] = len(Tabel[i])
	}

	for i = 0; i < n; i++ {
		MaxTab[0] = Max(MaxTab[0], len(T[i].Judul))
		MaxTab[1] = Max(MaxTab[1], len(T[i].Pengarang))
		MaxTab[2] = Max(MaxTab[2], len(T[i].Kode))
		MaxTab[3] = Max(MaxTab[3], len(ToStr(T[i].T_Terbit)))
		MaxTab[4] = Max(MaxTab[4], len(ToStr(T[i].JmlBuku)))
	}

	Pemisah = "+"

	for i = 0; i < TabCari; i++ {
		Pemisah += strings.Repeat("-", MaxTab[i]+2)
		Pemisah += "+"
	}

	fmt.Println(Pemisah)

	for i = 0; i < TabCari; i++ {
		fmt.Printf("| %s ", LJust(Tabel[i], MaxTab[i]))
	}

	fmt.Println("|")
	fmt.Println(Pemisah)

	for i = 0; i < n; i++ {
		fmt.Printf(
			"| %s | %s | %s | %s | %s |\n",
			LJust(T[i].Judul, MaxTab[0]),
			LJust(T[i].Kode, MaxTab[1]),
			LJust(T[i].Pengarang, MaxTab[2]),
			LJust(ToStr(T[i].T_Terbit), MaxTab[3]),
			LJust(ToStr(T[i].JmlBuku), MaxTab[4]),
		)
	}
	fmt.Println(Pemisah)
}
