package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan Insertion Sort sesuai petunjuk soal
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		// Geser elemen-elemen yang lebih besar dari key ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Fungsi untuk menghitung median dengan pembulatan ke bawah (floor) jika jumlah genap
func hitungMedian(arr []int) int {
	n := len(arr)
	if n%2 != 0 {
		// Jika jumlah data ganjil, ambil nilai tengahnya langsung
		return arr[n/2]
	} else {
		// Jika jumlah data genap, ambil rerata dari dua nilai tengah
		// Karena pembagian integer di Go otomatis membulatkan ke bawah, kita bisa langsung pakai / 2
		tengah1 := arr[(n/2)-1]
		tengah2 := arr[n/2]
		return (tengah1 + tengah2) / 2
	}
}

func main() {
	var data []int
	var val int

	for {
		// Membaca input satu per satu secara kontinu
		_, err := fmt.Scan(&val)
		if err != nil {
			break
		}

		// Marker penanda akhir dari seluruh rangkaian input
		if val == -5313 {
			break
		}

		if val == 0 {
			// Jika membaca angka 0, saatnya mengurutkan data yang ada dan mencetak median
			if len(data) > 0 {
				insertionSort(data)
				median := hitungMedian(data)
				fmt.Println(median)
			}
		} else {
			// Jika bukan 0 dan bukan -5313, masukkan angka ke dalam kumpulan data
			data = append(data, val)
		}
	}
}
