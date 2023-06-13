package PnK

import (
	"TUGASBESAR/HitungHari"
	"fmt"
	"strconv"
	"strings"
)

const NmaxPinjam int = 1000

type Pinjam struct {
	Nama         string
	Judul        string
	Tgl_Pinjam1  HitungHari.Date
	Tgl_Kembali1 HitungHari.Date
}

type TabPinjam [NmaxPinjam]Pinjam

const NmaxPengembali = 100

type Pengembali struct {
	Nama         string
	Judul        string
	Tgl_Pinjam2  HitungHari.Date
	Tgl_Kembali2 HitungHari.Date
}
type TabPengembali [NmaxPengembali]Pengembali

const TabCetakPeminjam = 6

var TabelPeminjam [TabCetakPeminjam]string = [TabCetakPeminjam]string{
	"NO",
	"NAMA",
	"JUDUL",
	"TANGGAL PINJAM",
	"TANGGAL KEMBALI",
	"LAMA",
}

const TabCetakPengembalian = 6

var TabelPengembalian [TabCetakPengembalian]string = [TabCetakPengembalian]string{
	"NO",
	"NAMA",
	"JUDUL",
	"TANGGAL PINJAM",
	"TANGGAL KEMBALI",
	"DENDA",
}

func InputPinjamBuku(n *Pinjam) {
	fmt.Print("Nama : ")
	fmt.Scan(&n.Nama)
	fmt.Print("Judul Buku : ")
	fmt.Scan(&n.Judul)
	fmt.Print("Tanggal Pinjam : ")
	fmt.Scan(&n.Tgl_Pinjam1.Tgl, &n.Tgl_Pinjam1.Bln, &n.Tgl_Pinjam1.Thn)
	fmt.Print("Tanggal Kembali : ")
	fmt.Scan(&n.Tgl_Kembali1.Tgl, &n.Tgl_Kembali1.Bln, &n.Tgl_Kembali1.Thn)
}

