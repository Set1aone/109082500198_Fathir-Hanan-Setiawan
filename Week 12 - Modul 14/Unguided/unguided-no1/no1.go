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

		for j := 0; j < m-1; j++ {
			minIdx := j
			for k := j + 1; k < m; k++ {
				if rumah[k] < rumah[minIdx] {
					minIdx = k
				}
			}
			temp := rumah[j]
			rumah[j] = rumah[minIdx]
			rumah[minIdx] = temp
		}

		for j := 0; j < m; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[j])
		}
		fmt.Println()
	}
}
