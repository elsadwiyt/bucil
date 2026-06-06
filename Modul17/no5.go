package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	var dalamPizza int

	fmt.Print("Banyak Topping: ")
	fmt.Scan(&n)

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= n; i++ {

		x := rand.Float64()
		y := rand.Float64()

		if (x-0.5)*(x-0.5)+(y-0.5)*(y-0.5) <= 0.25 {
			dalamPizza++
		}
	}

	pi := 4.0 * float64(dalamPizza) / float64(n)

	fmt.Println("Topping pada Pizza:", dalamPizza)
	fmt.Printf("PI : %.10f\n", pi)
}
