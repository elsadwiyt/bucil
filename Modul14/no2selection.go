package main

import "fmt"

func sortAsc(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func sortDesc(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func main() {
	var n, m, x int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var ganjil, genap []int

		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&x)

			if x%2 == 0 {
				genap = append(genap, x)
			} else {
				ganjil = append(ganjil, x)
			}
		}

		sortAsc(ganjil)
		sortDesc(genap)

		for _, v := range ganjil {
			fmt.Print(v, " ")
		}

		for j, v := range genap {
			if j == len(genap)-1 {
				fmt.Print(v)
			} else {
				fmt.Print(v, " ")
			}
		}

		fmt.Println()
	}
}
