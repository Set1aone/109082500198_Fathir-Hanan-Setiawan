package main

import "fmt"

func cetakBintang(n int) {
	if n == 0 {
		return
	}
	cetakBintang(n - 1)
	fmt.Print("*")
}

func baris(n int, max int) {
	if n > max {
		return
	}
	cetakBintang(n)
	fmt.Println()
	baris(n+1, max)
}

func main() {
	var n int
	fmt.Scan(&n)
	baris(1, n)
}
