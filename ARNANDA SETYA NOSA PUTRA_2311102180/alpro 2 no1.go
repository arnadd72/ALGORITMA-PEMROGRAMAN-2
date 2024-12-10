package main

import "fmt"

type set [2022]int

// Fungsi untuk mengecek apakah elemen sudah ada dalam himpunan
func exist(T []int, val int) bool {
	for _, v := range T {
		if v == val {
			return true
		}
	}
	return false
}

// Fungsi untuk membaca himpunan dari input pengguna
func inputSet(nama string) (set, int) {
	var T set
	var n int
	fmt.Printf("Masukkan elemen himpunan %s (akhiri dengan memasukkan elemen yang sudah pernah dimasukkan):\n", nama)
	for {
		var x int
		fmt.Scan(&x)
		if exist(T[:n], x) {
			fmt.Printf("Elemen %d sudah ada di himpunan %s, input selesai.\n", x, nama)
			break
		}
		T[n] = x
		n++
	}
	return T, n
}

// Fungsi untuk menemukan irisan dua himpunan
func findIntersection(T1 []int, n1 int, T2 []int, n2 int) (set, int) {
	var result set
	var k int
	for i := 0; i < n1; i++ {
		if exist(T2[:n2], T1[i]) {
			result[k] = T1[i]
			k++
		}
	}
	return result, k
}

// Fungsi untuk mencetak elemen-elemen himpunan
func printSet(T []int, n int) {
	if n == 0 {
		fmt.Println("Himpunan kosong.")
		return
	}
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(T[i])
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== Program Mencari Irisan Himpunan ===")
	var s1, s2, s3 set
	var n1, n2, n3 int

	// Input himpunan pertama
	s1, n1 = inputSet("A")

	// Input himpunan kedua
	s2, n2 = inputSet("B")

	// Cari irisan kedua himpunan
	s3, n3 = findIntersection(s1[:], n1, s2[:], n2)

	// Tampilkan hasil irisan
	fmt.Println("\nIrisan dari kedua himpunan adalah:")
	printSet(s3[:], n3)
	fmt.Println("=== Program Selesai ===")
}
