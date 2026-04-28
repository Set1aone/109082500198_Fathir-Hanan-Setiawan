package main

import (
	"fmt"
	"math"
)

type titik struct {
	x int
	y int
}

type lingkaran struct {
	pusat titik
	r     int
}

func jarak(p, q titik) float64 {
	a := float64(p.x - q.x)
	b := float64(p.y - q.y)
	return math.Sqrt(a*a + b*b)
}

func diDalam(c lingkaran, p titik) bool {
	j := jarak(c.pusat, p)
	return j <= float64(c.r)
}

func main() {
	var l1, l2 lingkaran
	var titikSembarang titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)

	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&titikSembarang.x, &titikSembarang.y)

	dalamL1 := diDalam(l1, titikSembarang)
	dalamL2 := diDalam(l2, titikSembarang)

	if dalamL1 && dalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
