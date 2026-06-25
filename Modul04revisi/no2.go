package main

import "fmt"

func hitungSkor(waktu [8]int, soal *int, skor *int) {
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		if waktu[i] <= 300 {
			*soal++
			*skor += waktu[i]
		}
	}
}

func main() {
	var nama string

	var pemenang string
	maxSoal := -1
	minWaktu := 999999

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		var waktu [8]int
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		var soal, skor int
		hitungSkor(waktu, &soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			maxSoal = soal
			minWaktu = skor
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}
