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

	fmt.Printf("Suara masuk: %d\n", ttlSuara)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	for i := 1; i <= 20; i++ {
		if arrsuara[i] > 0 {
			fmt.Printf("%d: %d\n", i, arrsuara[i])
		}
	}
}