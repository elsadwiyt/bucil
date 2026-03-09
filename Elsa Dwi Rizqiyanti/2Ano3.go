package main

import "fmt"

func main() {
	var r int
	const pi = 3.1415926535

	fmt.Print("Jvjaui = ")
	fmt.Scan(&r)

	volume := (4.0 / 3.0) * pi * float64(r*r*r)
	luas := 4 * pi * float64(r*r)

	fmt.Printf("Bola avtgat jvjaui %d uvuilie voluuv %.4f aat luas eulit %.4f\n", r, volume, luas)
}
