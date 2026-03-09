package main

import "fmt"

func main() {
	var K int

	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	akar2 := 1.0
	for k := 0; k <= K; k++ {
		kf := float64(k)

		pembilang := (4*kf + 2) * (4*kf + 2)
		penyebut := (4*kf + 1) * (4*kf + 3)
		fk := pembilang / penyebut

		akar2 *= fk
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}
