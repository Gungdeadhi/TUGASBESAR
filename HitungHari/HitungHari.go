package HitungHari

import "fmt"

type Date struct {
	Tgl, Bln, Thn int
}

func FormatTanggal(d Date) string {
	return fmt.Sprintf("%02d-%02d-%d", d.Tgl, d.Bln, d.Thn)
}

func Kabisat(Tahun int) bool {
	return Tahun%4 == 0 && (Tahun%400 == 0 || Tahun%100 != 0)
}

func TotalHari(Bulan, Tahun int) int {
	if Bulan == 2 {
		if Kabisat(Tahun) {
			return 29
		} else {
			return 28
		}
	} else {
		if Bulan <= 7 {
			return 30 + (Bulan % 2)
		} else {
			return 31 - (Bulan % 2)
		}
	}
}

func CekValid(d Date) bool {
	if d.Tgl < 1 {
		return false
	}

	if d.Bln < 1 || d.Bln > 12 {
		return false
	}

	if d.Thn < 0 {
		return false
	}

	if d.Tgl > TotalHari(d.Bln, d.Thn) {
		return false
	}

	return true
}

func LebihDulu(d1, d2 Date) bool {
	if d1.Thn < d2.Thn {
		return true
	} else if d1.Thn > d2.Thn {
		return false
	}

	if d1.Bln < d2.Bln {
		return true
	} else if d1.Bln > d2.Bln {
		return false
	}

	if d1.Tgl < d2.Tgl {
		return true
	}

	return false
}
func SisaHari(d1, d2 Date) int {
	var SisaHari int

	if !LebihDulu(d1, d2) {
		d1, d2 = d2, d1
	}

	for !(d1.Thn == d2.Thn && d1.Bln == d2.Bln) {
		SisaHari += TotalHari(d1.Bln, d1.Thn) - d1.Tgl + 1

		if d1.Bln < 12 {
			d1.Tgl = 1
			d1.Bln = d1.Bln + 1
		} else {
			d1.Tgl = 1
			d1.Bln = 1
			d1.Thn = d1.Thn + 1
		}
	}

	SisaHari += d2.Tgl - d1.Tgl

	return SisaHari
}
