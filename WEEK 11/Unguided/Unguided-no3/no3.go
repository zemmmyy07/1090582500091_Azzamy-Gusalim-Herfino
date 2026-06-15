package main

import "fmt"

const NMAX = 1000000
var data [NMAX]int

func isiArray(n int) {
	fmt.Print("Masukan data: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var kr = 0
	var kn = n - 1
	var med int
	var found = -1

	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if k < data[med] {
			kn = med - 1
		} else if k > data[med] {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}

func main() {
	var n, k int
	fmt.Print("Masukan jumlah data dan data yang ingin dicari: ")
	fmt.Scan(&n, &k)
	
	isiArray(n)
	hasil := posisi(n, k)
	
	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Printf("Data %d berada di posisi ke-%d", k,hasil)
	}
}