package main

import "fmt"

type suhu float64

func CelciusToReamur(Celcius suhu) suhu {
	return Celcius * 4 / 5
}

func CelciusToFahrenheit(Celcius suhu) suhu {
	return Celcius*9/5 + 32
}

func CelciusToKelvin(Celcius suhu) suhu {
	return Celcius + 273.15
}

func main() {
	var c suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius): ")
	fmt.Scan(&c)

	fmt.Println()
	fmt.Println(c, "celcius =", CelciusToReamur(c), "reamur")
	fmt.Println(c, "celcius =", CelciusToFahrenheit(c), "fahrenheit")
	fmt.Println(c, "celcius =", CelciusToKelvin(c), "kelvin")
}
