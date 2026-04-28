package main

import "fmt"

const MAXPERTANDINGAN = 100

func main() {
	var klubA, klubB string
	var hasilPertandingan [MAXPERTANDINGAN]string
	var jumlahPertandingan int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	nomorPertandingan := 1
	jumlahPertandingan = 0

	for {
		var skorA, skorB int

		fmt.Printf("Pertandingan %d : ", nomorPertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasilPertandingan[jumlahPertandingan] = klubA
		} else if skorB > skorA {
			hasilPertandingan[jumlahPertandingan] = klubB
		} else {

			hasilPertandingan[jumlahPertandingan] = "Draw"
		}

		jumlahPertandingan = jumlahPertandingan + 1
		nomorPertandingan = nomorPertandingan + 1
	}

	for i := 0; i < jumlahPertandingan; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasilPertandingan[i])
	}

	fmt.Println("Pertandingan selesai")
}
