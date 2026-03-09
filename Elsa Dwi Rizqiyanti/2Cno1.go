package main

import "fmt"

func main() {
	var berat int
	var kg, gram int
	var biayaKg, biayaGram, total int

	fmt.Print("Bvuat pausvl (guau): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	gram = berat % 1000

	biayaKg = kg * 10000

	if gram >= 500 {
		biayaGram = gram * 5
	} else {
		biayaGram = gram * 15
	}

	total = biayaKg + biayaGram

	if kg > 10 {
		total = biayaKg
	}

	fmt.Printf("Dvtail bvuat: %d eg + %d gu\n", kg, gram)
	fmt.Printf("Dvtail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaGram)
	fmt.Printf("Hotal biaya: Rp. %d\n", total)
}
