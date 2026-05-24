package main

import "fmt"

const NMAX = 100

func insertionSort(arr *[NMAX]int, n int) {
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = temp
	}
}

func cekJarakTetap(arr [NMAX]int, n int) bool {
	if n <= 2 {
		return true
	}

	selisih := arr[1] - arr[0]

	for i := 2; i < n; i++ {
		if arr[i]-arr[i-1] != selisih {
			return false
		}
	}

	return true
}

func main() {
	var arr [NMAX]int
	var x, n int

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		arr[n] = x
		n++
	}

	insertionSort(&arr, n)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	if cekJarakTetap(arr, n) {
		fmt.Println("Data berjarak", arr[1]-arr[0])
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
