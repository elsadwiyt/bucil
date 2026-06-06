package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var banyakTopping int
	fmt.Print("Banyak Topping: ")
	fmt.Scan(&banyakTopping)

	toppingPadaPizza := 0
	xc, yc := 0.5, 0.5
	rKuadrat := 0.25

	for i := 0; i < banyakTopping; i++ {
		x := rand.Float64()
		y := rand.Float64()
		dx := x - xc
		dy := y - yc
		jarakKuadrat := (dx * dx) + (dy * dy)

		if jarakKuadrat <= rKuadrat {
			toppingPadaPizza++
		}
	}

	pi := 4.0 * float64(toppingPadaPizza) / float64(banyakTopping)

	fmt.Printf("Topping pada Pizza: %d\n", toppingPadaPizza)
	fmt.Printf("PI : %.10f\n", pi)
}
