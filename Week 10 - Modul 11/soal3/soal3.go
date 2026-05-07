package main

import "fmt"

const nProv = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func inputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func provinsiTercepat(tumbuh TumbuhProv) int {
	idx := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idx] {
			idx = i
		}
	}
	return idx
}

func prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println()
	fmt.Println("=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 2.0 {
			prediksi := int((tumbuh[i] + 1) * float64(pop[i]))
			fmt.Println(prov[i], prediksi)
		}
	}
}

func indexProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv

	inputData(&prov, &pop, &tumbuh)

	var cari string
	fmt.Scan(&cari)

	fmt.Println()
	idx := provinsiTercepat(tumbuh)
	fmt.Println("Provinsi dengan angka pertumbuhan tercepat :", prov[idx])
	fmt.Println()

	idxCari := indexProvinsi(prov, cari)
	if idxCari == -1 {
		fmt.Println("Data provinsi yang dicari : tidak ditemukan")
	} else {
		fmt.Println("Data provinsi yang dicari :", prov[idxCari])
	}

	prediksi(prov, pop, tumbuh)
}
