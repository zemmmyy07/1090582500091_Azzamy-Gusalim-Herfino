package main

import (
	"fmt"
	"math"
)

const NMAX int = 100

func main() {
	var arr [NMAX]int
	var n, x, hapusIdx, cariFreq int

	fmt.Print("Masukkan jumlah elemen N: ")
	fmt.Scan(&n)

	fmt.Printf("Masukkan %d buah bilangan bulat:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("\na. Isi seluruh array: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("b. Elemen indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Elemen indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Masukkan nilai X untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("d. Elemen indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if x != 0 && i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Print("Masukkan indeks elemen yang akan dihapus: ")
	fmt.Scan(&hapusIdx)
	if hapusIdx >= 0 && hapusIdx < n {
		for i := hapusIdx; i < n-1; i++ {
			arr[i] = arr[i+1]
		}
		n-- 
		fmt.Print("e. Isi array setelah dihapus: ")
		for i := 0; i < n; i++ {
			fmt.Print(arr[i], " ")
		}
		fmt.Println()
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		sum += float64(arr[i])
	}
	rata := sum / float64(n)
	fmt.Printf("f. Rata-rata: %.2f\n", rata)

	varianSum := 0.0
	for i := 0; i < n; i++ {
		varianSum += math.Pow(float64(arr[i])-rata, 2)
	}
	stdDev := math.Sqrt(varianSum / float64(n))
	fmt.Printf("g. Standar Deviasi: %.2f\n", stdDev)

	fmt.Print("Masukkan bilangan untuk dicari frekuensinya: ")
	fmt.Scan(&cariFreq)
	freq := 0
	for i := 0; i < n; i++ {
		if arr[i] == cariFreq {
			freq++
		}
	}
	fmt.Printf("h. Frekuensi bilangan %d adalah: %d kali\n", cariFreq, freq)
}