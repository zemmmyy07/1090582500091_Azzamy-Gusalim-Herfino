package main

import "fmt"

const nProv = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIdx := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("--- Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ---")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := (tumbuh[i] + 1.0) * float64(pop[i])
			fmt.Printf("%s %d\n", prov[i], int(prediksi))
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var cariProv string

	fmt.Println("--- Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ---")
	InputData(&prov, &pop, &tumbuh)

	fmt.Scan(&cariProv)

	idxTercepat := ProvinsiTercepat(tumbuh)
	fmt.Println("\nProvinsi dengan angka pertumbuhan tercepat :", prov[idxTercepat])

	idxCari := IndeksProvinsi(prov, cariProv)
	fmt.Printf("Data provinsi yang dicari : %s (Indeks ke-%d)\n\n", cariProv, idxCari)

	Prediksi(prov, pop, tumbuh)
}