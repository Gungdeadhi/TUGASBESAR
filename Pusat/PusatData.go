package Pusat

import "fmt"

const NMAX = 10000

// func bubble_sort_1(A *[5]int, n int){
// 	var i, k int
// 	for k = 1; k <= n-1; k++{
// 		for i = 0;
// 	}
//}
// func selection_sort(A *[10]int, n int){
// 	var i pas
// }

func Login(usser, password *string) bool {
	var i int = 3

	fmt.Print("Masukan Usser : ")
	fmt.Scan(usser)
	fmt.Print("Masukan Password : ")
	fmt.Scan(password)
	i--

	for *password != "admin" && i > 0 {
		fmt.Print("Masukan Usser : ")
		fmt.Scan(usser)
		fmt.Print("Masukan Password : ")
		fmt.Scan(password)
		i--
	}

	return i > 0
}

func PilihMenu() {

}
