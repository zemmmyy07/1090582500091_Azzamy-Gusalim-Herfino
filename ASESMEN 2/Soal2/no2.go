package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func nilaiPertama(arr arrayMahasiswa, n int, nimCari string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == nimCari {
			return arr[i].nilai 
		}
	}
	return -1
}

func nilaiTerbesar(arr arrayMahasiswa, n int, nimCari string) int {
	max := -1
	for i := 0; i < n; i++ {
		if arr[i].NIM == nimCari {
			if arr[i].nilai > max {
				max = arr[i].nilai
			}
		}
	}
	return max
}

func main() {
	var data arrayMahasiswa
	var n int

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&data[i].NIM, &data[i].nama, &data[i].nilai)
	}

	var searchNIM string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&searchNIM)

	pertama := nilaiPertama(data, n, searchNIM)
	terbesar := nilaiTerbesar(data, n, searchNIM)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", searchNIM, pertama)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", searchNIM, terbesar)
}