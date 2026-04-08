package main

import "fmt"

// Fungsi untuk menghitung biaya parkir
func hitungBiaya(jenis string, masuk int, keluar int) int {
	biaya := 0

	// Looping per jam untuk mengecek tarif yang berlaku di jam tersebut
	for jam := masuk; jam < keluar; jam++ {
		switch jenis {
		case "motor":
			// Di bawah jam 17 (jam 5 sore), tarif 4000. Mulai jam 17 ke atas, tarif 5000.
			if jam < 17 {
				biaya += 4000
			} else {
				biaya += 5000
			}
		case "mobil":
			// Di bawah jam 17 (jam 5 sore), tarif 6000. Mulai jam 17 ke atas, tarif 7000.
			if jam < 17 {
				biaya += 6000
			} else {
				biaya += 7000
			}
		}
	}

	return biaya
}

func main() {
	var jenis string
	var masuk, keluar int
	total := 0
	no := 1

	fmt.Println("==== Rekap Tarif Parkir Cafe Per Hari ====")

	for {
		fmt.Printf("\n*Kendaraan %d\n", no)
		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenis)

		// Kondisi untuk menghentikan input jika user mengetik "-"
		if jenis == "-" {
			break
		}

		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&masuk)

		fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&keluar)

		// Memanggil fungsi hitungBiaya
		biaya := hitungBiaya(jenis, masuk, keluar)

		fmt.Printf("Biaya parkir %s %d: %d\n", jenis, no, biaya)
		fmt.Println("================================")

		// Menambahkan biaya kendaraan saat ini ke total pendapatan harian
		total += biaya
		no++
	}

	// Menampilkan total pendapatan setelah loop selesai
	fmt.Printf("\n*** Total Pendapatan Parkir Hari Ini Adalah %d ***\n", total)
}