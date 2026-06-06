package main

import "fmt"

type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

var pustaka []Buku
var nPustaka int

func daftarkanBuku() {
	fmt.Scan(&nPustaka)
	pustaka = make([]Buku, nPustaka)
	for i := 0; i < nPustaka; i++ {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
	}
}

func cetakTerfavorit() {
	maxIdx := 0
	for i := 1; i < nPustaka; i++ {
		if pustaka[i].rating > pustaka[maxIdx].rating {
			maxIdx = i
		}
	}
	fmt.Println(pustaka[maxIdx].judul, pustaka[maxIdx].penulis, pustaka[maxIdx].penerbit, pustaka[maxIdx].tahun)
}

func urutBuku() {
	for i := 1; i < nPustaka; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func cetak5Terbaru() {
	limit := 5
	if nPustaka < 5 {
		limit = nPustaka
	}
	for i := 0; i < limit; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func cariBuku(r int) {
	lo := 0
	hi := nPustaka - 1
	ketemu := -1
	for lo <= hi {
		mid := (lo + hi) / 2
		if pustaka[mid].rating == r {
			ketemu = mid
			break
		} else if pustaka[mid].rating > r {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if ketemu == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		fmt.Println(pustaka[ketemu].judul, pustaka[ketemu].penulis, pustaka[ketemu].penerbit, pustaka[ketemu].tahun, pustaka[ketemu].eksemplar, pustaka[ketemu].rating)
	}
}

func main() {
	daftarkanBuku()
	cetakTerfavorit()
	urutBuku()
	cetak5Terbaru()
	var r int
	fmt.Scan(&r)
	cariBuku(r)
}
