package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)

	if tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0 {
		fmt.Print("Kabisat: ")
		fmt.Print("True")
	} else {
		fmt.Print("Kabisat: ")
		fmt.Print("False")
	}
}
