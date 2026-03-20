package main

import "fmt"

const PI = 3.141592653589793

func hitungPersegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi
	fmt.Println("Luas persegi    :", luas)
	fmt.Println("Keliling persegi:", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	fmt.Println("Luas persegi panjang    :", luas)
	fmt.Println("Keliling persegi panjang:", keliling)
}

func hitungLingkaran(jarijari float64) {
	luas := PI * jarijari * jarijari
	keliling := 2 * PI * jarijari
	fmt.Printf("Luas lingkaran    : %f\n", luas)
	fmt.Printf("Keliling lingkaran: %g\n", keliling)
}

func main() {
	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")

	var pilihan int
	fmt.Scan(&pilihan)
	fmt.Println()

	switch pilihan {
	case 1:
		var sisi int
		fmt.Print("Masukkan sisi : ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)
	case 2:
		var panjang, lebar int
		fmt.Print("Masukkan panjang : ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar   : ")
		fmt.Scan(&lebar)
		hitungPersegiPanjang(panjang, lebar)
	case 3:
		var jarijari float64
		fmt.Print("Masukkan jari-jari : ")
		fmt.Scan(&jarijari)
		hitungLingkaran(jarijari)
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}