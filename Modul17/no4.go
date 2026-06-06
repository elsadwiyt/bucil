package main

import "fmt"

func main() {
	var n int
	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	var s float64 = 0.0

	for i := 1; i <= n; i++ {
		penyebut := float64(2*i - 1)

		if i%2 != 0 {
			s += 1.0 / penyebut
		} else {
			s -= 1.0 / penyebut
		}
	}

	pi := 4.0 * s
	fmt.Printf("Hasil PI: %.6f\n", pi)
}