func TambahPeminjam(n *TabPinjam, i *int, r Pinjam) {
	n[*i] = r
	*i++
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

func CetakPeminjam(T TabPinjam, n int) {
	var i int
	var MaxTab [TabCetakPeminjam]int
	var Pemisah, fd1, fd2 string

	for i = 0; i < TabCetakPeminjam; i++ {
		MaxTab[i] = len(TabelPeminjam[i])
	}

	for i = 0; i < n; i++ {

		fd1 = HitungHari.FormatTanggal(T[i].Tgl_Pinjam1)
		fd2 = HitungHari.FormatTanggal(T[i].Tgl_Kembali1)

		MaxTab[0] = Max(MaxTab[0], len(ToStr(i+1)))
		MaxTab[1] = Max(MaxTab[1], len(T[i].Nama))
		MaxTab[2] = Max(MaxTab[2], len(T[i].Judul))
		MaxTab[3] = Max(MaxTab[3], len(fd1))
		MaxTab[4] = Max(MaxTab[3], len(fd2))
	}

	Pemisah = "+"

	for i = 0; i < TabCetakPeminjam; i++ {
		Pemisah += strings.Repeat("-", MaxTab[i]+2)
		Pemisah += "+"
	}

	fmt.Println(Pemisah)

	for i = 0; i < TabCetakPeminjam; i++ {
		fmt.Printf("| %s", LJust(TabelPeminjam[i], MaxTab[i]))
	}

	fmt.Println("|")
	fmt.Println(Pemisah)

	for i = 0; i < n; i++ {

		fd1 = HitungHari.FormatTanggal(T[i].Tgl_Pinjam1)
		fd2 = HitungHari.FormatTanggal(T[i].Tgl_Kembali1)

		fmt.Printf(
			"| %s | %s | %s | %s | %s |\n",
			LJust(ToStr(i+1), MaxTab[0]),
			LJust(T[i].Nama, MaxTab[1]),
			LJust(T[i].Judul, MaxTab[2]),
			LJust(fd1, MaxTab[3]),
			LJust(fd2, MaxTab[4]),
		)
	}
	fmt.Println(Pemisah)
}

func CariPeminjamBuku(T TabPinjam, n int, Nama string) int {
	for i := 0; i < n; i++ {
		if T[i].Nama == Nama {
			return 1
		}
	}
	return -1
}

func HapusPeminjam(T *TabPinjam, n *int, Nama string) {
	idx := CariPeminjamBuku(*T, *n, Nama)
	for i := idx; i < *n; i++ {
		T[i] = T[i+1]
	}
	T[*n-1] = Pinjam{}
	*n--
}

func InputPengembalianBuku(n *Pengembali) {
	fmt.Print("Nama : ")
	fmt.Scan(&n.Nama)
	fmt.Print("Judul : ")
	fmt.Scan(&n.Judul)
	fmt.Print("Tanggal Pinjam : ")
	fmt.Scan(&n.Tgl_Pinjam2.Tgl, &n.Tgl_Pinjam2.Bln, &n.Tgl_Pinjam2.Thn)
	fmt.Print("Tanggal Kembali : ")
	fmt.Scan(&n.Tgl_Kembali2.Tgl, &n.Tgl_Pinjam2.Bln, &n.Tgl_Pinjam2.Thn)
}
func TambahPengembalian(n *TabPinjam, i *int, r Pinjam) {
	n[*i] = r
	*i++
}

// func CetakPengembalian(T TabPengembali, n int) {
// 	var i int
// 	var MaxTab [TabCetakPengembalian]int
// 	var Pemisah, fd1, fd2 string

// 	for i = 0; i < TabCetakPengembalian; i++ {
// 		MaxTab[i] = len(TabelPengembalian[i])
// 	}

// 	for i = 0; i < n; i++ {

// 		fd1 = HitungHari.FormatTanggal(T[i].Tgl_Pinjam2)
// 		fd2 = HitungHari.FormatTanggal(T[i].Tgl_Kembali2)

// 		MaxTab[0] = Max(MaxTab[0], len(ToStr(i+1)))
// 		MaxTab[1] = Max(MaxTab[1], len(T[i].Nama))
// 		MaxTab[2] = Max(MaxTab[2], len(T[i].Judul))
// 		MaxTab[3] = Max(MaxTab[3], len(fd1))
// 		MaxTab[4] = Max(MaxTab[3], len(fd2))
// 		// MaxTab[5] = Max(MaxTab[5])
// 	}

// 	Pemisah = "+"

// 	for i = 0; i < TabCetakPeminjam; i++ {
// 		Pemisah += strings.Repeat("-", MaxTab[i]+2)
// 		Pemisah += "+"
// 	}

// 	fmt.Println(Pemisah)

// 	for i = 0; i < TabCetakPengembalian; i++ {
// 		fmt.Printf("| %s", LJust(TabelPengembalian[i], MaxTab[i]))
// 	}

// 	fmt.Println("|")
// 	fmt.Println(Pemisah)

// 	for i = 0; i < n; i++ {

// 		fd1 = HitungHari.FormatTanggal(T[i].Tgl_Pinjam2)
// 		fd2 = HitungHari.FormatTanggal(T[i].Tgl_Kembali2)

// 		fmt.Printf(
// 			"| %s | %s | %s | %s | %s | %s |\n",
// 			LJust(ToStr(i+1), MaxTab[0]),
// 			LJust(T[i].Nama, MaxTab[1]),
// 			LJust(T[i].Judul, MaxTab[2]),
// 			LJust(fd1, MaxTab[3]),
// 			LJust(fd2, MaxTab[4]),
// 		)
// 	}
// 	fmt.Println(Pemisah)
// }

// func CariPengembaliBuku(T TabPinjam, n int, Nama string) int {
// 	for i := 0; i < n; i++ {
// 		if T[i].Nama == Nama {
// 			return 1
// 		}
// 	}
// 	return -1
// }
// func HapusPengembali(T *TabPinjam, n *int, Nama string) {
// 	idx := CariPengembaliBuku(*T, *n, Nama)
// 	for i := idx; i < *n; i++ {
// 		T[i] = T[i+1]
// 	}
// 	T[*n-1] = Pinjam{}
// 	*n--
// }

// func BiayaPinjam(T Pinjam, Biaya *int) {
// 	var d1, d2 HitungHari.Date

// 	HitungHari.SisaHari(d1, d2)
// 	*Biaya = 10000 * HitungHari.SisaHari(d1, d2)

// }

// func CekKeterlambat(T1 Pinjam, T2 Pengembali) bool {

// 	if T1.Tgl_Kembali1.Tgl > T2.Tgl_Kembali2.Tgl {
// 		return true
// 	} else {
// 		return false
// 	}

// }

// func Denda(T Pinjam, Denda *int) {
// 	var d1, d2 HitungHari.Date
// 	var Hari, Biaya int

// 	BiayaPinjam(T, Hari, &Biaya)
// 	if CekKeterlambat(T) {
// 		*Denda = ((T.Tanggal3 - T.Tanggal2) * 5000) + HitungHari.SisaHari(d1, d2)
// 	}

// }
