package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 || n >= 1000 {
		return
	}

	hasil := make([][]int, n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		daerah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&daerah[j])
		}

		selectionSort(daerah)

		hasil[i] = daerah
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(hasil[i]); j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(hasil[i][j])
		}
		fmt.Println()
	}
}
