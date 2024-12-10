package main

import (
	"fmt"
	"sort"
)

type Partai struct {
	nama  int
	suara int
}

func main() {
	const MAX = 1000000
	var suara [MAX + 1]int
	var input int

	fmt.Println("=== Program Perhitungan Suara Partai ===")
	fmt.Println("Masukkan nomor partai (1 - 1000000), akhiri dengan -1:")

	// Input data suara
	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}
		if input < 1 || input > MAX {
			fmt.Println("Nomor partai tidak valid. Harap masukkan nomor partai antara 1 dan 1000000.")
			continue
		}
		suara[input]++
	}

	// Proses hasil suara
	var hasil []Partai
	for i := 1; i <= MAX; i++ {
		if suara[i] > 0 {
			hasil = append(hasil, Partai{nama: i, suara: suara[i]})
		}
	}

	// Urutkan hasil berdasarkan suara (menurun) dan nomor partai (meningkat jika suara sama)
	sort.Slice(hasil, func(i, j int) bool {
		if hasil[i].suara == hasil[j].suara {
			return hasil[i].nama < hasil[j].nama
		}
		return hasil[i].suara > hasil[j].suara
	})

	// Output hasil
	fmt.Println("\nHasil Perhitungan Suara:")
	if len(hasil) == 0 {
		fmt.Println("Tidak ada suara yang dihitung.")
	} else {
		for i, p := range hasil {
			fmt.Printf("%d. Partai %d: %d suara\n", i+1, p.nama, p.suara)
		}
	}

	fmt.Println("\n=== Program Selesai ===")
}
