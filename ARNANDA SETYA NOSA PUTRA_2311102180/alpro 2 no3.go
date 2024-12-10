// Sebuah program digunakan untuk mengolah data nama provinsi, populasi, dan angka pertumbuhan penduduk provinsi di Indonesia pada tahun 2018
// Masukan terdiri dari 35 baris, yang mana masing-masing barisnya terdiri dari tiga nilal yang menyatakan nama provinsi, jumlah populasi provinsi (bilangan bulat), dan angka pertumbuhan (rill) provinsi tersebut. Pada baris terakhir hanya sebuah string yang menyatakan nama provinsi yang akan dicari.
// Keluaran terdiri dari 36 baris. Baris pertama adalah nama provinsi dengan angka pertumbuhan tercepat. Baris kedua adalah indeks provinsi yang dicari sesuai dengan nama provinsi yang ditulis pada masukan baris terakhir. Terakhir terdiri dari 34 baris yang menampilkan nama provinsi beserta prediksi jumlah penduduk pada provinsi tersebut di tahun depannya, khusus yang memiliki pertumbuhan di atas 2%.
// Lengkapi program berikut sesuai dengan spesifikasi dari subprogram yang diberikut
package main

import (
	"fmt"
	"strings"
)

const nProv = 34

type Provinsi struct {
	Nama        string
	Populasi    int
	Pertumbuhan float64
}

func ProvinsiTercepat(provinsi []Provinsi) int {
	tercepatIdx := 0
	for i := 1; i < len(provinsi); i++ {
		if provinsi[i].Pertumbuhan > provinsi[tercepatIdx].Pertumbuhan {
			tercepatIdx = i
		}
	}
	return tercepatIdx
}

func Prediksi(provinsi []Provinsi) {
	for _, p := range provinsi {
		if p.Pertumbuhan > 0.02 {
			prediksi := float64(p.Populasi) * (1 + p.Pertumbuhan)
			fmt.Printf("%s %d\n", p.Nama, int(prediksi))
		}
	}
}

func IndeksProvinsi(provinsi []Provinsi, nama string) int {
	for i, p := range provinsi {
		if strings.EqualFold(p.Nama, nama) {
			return i
		}
	}
	return -1
}

func main() {
	var provinsi []Provinsi
	var namaCari string

	for i := 0; i < nProv; i++ {
		var nama string
		var populasi int
		var pertumbuhan float64
		fmt.Scan(&nama, &populasi, &pertumbuhan)
		provinsi = append(provinsi, Provinsi{Nama: nama, Populasi: populasi, Pertumbuhan: pertumbuhan})
	}

	fmt.Scan(&namaCari)
	tercepatIdx := ProvinsiTercepat(provinsi)
	fmt.Println(provinsi[tercepatIdx].Nama)
	indeks := IndeksProvinsi(provinsi, namaCari)
	fmt.Println(indeks)
	Prediksi(provinsi)
}
