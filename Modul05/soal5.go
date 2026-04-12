package main

import "fmt"

// fungsi rekursif
func ganjil(n, i int) {
	if i > n {
		return
	}

	fmt.Print(i, " ")
	ganjil(n, i+2)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	ganjil(n, 1)
}
