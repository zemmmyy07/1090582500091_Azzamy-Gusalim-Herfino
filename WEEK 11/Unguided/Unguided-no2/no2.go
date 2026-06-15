package main

import "fmt"

func main() {
	var suara int
	var ttlSuara, suaraSah int
	var arrsuara [21]int

	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		ttlSuara++
		if suara >= 1 && suara <= 20 {
			suaraSah++
			arrsuara[suara]++
		}
	}

	var ketua, wakil int
	var max1, max2 int

	for i := 1; i <= 20; i++ {
		if arrsuara[i] > max1 {
			max2 = max1
			wakil = ketua
			max1 = arrsuara[i]
			ketua = i
		} else if arrsuara[i] > max2 {
			max2 = arrsuara[i]
			wakil = i
		}
	}

	fmt.Printf("Suara masuk: %d\n", ttlSuara)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}