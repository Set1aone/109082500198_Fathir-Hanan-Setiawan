package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		ganjil := []int{}
		genap := []int{}

		for j := 0; j < m; j++ {
			if rumah[j]%2 != 0 {
				ganjil = append(ganjil, rumah[j])
			} else {
				genap = append(genap, rumah[j])
			}
		}

		for j := 0; j < len(ganjil)-1; j++ {
			minIdx := j
			for k := j + 1; k < len(ganjil); k++ {
				if ganjil[k] < ganjil[minIdx] {
					minIdx = k
				}
			}
			temp := ganjil[j]
			ganjil[j] = ganjil[minIdx]
			ganjil[minIdx] = temp
		}

		for j := 0; j < len(genap)-1; j++ {
			maxIdx := j
			for k := j + 1; k < len(genap); k++ {
				if genap[k] > genap[maxIdx] {
					maxIdx = k
				}
			}
			temp := genap[j]
			genap[j] = genap[maxIdx]
			genap[maxIdx] = temp
		}

		hasil := append(ganjil, genap...)

		for j := 0; j < len(hasil); j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(hasil[j])
		}
		fmt.Println()
	}
}
