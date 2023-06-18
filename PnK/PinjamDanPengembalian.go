package PnK

import (
	"TUGASBESAR/HitungHari"
	"fmt"
	"strings"
)

const NmaxPinjam int = 1000

type Pinjam struct {
	Nama         string
	Judul        string
	Tgl_Pinjam1  HitungHari.Date
	Tgl_Kembali1 HitungHari.Date
	Biaya1       int
}

type TabPinjam [NmaxPinjam]Pinjam

const NmaxPengembali = 1000

type Pengembali struct {
	Nama         string
	Judul        string
	Tgl_Kembali1 HitungHari.Date
	Tgl_Kembali2 HitungHari.Date
	Biaya2       int
}
type TabPengembali [NmaxPengembali]Pengembali

const NmaxCetakPeminjam = 6

var TabelPeminjam [NmaxCetakPeminjam]string = [NmaxCetakPeminjam]string{
	"NO",
	"NAMA",
	"JUDUL",
	"TANGGAL PINJAM",
	"TANGGAL KEMBALI",
	"BIAYA",
}

const NmaxCetakPengembalian = 7

var TabelPengembalian [NmaxCetakPengembalian]string = [NmaxCetakPengembalian]string{
	"NO",
	"NAMA",
	"JUDUL",
	"TANGGAL KEMBALI SEBENARNYA",
	"TANGGAL KEMBALI HARI INI",
	"DENDA",
	"TOTAL",
}

func InputPinjamBuku(n *Pinjam) {
	fmt.Print("Nama 				: ")
	fmt.Scan(&n.Nama)
	fmt.Print("Judul Buku 			: ")
	fmt.Scan(&n.Judul)
	fmt.Print("Tanggal Pinjam 			: ")
	fmt.Scan(&n.Tgl_Pinjam1.Tgl, &n.Tgl_Pinjam1.Bln, &n.Tgl_Pinjam1.Thn)
	fmt.Print("Tanggal Kembali 		: ")
	fmt.Scan(&n.Tgl_Kembali1.Tgl, &n.Tgl_Kembali1.Bln, &n.Tgl_Kembali1.Thn)
	fmt.Print("Masukan Biaya Pinjam 		: ")
	fmt.Scan(&n.Biaya1)
}

func TambahPeminjam(n *TabPinjam, i *int, r Pinjam) {
	n[*i] = r
	*i++
}

func ToStr(n int) string {
	return fmt.Sprintf("%d", n)
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
	var i, fd3 int
	var MaxTab [NmaxCetakPeminjam]int
	var Pemisah, fd1, fd2 string

	for i = 0; i < NmaxCetakPeminjam; i++ {
		MaxTab[i] = len(TabelPeminjam[i])
	}

	for i = 0; i < n; i++ {

		fd1 = HitungHari.FormatTanggal(T[i].Tgl_Pinjam1)
		fd2 = HitungHari.FormatTanggal(T[i].Tgl_Kembali1)
		fd3 = BiayaPinjam(T, i)

		MaxTab[0] = Max(MaxTab[0], len(ToStr(i+1)))
		MaxTab[1] = Max(MaxTab[1], len(T[i].Nama))
		MaxTab[2] = Max(MaxTab[2], len(T[i].Judul))
		MaxTab[3] = Max(MaxTab[3], len(fd1))
		MaxTab[4] = Max(MaxTab[3], len(fd2))
		MaxTab[5] = Max(MaxTab[5], len(ToStr(fd3)))
	}

	Pemisah = "+"

	for i = 0; i < NmaxCetakPeminjam; i++ {
		Pemisah += strings.Repeat("-", MaxTab[i]+2)
		Pemisah += "+"
	}

	fmt.Println(Pemisah)

	for i = 0; i < NmaxCetakPeminjam; i++ {
		fmt.Printf("| %s ", LJust(TabelPeminjam[i], MaxTab[i]))
	}

	fmt.Println("|")
	fmt.Println(Pemisah)

	for i = 0; i < n; i++ {

		fd1 = HitungHari.FormatTanggal(T[i].Tgl_Pinjam1)
		fd2 = HitungHari.FormatTanggal(T[i].Tgl_Kembali1)
		fd3 = BiayaPinjam(T, i)

		fmt.Printf(
			"| %s | %s | %s | %s | %s | %s |\n",
			LJust(ToStr(i+1), MaxTab[0]),
			LJust(T[i].Nama, MaxTab[1]),
			LJust(T[i].Judul, MaxTab[2]),
			LJust(fd1, MaxTab[3]),
			LJust(fd2, MaxTab[4]),
			LJust(ToStr(fd3), MaxTab[5]),
		)
	}
	fmt.Println(Pemisah)
}

func CariPeminjamBuku(T TabPinjam, n int, Nama string) int {
	for i := 0; i < n; i++ {
		if T[i].Nama == Nama {
			return i
		}
	}
	return -1
}

