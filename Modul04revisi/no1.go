package main

import (
	"fmt"
)

func hitungFaktorial(angka int) int {
	total := 1
	for i := 2; i <= angka; i++ {
		total *= i
	}
	return total
}

func hitungPermutasi(n, r int) int {
	return hitungFaktorial(n) / hitungFaktorial(n-r)
}

func hitungKombinasi(n, r int) int {
	return hitungFaktorial(n) / (hitungFaktorial(r) * hitungFaktorial(n-r))
}

func main() {
	var w, x, y, z int
	fmt.Scan(&w, &x, &y, &z) // input w x y z

	hasilPermutasi1 := hitungPermutasi(w, y)
	hasilKombinasi1 := hitungKombinasi(w, y)
	fmt.Println(hasilPermutasi1, hasilKombinasi1)

	hasilPermutasi2 := hitungPermutasi(x, z)
	hasilKombinasi2 := hitungKombinasi(x, z)
	fmt.Println(hasilPermutasi2, hasilKombinasi2)
}
