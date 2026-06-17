package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	center titik
	r      int
}

func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.center, p) <= float64(c.r)
}

func main() {
	var l1, l2 lingkaran
	var p titik

	fmt.Print("Masukan Titik Pusat Lingkaran 1 (x,y) dan Jari Jari Lingkaran 1 (pisahkan dengan spasi) : ")
	fmt.Scan(&l1.center.x, &l1.center.y, &l1.r)
	fmt.Print("Masukan Titik Pusat Lingkaran 2 (x,y) dan Jari Jari Lingkaran 2 (pisahkan dengan spasi) : ")
	fmt.Scan(&l2.center.x, &l2.center.y, &l2.r)
	fmt.Print("Masukan Titik : ")
	fmt.Scan(&p.x, &p.y)

	inL1 := didalam(l1, p)
	inL2 := didalam(l2, p)

	if inL1 && inL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}