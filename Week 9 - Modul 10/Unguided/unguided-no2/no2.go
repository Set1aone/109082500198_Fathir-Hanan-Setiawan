package main

import "fmt"

func main() {
	var x, y int
	var berat [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}

	for i := 0; i < jumlahWadah; i++ {
		total := 0.0
		ikanDalamWadah := y

		if i == jumlahWadah-1 && x%y != 0 {
			ikanDalamWadah = x % y
		}

		for j := 0; j < ikanDalamWadah; j++ {
			total += berat[i*y+j]
		}
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	for i := 0; i < jumlahWadah; i++ {
		total := 0.0
		ikanDalamWadah := y

		if i == jumlahWadah-1 && x%y != 0 {
			ikanDalamWadah = x % y
		}

		for j := 0; j < ikanDalamWadah; j++ {
			total += berat[i*y+j]
		}
		rataRata := total / float64(ikanDalamWadah)
		fmt.Printf("%.2f ", rataRata)
	}
	fmt.Println()
}
