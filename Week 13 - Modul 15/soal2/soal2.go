package main

import "fmt"

const NMAX = 1001

type Pemain struct {
	nama1  string
	nama2  string
	gol    int
	assist int
}

type arrPemain [NMAX]Pemain

func SelectionSort(T *arrPemain, n int) {
	var i, j, idxMax int
	var temp Pemain
	for i = 0; i < n-1; i++ {
		idxMax = i
		for j = i + 1; j < n; j++ {
			if T[j].gol > T[idxMax].gol {
				idxMax = j
			} else if T[j].gol == T[idxMax].gol {
				if T[j].assist > T[idxMax].assist {
					idxMax = j
				}
			}
		}
		temp = T[i]
		T[i] = T[idxMax]
		T[idxMax] = temp
	}
}

func main() {
	var data arrPemain
	var n, i int

	fmt.Println("Masukkan Data Input :")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&data[i].nama1)
		fmt.Scan(&data[i].nama2)
		fmt.Scan(&data[i].gol)
		fmt.Scan(&data[i].assist)
	}

	SelectionSort(&data, n)

	fmt.Println()
	fmt.Println("Hasil Sorting :")
	for i = 0; i < n; i++ {
		fmt.Println(data[i].nama1, data[i].nama2, data[i].gol, data[i].assist)
	}
}
