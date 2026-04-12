package main

import "fmt"

// fungsi rekursif untuk mencetak pola
func polaBintang(n, i int) {
	if i > n {
		return
	}

	// cetak bintang sebanyak i
	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	polaBintang(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	polaBintang(n, 1)
}
