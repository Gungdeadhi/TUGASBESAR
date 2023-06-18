package Member

import (
	"fmt"
	"strconv"
	"strings"
)

const NmaxMember = 10000

type Member struct {
	Asal, Nama string
	Umur       int
}

type TabMember [NmaxMember]Member

const TabCetak = 4

var Tabel = [TabCetak]string{
	"NO",
	"NAMA",
	"UMUR",
	"ALAMAT MEMBER",
}

func InputMember(r *Member) {
	fmt.Print("Nama Member Baru : ")
	fmt.Scan(&r.Nama)
	fmt.Print("Umur Member : ")
	fmt.Scan(&r.Umur)
	fmt.Print("Alamat Member : ")
	fmt.Scan(&r.Asal)
}

func TambahMember(T *TabMember, i *int, r Member) {
	T[*i] = r
	*i++
}

func CariMember(T TabMember, n int, Nama string) int {
	for i := 0; i < n; i++ {
		if T[i].Nama == Nama {
			return i
		}
	}
	return -1
}

func HapusMember(T *TabMember, n *int, Nama string) {
	idx := CariMember(*T, *n, Nama)
	for i := idx; i < *n; i++ {
		T[i] = T[i+1]
	}
	T[*n-1] = Member{}
	*n--
}

func EditMember(T *TabMember, n int, Nama string, r Member) {
	idx := CariMember(*T, n, Nama)
	T[idx] = r
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

func CetakMember(T TabMember, n int) {
	var i int
	var MaxTab [TabCetak]int
	var Pemisah string

	for i = 0; i < TabCetak; i++ {
		MaxTab[i] = len(Tabel[i])
	}

	for i = 0; i < n; i++ {
		MaxTab[0] = Max(MaxTab[0], len(ToStr(i+1)))
		MaxTab[1] = Max(MaxTab[1], len(T[i].Nama))
		MaxTab[2] = Max(MaxTab[2], len(ToStr(T[i].Umur)))
		MaxTab[3] = Max(MaxTab[3], len(T[i].Asal))
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
			"| %s | %s | %s | %s |\n",
			LJust(ToStr(i+1), MaxTab[0]),
			LJust(T[i].Nama, MaxTab[1]),
			LJust(ToStr(T[i].Umur), MaxTab[2]),
			LJust(T[i].Asal, MaxTab[3]),
		)
	}
	fmt.Println(Pemisah)
}
