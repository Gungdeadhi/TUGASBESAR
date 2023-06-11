package Peminjaman

import (
	"fmt"
)

const NmaxPinjam int = 1000

type Pinjam struct {
	Nama        string
	Judul       string
	Tgl_Pinjam  int
	Tgl_kembali int
	Bulan       int
	Tahun       int
}
type TabPinjam [NmaxPinjam]Pinjam

func InputPinjamBuku(n *Pinjam) {
	fmt.Print("Nama : ")
	fmt.Scan(&n.Nama)
	fmt.Print("Judul Buku : ")
	fmt.Scan(&n.Judul)
	fmt.Print("Tanggal Pinjam : ")
	fmt.Scan(&n.Tgl_Pinjam)
	fmt.Print("Bulan Pinjam : ")
	fmt.Scan(&n.Bulan)
	fmt.Print("Tahun Pinjam : ")
	fmt.Scan(&n.Tahun)
	fmt.Print("Tanggal Kembali : ")
	fmt.Scan(&n.Tgl_kembali)
	fmt.Print("Bulan Kembali : ")
	fmt.Scan(&n.Bulan)
	fmt.Print("Tahun Kembali : ")
	fmt.Scan(&n.Tahun)
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
	fmt.Print("Tanggal Kembali : ")
	fmt.Print("Tanggal Pinjam : ")
	fmt.Scan(&n.Tgl_Pinjam)
	fmt.Print("Bulan Pinjam : ")
	fmt.Scan(&n.Bulan)
	fmt.Print("Tahun Pinjam : ")
	fmt.Scan(&n.Tahun)
	fmt.Print("Tanggal Kembali : ")
	fmt.Scan(&n.Tgl_kembali)
	fmt.Print("Bulan Kembali : ")
	fmt.Scan(&n.Bulan)
	fmt.Print("Tahun Kembali : ")
	fmt.Scan(&n.Tahun)
}

func Kabisat(T *Pinjam) bool {
	if T.Tahun%4 == 0 && (T.Tahun%400 == 0 || T.Tahun%100 != 0) {
		return true
	}
	return false
}
func getJumlahHari(T Pinjam, jmlHari *int) {
	if T.Bulan == 2 {
		if Kabisat(&T) {
			*jmlHari = 29
		} else {
			*jmlHari = 28
		}
	} else {
		if T.Bulan <= 7 {
			*jmlHari = 30 + (T.Bulan % 2)
		} else {
			*jmlHari = 31 - (T.Bulan % 2)
		}
	}
}
func BiayaPeminjaman()
