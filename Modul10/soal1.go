package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 || n > 1000 {
		return
	}

	arr := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	min := arr[0]
	max := arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Println(min, max)
}
