package main

import "fmt"

func main() {
	// Nilai tagihan yang akan diuji
	tagihan1 := 275.0
	tagihan2 := 40.0
	tagihan3 := 430.0

	// Fungsi untuk menghitung tip
	hitungTip := func(tagihan float64) float64 {
		var tip float64
		if tagihan >= 50 && tagihan <= 300 {
			tip = tagihan * 0.15
		} else {
			tip = tagihan * 0.20
		}
		return tip
	}

	// Menghitung tip dan total untuk setiap nilai tagihan
	tip1 := hitungTip(tagihan1)
	total1 := tagihan1 + tip1

	tip2 := hitungTip(tagihan2)
	total2 := tagihan2 + tip2

	tip3 := hitungTip(tagihan3)
	total3 := tagihan3 + tip3

	// Menampilkan hasil
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan1, tip1, total1)
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan2, tip2, total2)
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan3, tip3, total3)
}
