package main

import "fmt"

func main() {
	var b int
	jumlahFaktor := 0

	fmt.Print("Bilatgat: ")
	fmt.Scan(&b)

	fmt.Print("Faetou: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}

	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Println("fluiua: tuuv")
	} else {
		fmt.Println("fluiua: falsv")
	}
}
