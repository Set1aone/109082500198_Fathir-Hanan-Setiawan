package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	*n = 0
	var s string
	fmt.Print("Teks : ")
	for {
		fmt.Scan(&s)

		if s == "." {
			break
		}
		t[*n] = rune(s[0])
		*n = *n + 1
		if *n >= NMAX {
			break
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	hasilPalindrom := palindrom(tab, m)

	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)

	fmt.Print("palindrom : ")
	if hasilPalindrom {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
