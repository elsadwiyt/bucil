package main

import "fmt"

func tampilkanBarisan(angka int) {
	for angka != 1 {
		fmt.Print(angka, " ")
		if angka%2 == 0 {
			angka = angka / 2
		} else {
			angka = 3*angka + 1
		}
	}
	fmt.Println(1) // suku terakhir selalu 1
}

func main() {
	var nilai int
	fmt.Scan(&nilai)
	tampilkanBarisan(nilai)
}
