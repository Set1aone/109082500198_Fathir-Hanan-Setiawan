package main

import "fmt"

const NMAX = 1000001

type arrInt [NMAX]int

func SelectionSort(T *arrInt, n int) {
	var i, j, idxMin, temp int
	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			if T[j] < T[idxMin] {
				idxMin = j
			}
		}
		temp = T[i]
		T[i] = T[idxMin]
		T[idxMin] = temp
	}
}

func median(T arrInt, n int) float64 {
	var med float64
	if n%2 == 1 {
		med = float64(T[n/2])
	} else {
		med = float64(T[n/2-1]+T[n/2]) / 2.0
	}
	return med
}

func main() {
	var A arrInt
	var x int
	var n int = 0

	fmt.Println("Input data masukan :")
	fmt.Scan(&x)
	for x != -5313541 && n < NMAX {
		if x == 0 {
			SelectionSort(&A, n)
			fmt.Println("Median :")
			fmt.Println(median(A, n))
		} else {
			A[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}