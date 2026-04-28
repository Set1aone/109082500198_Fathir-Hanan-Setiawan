package main

import (
	"fmt"
	"math"
)

const MAXN = 100

func main() {
	var arr [MAXN]int
	var n int

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan", n, "elemen:")
	for i := 0; i < n; i++ {
		fmt.Printf("arr[%d] = ", i)
		fmt.Scan(&arr[i])
	}

	var pilihan string
	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("a. Tampilkan semua isi array")
		fmt.Println("b. Tampilkan elemen indeks ganjil")
		fmt.Println("c. Tampilkan elemen indeks genap")
		fmt.Println("d. Tampilkan elemen indeks kelipatan x")
		fmt.Println("e. Hapus elemen pada indeks tertentu")
		fmt.Println("f. Tampilkan rata-rata")
		fmt.Println("g. Tampilkan standar deviasi")
		fmt.Println("h. Tampilkan frekuensi bilangan tertentu")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == "0" {
			fmt.Println("Program selesai.")
			break
		}

		switch pilihan {
		case "a":
			fmt.Print("Semua elemen: ")
			for i := 0; i < n; i++ {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()

		case "b":
			fmt.Print("Elemen indeks ganjil: ")
			for i := 1; i < n; i += 2 {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()

		case "c":
			fmt.Print("Elemen indeks genap: ")
			for i := 0; i < n; i += 2 {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()

		case "d":
			var x int
			fmt.Print("Masukkan nilai x: ")
			fmt.Scan(&x)
			fmt.Print("Elemen indeks kelipatan", x, ": ")
			for i := 0; i < n; i++ {
				if i%x == 0 {
					fmt.Print(arr[i], " ")
				}
			}
			fmt.Println()

		case "e":
			var idx int
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			fmt.Scan(&idx)

			for i := idx; i < n-1; i++ {
				arr[i] = arr[i+1]
			}
			n = n - 1

			fmt.Print("Array setelah dihapus: ")
			for i := 0; i < n; i++ {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()

		case "f":
			jumlah := 0
			for i := 0; i < n; i++ {
				jumlah = jumlah + arr[i]
			}
			rataRata := float64(jumlah) / float64(n)
			fmt.Printf("Rata-rata: %.2f\n", rataRata)

		case "g":
			jumlah := 0
			for i := 0; i < n; i++ {
				jumlah = jumlah + arr[i]
			}
			mean := float64(jumlah) / float64(n)

			jumlahKuadrat := 0.0
			for i := 0; i < n; i++ {
				selisih := float64(arr[i]) - mean
				jumlahKuadrat = jumlahKuadrat + selisih*selisih
			}

			standarDeviasi := math.Sqrt(jumlahKuadrat / float64(n))
			fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi)

		case "h":
			var cari int
			fmt.Print("Masukkan bilangan yang dicari: ")
			fmt.Scan(&cari)

			frekuensi := 0
			for i := 0; i < n; i++ {
				if arr[i] == cari {
					frekuensi = frekuensi + 1
				}
			}
			fmt.Printf("Frekuensi %d di dalam array: %d\n", cari, frekuensi)

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
