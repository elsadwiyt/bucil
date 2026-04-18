package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang [100]string
	var n int = 0

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	i := 1
	for {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[n] = klubA
			n++
		} else if skorB > skorA {
			pemenang[n] = klubB
			n++
		} else {
			pemenang[n] = "Draw"
			n++
		}

		i++
	}

	for j := 0; j < n; j++ {
		fmt.Printf("Hasil %d : %s\n", j+1, pemenang[j])
	}

	fmt.Println("Pertandingan selesai")
}
