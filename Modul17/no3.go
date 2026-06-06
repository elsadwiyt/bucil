package main

import "fmt"

func main() {
	var totalTetesan int

	fmt.Print("Masukkan banyaknya tetesan air hujan: ")
	fmt.Scan(&totalTetesan)

	countA := 0
	countB := 0
	countC := 0
	countD := 0

	var seed uint64 = 123456789
	var a uint64 = 1103515245
	var c uint64 = 12345
	var m uint64 = 1 << 31

	for i := 0; i < totalTetesan; i++ {
		seed = (a*seed + c) % m
		x := float64(seed) / float64(m)

		seed = (a*seed + c) % m
		y := float64(seed) / float64(m)

		if x <= 0.5 && y <= 0.5 {
			countA++
		} else if x > 0.5 && y <= 0.5 {
			countB++
		} else if x > 0.5 && y > 0.5 {
			countC++
		} else if x <= 0.5 && y > 0.5 {
			countD++
		}
	}

	const konversi = 0.0001

	curahA := float64(countA) * konversi
	curahB := float64(countB) * konversi
	curahC := float64(countC) * konversi
	curahD := float64(countD) * konversi

	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", curahA)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", curahB)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", curahC)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", curahD)
}
