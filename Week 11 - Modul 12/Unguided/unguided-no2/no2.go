package main

import "fmt"

func main() {
	var suaraMasuk int
	var suaraSah int
	var calon [21]int

	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		suaraMasuk++
		if n >= 1 && n <= 20 {
			suaraSah++
			calon[n]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	ketua := 0
	for i := 1; i <= 20; i++ {
		if ketua == 0 || calon[i] > calon[ketua] {
			ketua = i
		}
	}

	wakil := 0
	for i := 1; i <= 20; i++ {
		if i == ketua {
			continue
		}
		if calon[i] == calon[ketua] {
			wakil = i
			break
		}
	}

	if wakil == 0 {
		for i := 1; i <= 20; i++ {
			if i == ketua {
				continue
			}
			if wakil == 0 || calon[i] > calon[wakil] {
				wakil = i
			}
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
