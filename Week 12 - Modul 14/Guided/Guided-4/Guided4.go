package main

import "fmt"

type mahasiswa struct {
	nama, nim string
	IPK       float64
}

type arrMHS [100]mahasiswa

func insertionSortStruct(T *arrMHS, n int) {
	var temp mahasiswa
	var i, j int
	for i = 1; i <= n-1; i++ {
		temp = T[i]
		j = i - 1
		for j >= 0 && T[j].nama < temp.nama {
			T[j+1] = T[j]
			j = j - 1
		}
		T[j+1] = temp
	}
}

func main() {
	var data arrMHS
	var n, i int

	fmt.Print("Masukkan jumlah mahasiswa : ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Println("\nData Mahasiswa ke-", i+1)

		fmt.Print("Nama : ")
		fmt.Scan(&data[i].nama)

		fmt.Print("NIM : ")
		fmt.Scan(&data[i].nim)

		fmt.Print("IPK : ")
		fmt.Scan(&data[i].IPK)
	}

	fmt.Println("\nData sebelum sorting : ")
	for i = 0; i < n; i++ {
		fmt.Println(data[i].nama, data[i].nim, data[i].IPK)
	}

	insertionSortStruct(&data, n)

	fmt.Println("\nData setelah sorting (Descending berdasarkan nama) : ")
	for i = 0; i < n; i++ {
		fmt.Println(data[i].nama, data[i].nim, data[i].IPK)
	}
}
