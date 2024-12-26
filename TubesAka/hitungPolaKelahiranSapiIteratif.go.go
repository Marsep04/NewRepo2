package main

import (
	"fmt"
	"time"
)

func HitungPolaKelahiran(populasiPeriode []int64 , populasiAwal, sapiProduktif, tingkatKelahiran, periode int64) {

	// Slice untuk melacak jumlah sapi produktif berdasarkan umur
	sapiProduktifUsia := make([]int64, 6) // Indeks 0-5 untuk usia 0-5 periode, di indeks 5 dan didiluar indeks untuk sapi yang tidak produktif

	populasiPeriode[0] = populasiAwal
	sapiProduktifUsia[0] = sapiProduktif // Semua sapi produktif awalnya berumur 0 periode

	for i := 1; i <= int(periode); i++ {
		// Hitung kelahiran berdasarkan sapi produktif usia 1-4 (hanya sapi usia 1-4 yang bisa beranak)
		kelahiran := (sapiProduktifUsia[1] + sapiProduktifUsia[2] + sapiProduktifUsia[3] + sapiProduktifUsia[4]) * tingkatKelahiran / 2

		// Update populasi total
		populasiPeriode[i] = populasiPeriode[i-1] + kelahiran

		// Perbarui usia sapi produktif
		for j := 5; j > 0; j-- {
			sapiProduktifUsia[j] = sapiProduktifUsia[j-1] // Geser usia sapi ke atas (usia bertambah 1 periode/tahun)
		}
	
		// Tambahkan sapi yang baru lahir ke usia 0
		sapiProduktifUsia[0] = kelahiran	

		// Sapi yang sudah lebih dari 5 periode menjadi tidak produktif
		// sapiProduktifUsia[5] sudah mencatat sapi tidak produktif dari iterasi sebelumnya
		fmt.Printf("Periode %d:\n", i)
		fmt.Printf("Populasi total: %d\n", populasiPeriode[i])
		fmt.Printf("Distribusi sapi produktif berdasarkan usia: %v\n", sapiProduktifUsia)
	}
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
	
	// Hitung pola kelahiran
	populasiPeriode := make([]int64, periode+1)

	mulai := time.Now()
	HitungPolaKelahiran(populasiPeriode, populasiAwal, sapiProduktif, tingkatKelahiran, periode)
	durasi := time.Since(mulai)

	// Tampilkan hasil
	fmt.Println("Pola Populasi Kelahiran Ternak Sapi per Periode:")
	for i := 1; i <= int(periode); i++ {
		fmt.Printf("Periode %d: %d sapi\n", i, populasiPeriode[i])
	}

	fmt.Printf("\n\n\nWaktu eksekusi iteratif: %s\n\n\n", durasi)
}
