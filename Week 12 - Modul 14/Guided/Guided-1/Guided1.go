package main

import "fmt"

func selectionSortArray(arr *[8]int) {
	var idx_min, i, j int
	for i = 0; i < len(arr); i++ {
		idx_min = i
		for j = i + 1; j < len(arr); j++ {
			if arr[j] < arr[idx_min] {
				idx_min = j
			}
		}
		if idx_min != i {
			arr[i], arr[idx_min] = arr[idx_min], arr[i]
		}
	}
}

func main() {
	var arrAngka [8]int

	for i := 0; i < len(arrAngka); i++ {
		fmt.Printf("Masukkan angka index ke-%d: ", i)
		fmt.Scan(&arrAngka[i])
	}
	fmt.Println()

	fmt.Println("=== SEBELUM SORTING ===")
	for i := 0; i < len(arrAngka); i++ {
		fmt.Print(arrAngka[i], " ")
	}
	fmt.Println()

	selectionSortArray(&arrAngka)
	fmt.Println("=== SETELAH SORTING ===")
	for i := 0; i < len(arrAngka); i++ {
		fmt.Print(arrAngka[i], " ")
	}
}
