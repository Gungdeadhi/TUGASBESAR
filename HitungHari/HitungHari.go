package HitungHari

import "fmt"

type Date struct {
	Tgl, Bln, Thn int
}

func FormatTanggal(d Date) string {
	return fmt.Sprintf("%02d-%02d-%d", d.Tgl, d.Bln, d.Thn)
}

func f(d Date) int {
	if d.Bln <= 2 {
		return d.Thn - 1
	} else {
		return d.Thn
	}
}

func g(d Date) int {
	if d.Bln <= 2 {
		return d.Bln + 13
	} else {
		return d.Bln + 1
	}
}

func N(d Date) int {
	return ((1461 * f(d)) / 4) + (153*g(d))/5 + d.Tgl
}
