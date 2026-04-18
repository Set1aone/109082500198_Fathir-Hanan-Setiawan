package main

import "fmt"

type angka int
type kata string

type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	var b Buku
	var judul, penulis, penerbit string
	var tahun, halaman int

	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukkan judul buku: ")
	fmt.Scan(&judul)
	b.judul = kata(judul)

	fmt.Print("Masukkan nama penulis: ")
	fmt.Scan(&penulis)
	b.penulis = kata(penulis)

	fmt.Print("Masukkan penerbit: ")
	fmt.Scan(&penerbit)
	b.penerbit = kata(penerbit)

	fmt.Print("Masukkan tahun terbit: ")
	fmt.Scan(&tahun)
	b.tahunTerbit = angka(tahun)

	fmt.Print("Masukkan jumlah halaman: ")
	fmt.Scan(&halaman)
	b.jumlahHalaman = angka(halaman)

	fmt.Println()
	fmt.Println("=== BIODATA BUKU ===")
	fmt.Println("Judul Buku:", b.judul)
	fmt.Println("Penulis:", b.penulis)
	fmt.Println("Penerbit:", b.penerbit)
	fmt.Println("Tahun Terbit:", b.tahunTerbit)
	fmt.Println("Jumlah Halaman:", b.jumlahHalaman)
}