func HapusPeminjam(T *TabPinjam, n int, Nama string) {
	idx := CariPeminjamBuku(*T, n, Nama)
	for i := idx; i < n; i++ {
		T[i] = T[i+1]
	}
	T[n-1] = Pinjam{}
	n--
}

func InputPengembalianBuku(n *Pengembali) {
	fmt.Print("Nama 			: ")
	fmt.Scan(&n.Nama)
	fmt.Print("Judul 			: ")
	fmt.Scan(&n.Judul)
	fmt.Print("Tanggal Kembali Seharusnya 	: ")
	fmt.Scan(&n.Tgl_Kembali1.Tgl, &n.Tgl_Kembali1.Bln, &n.Tgl_Kembali1.Thn)
	fmt.Print("Tanggal Kembali Buku Hari Ini : ")
	fmt.Scan(&n.Tgl_Kembali2.Tgl, &n.Tgl_Kembali2.Bln, &n.Tgl_Kembali2.Thn)
}

func TambahPengembalian(n *TabPengembali, i *int, r Pengembali) {
	n[*i] = r
	*i++
}

func CetakPengembalian(T1 TabPinjam, T2 TabPengembali, n int) {
	var i, fd3, fd4 int
	var MaxTab [NmaxCetakPengembalian]int
	var Pemisah, fd1, fd2 string

	for i = 0; i < NmaxCetakPengembalian; i++ {
		MaxTab[i] = len(TabelPengembalian[i])
	}

	for i = 0; i < n; i++ {
		fd1 = HitungHari.FormatTanggal(T2[i].Tgl_Kembali1)
		fd2 = HitungHari.FormatTanggal(T2[i].Tgl_Kembali2)
		fd3 = Denda(T2, i)
		fd4 = Total(T1, T2, i)

		MaxTab[0] = Max(MaxTab[0], len(ToStr(i+1)))
		MaxTab[1] = Max(MaxTab[1], len(T2[i].Nama))
		MaxTab[2] = Max(MaxTab[2], len(T2[i].Judul))
		MaxTab[3] = Max(MaxTab[3], len(fd1))
		MaxTab[4] = Max(MaxTab[4], len(fd2))
		MaxTab[5] = Max(MaxTab[5], len(ToStr(fd3)))
		MaxTab[6] = Max(MaxTab[6], len(ToStr(fd4)))
	}

	Pemisah = "+"

	for i = 0; i < NmaxCetakPengembalian; i++ {
		Pemisah += strings.Repeat("-", MaxTab[i]+2)
		Pemisah += "+"
	}

	fmt.Println(Pemisah)

	for i = 0; i < NmaxCetakPengembalian; i++ {
		fmt.Printf("| %s ", LJust(TabelPengembalian[i], MaxTab[i]))
	}

	fmt.Println("|")
	fmt.Println(Pemisah)

	for i = 0; i < n; i++ {

		fd1 = HitungHari.FormatTanggal(T2[i].Tgl_Kembali1)
		fd2 = HitungHari.FormatTanggal(T2[i].Tgl_Kembali2)
		fd3 = Denda(T2, i)
		fd4 = Total(T1, T2, i)

		fmt.Printf(
			"| %s | %s | %s | %s | %s | %s | %s |\n",
			LJust(ToStr(i+1), MaxTab[0]),
			LJust(T2[i].Nama, MaxTab[1]),
			LJust(T2[i].Judul, MaxTab[2]),
			LJust(fd1, MaxTab[3]),
			LJust(fd2, MaxTab[4]),
			LJust(ToStr(fd3), MaxTab[5]),
			LJust(ToStr(fd4), MaxTab[6]),
		)
	}
	fmt.Println(Pemisah)
}

func CariPengembaliBuku(T TabPengembali, n int, Nama string) int {
	for i := 0; i < n; i++ {
		if T[i].Nama == Nama {
			return i
		}
	}
	return -1
}
func HapusPengembali(T *TabPengembali, n *int, Nama string) {
	idx := CariPengembaliBuku(*T, *n, Nama)
	for i := idx; i < *n; i++ {
		T[i] = T[i+1]
	}
	T[*n-1] = Pengembali{}
	*n--
}

func BiayaPinjam(T TabPinjam, idx int) int {
	return T[idx].Biaya1 * (HitungHari.N(T[idx].Tgl_Kembali1) - HitungHari.N(T[idx].Tgl_Pinjam1))
}

func CekKeterlambat(T TabPengembali, idx int) bool {
	return HitungHari.Mendahului(T[idx].Tgl_Kembali1, T[idx].Tgl_Kembali2)
}

func Denda(T TabPengembali, idx int) int {
	if CekKeterlambat(T, idx) {
		return 10000 * (HitungHari.N(T[idx].Tgl_Kembali2) - HitungHari.N(T[idx].Tgl_Kembali1))
	} else {
		return 0
	}

}

func Total(T1 TabPinjam, T2 TabPengembali, idx1 int) int {
	return BiayaPinjam(T1, idx1) + Denda(T2, idx1)
}
