package Peminjaman

import (
	"fmt"
)

const NmaxPinjam int = 1000

type Pinjam struct {
	Judul        string
	Nama         string
	Tgl_Pinjam   Date
	Tgl_kembali1 Date
	Tgl_kembali2 Date
}
type Date struct {
	Tgl, Bln, Thn int
}
type TabPinjam [NmaxPinjam]Pinjam

func InputPinjamBuku(n *Pinjam) {

	fmt.Print("Nama : ")
	fmt.Scan(&n.Nama)
	fmt.Print("Judul Buku : ")
	fmt.Scan(&n.Judul)
	fmt.Print("Tanggal Pinjam : ")
	fmt.Scan(&n.Tgl_Pinjam.Tgl)
	fmt.Print("Bulan Pinjam : ")
	fmt.Scan(&n.Tgl_Pinjam.Bln)
	fmt.Print("Tahun Pinjam : ")
	fmt.Scan(&n.Tgl_Pinjam.Thn)
	fmt.Print("Tanggal Kembali : ")
	fmt.Scan(&n.Tgl_kembali1.Tgl)
	fmt.Print("Bulan Kembali : ")
	fmt.Scan(&n.Tgl_kembali1.Bln)
	fmt.Print("Tahun Kembali : ")
	fmt.Scan(&n.Tgl_kembali1.Thn)
}

func TambahPeminjam(n *TabPinjam, i *int, r Pinjam) {
	n[*i] = r
	*i++
}

func PengembalianBuku(n *Pinjam) {

	fmt.Print("Nama : ")
	fmt.Scan(&n.Nama)
	fmt.Print("Judul Buku : ")
	fmt.Scan(&n.Judul)
	fmt.Print("Tanggal Pinjam : ")
	fmt.Scan(&n.Tgl_Pinjam.Tgl)
	fmt.Print("Bulan Pinjam : ")
	fmt.Scan(&n.Tgl_Pinjam.Bln)
	fmt.Print("Tahun Pinjam : ")
	fmt.Scan(&n.Tgl_Pinjam.Thn)
	fmt.Print("Tanggal Kembali : ")
	fmt.Scan(&n.Tgl_kembali2.Tgl)
	fmt.Print("Bulan Kembali : ")
	fmt.Scan(&n.Tgl_kembali2.Bln)
	fmt.Print("Tahun Kembali : ")
	fmt.Scan(&n.Tgl_kembali2.Thn)
}

func Kabisat(T Date) bool {
	return T.Thn%4 == 0 && (T.Thn%400 == 0 || T.Thn%100 != 0)
}

func HitungSelisihHari()

// func GetJumlahHari(T Pinjam, jmlHari *int) {
// 	if T.Bulan == 2 {
// 		if Kabisat(T) {
// 			*jmlHari = 29
// 		} else {
// 			*jmlHari = 28
// 		}
// 	} else {
// 		if T.Bulan <= 7 {
// 			*jmlHari = 30 + (T.Bulan % 2)
// 		} else {
// 			*jmlHari = 31 - (T.Bulan % 2)
// 		}
// 	}
// }
// func BiayaPeminjaman(T Pinjam, jmlHari *int){
// 	var Biaya, Toatal int

// 	fmt.Scan(&Biaya)
// 	Total = GetJumlahHari(T, jmlHari) * Biaya
// }
