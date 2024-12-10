package main

import (
	"fmt"
)

const nMax = 51

type Mahasiswa struct {
	NIM   string
	Nama  string
	Nilai int
}

type ArrayMahasiswa [nMax]Mahasiswa

func main() {
	var dataMahasiswa ArrayMahasiswa
	var N int

	fmt.Println("=== Program Pencarian Data Mahasiswa ===")
	fmt.Print("Masukkan jumlah mahasiswa (maksimal 50): ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("\nMasukkan data mahasiswa ke-%d:\n", i+1)
		fmt.Print("  NIM: ")
		fmt.Scan(&dataMahasiswa[i].NIM)
		fmt.Print("  Nama: ")
		fmt.Scan(&dataMahasiswa[i].Nama)
		fmt.Print("  Nilai: ")
		fmt.Scan(&dataMahasiswa[i].Nilai)
	}

	fmt.Println("\n=== Pencarian Data ===")
	var cariNIM string
	fmt.Print("Masukkan NIM mahasiswa untuk mencari nilai pertama: ")
	fmt.Scan(&cariNIM)
	nilaiPertama := cariNilaiPertama(dataMahasiswa, N, cariNIM)
	if nilaiPertama != -1 {
		fmt.Printf("\nNilai pertama mahasiswa dengan NIM '%s' adalah: %d\n", cariNIM, nilaiPertama)
	} else {
		fmt.Printf("\nMaaf, mahasiswa dengan NIM '%s' tidak ditemukan.\n", cariNIM)
	}

	fmt.Print("\nMasukkan NIM mahasiswa untuk mencari nilai terbesar: ")
	fmt.Scan(&cariNIM)
	nilaiTerbesar := cariNilaiTerbesar(dataMahasiswa, N, cariNIM)
	if nilaiTerbesar != -1 {
		fmt.Printf("\nNilai terbesar mahasiswa dengan NIM '%s' adalah: %d\n", cariNIM, nilaiTerbesar)
	} else {
		fmt.Printf("\nMaaf, mahasiswa dengan NIM '%s' tidak ditemukan.\n", cariNIM)
	}

	fmt.Println("\n=== Program Selesai ===")
}

func cariNilaiPertama(data ArrayMahasiswa, N int, cariNIM string) int {
	for i := 0; i < N; i++ {
		if data[i].NIM == cariNIM {
			return data[i].Nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(data ArrayMahasiswa, N int, cariNIM string) int {
	nilaiMax := -1
	found := false
	for i := 0; i < N; i++ {
		if data[i].NIM == cariNIM {
			found = true
			if data[i].Nilai > nilaiMax {
				nilaiMax = data[i].Nilai
			}
		}
	}
	if !found {
		return -1
	}
	return nilaiMax
}
