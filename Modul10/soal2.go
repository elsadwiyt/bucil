package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x <= 0 || x > 1000 || y <= 0 {
		return
	}

	ikan := make([]float64, x)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	wadah := make([]float64, y)

	base := x / y
	sisa := x % y

	index := 0
	for i := 0; i < y; i++ {
		count := base
		if i < sisa {
			count++
		}

		for j := 0; j < count; j++ {
			wadah[i] += ikan[index]
			index++
		}
	}

	for i := 0; i < y; i++ {
		fmt.Print(wadah[i], " ")
	}
	fmt.Println()

	var total float64
	for i := 0; i < y; i++ {
		total += wadah[i]
	}

	rata := total / float64(y)
	fmt.Println(rata)
}
