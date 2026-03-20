package main

import "fmt"

func hitungBiaya(jenis string, masuk int, keluar int) int {
	biaya := 0

	var tarifMurah, tarifMahal int
	if jenis == "motor" {
		tarifMurah = 4000
		tarifMahal = 5000
	} else {
		tarifMurah = 6000
		tarifMahal = 7000
	}

	for jam := masuk; jam < keluar; jam++ {
		if jam < 17 {
			biaya += tarifMurah
		} else {
			biaya += tarifMahal
		}
	}

	return biaya
}

func main() {
	fmt.Println("==== Rekap Tarif Parkir Cafe Per Hari ====")

	totalPendapatan := 0
	kendaraanKe := 1

	for {
		fmt.Printf("\n*Kendaraan %d\n", kendaraanKe)

		var jenis string
		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenis)

		if jenis == "-" {
			break
		}

		var masuk, keluar int
		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&masuk)
		fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&keluar)

		biaya := hitungBiaya(jenis, masuk, keluar)
		totalPendapatan += biaya

		fmt.Printf("Biaya parkir %s %d: %d\n", jenis, kendaraanKe, biaya)
		fmt.Println("========================================")

		kendaraanKe++
	}

	fmt.Printf("\n*** Total Pendapatan Parkir Hari Ini Adalah %d ***\n", totalPendapatan)
}