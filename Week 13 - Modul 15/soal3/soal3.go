package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	var i int
	i = 0
	for i < n {
		if t[i].nama == nama {
			return i
		}
		i++
	}
	return -1
}

func main() {
	var p tabPartai
	var n int
	var x int
	var idx int
	var i, j int
	var key partai

	n = 0
	fmt.Println("Masukkan proses input suara :")
	fmt.Scan(&x)
	for x != -1 {
		idx = posisi(p, n, x)
		if idx == -1 {
			p[n].nama = x
			p[n].suara = 1
			n++
		} else {
			p[idx].suara = p[idx].suara + 1
		}
		fmt.Scan(&x)
	}

	for i = 1; i < n; i++ {
		key = p[i]
		j = i - 1
		for j >= 0 && p[j].suara < key.suara {
			p[j+1] = p[j]
			j = j - 1
		}
		p[j+1] = key
	}

	fmt.Println()
	fmt.Println("Hasil Perhitungan suara :")
	for i = 0; i < n; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
	fmt.Println()
}
