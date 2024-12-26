package main

import (
	"fmt"
	"time"
)

var callCount int

func geserUsiaSapi(sapiProduktifUsia []int64, j int64)  {
	if j < 1 {
		return
	}
	sapiProduktifUsia[j] = sapiProduktifUsia[j-1] // Geser usia sapi ke atas (usia bertambah 1 periode/tahun)

    geserUsiaSapi(sapiProduktifUsia, j-1)
}

func HitungPolaKelahiranRekursif(populasiPeriode []int64, sapiProduktifUsia []int64, tingkatKelahiran, periode, i int64) {
	callCount++
	if i > periode {
		return
	}

	// Hitung kelahiran berdasarkan sapi produktif usia 1-4
	kelahiran := (sapiProduktifUsia[1] + sapiProduktifUsia[2] + sapiProduktifUsia[3] + sapiProduktifUsia[4]) * tingkatKelahiran / 2

	// Update populasi total
	populasiPeriode[i] = populasiPeriode[i-1] + kelahiran

	// Perbarui usia sapi produktif
	
	geserUsiaSapi(sapiProduktifUsia, 5)

	// Tambahkan sapi yang baru lahir ke usia 0
	sapiProduktifUsia[0] = kelahiran

	// Tampilkan informasi periode saat ini
	fmt.Printf("Periode %d:\n", i)
	fmt.Printf("Populasi total: %d\n", populasiPeriode[i])
	fmt.Printf("Distribusi sapi produktif berdasarkan usia: %v\n", sapiProduktifUsia)

	// Panggilan rekursif untuk periode berikutnya
	HitungPolaKelahiranRekursif(populasiPeriode, sapiProduktifUsia, tingkatKelahiran, periode, i+1)
}

func main() {
	var populasiAwal, sapiProduktif, tingkatKelahiran, periode int64

	fmt.Print("Masukkan populasi awal sapi: ")
	fmt.Scan(&populasiAwal)

	fmt.Print("Masukkan jumlah sapi produktif awal: ")
	fmt.Scan(&sapiProduktif)

	fmt.Print("Masukkan tingkat kelahiran per sapi produktif (contoh: 1): ")
	fmt.Scan(&tingkatKelahiran)

	fmt.Print("Masukkan jumlah periode (tahun): ")
	fmt.Scan(&periode)

	// Inisialisasi slice untuk menyimpan pola populasi
	populasiPeriode := make([]int64, periode+1)
	populasiPeriode[0] = populasiAwal

	// Inisialisasi distribusi sapi produktif berdasarkan usia
	sapiProduktifUsia := make([]int64, 6)
	sapiProduktifUsia[0] = sapiProduktif

	callCount = 0
	mulai := time.Now()
	// Hitung pola kelahiran secara rekursif
	HitungPolaKelahiranRekursif(populasiPeriode, sapiProduktifUsia, tingkatKelahiran, periode, 1)
	durasi := time.Since(mulai)

	// Tampilkan hasil
	fmt.Println("Pola Populasi Kelahiran Ternak Sapi per Periode:")
	for i := 1; i <= int(periode); i++ {
		fmt.Printf("Periode %d: %d sapi\n", i, populasiPeriode[i])
	}

	fmt.Printf("Jumlah panggilan rekursif: %d\n", callCount)
	fmt.Printf("\n\n\nWaktu eksekusi rekursif : %s\n\n\n", durasi)
}
