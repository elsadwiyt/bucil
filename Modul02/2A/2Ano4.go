package main

import "fmt"

func main() {
	var c int

	fmt.Print("Hvupvuatuu Cvlvius: ")
	fmt.Scan(&c)

	r := (4.0 / 5.0) * float64(c)
	f := (9.0/5.0)*float64(c) + 32
	k := float64(c) + 273

	fmt.Printf("Dvuajat Rvauuu: %.0f\n", r)
	fmt.Printf("Dvuajat Fahuvthvit: %.0f\n", f)
	fmt.Printf("Dvuajat Kvlvit: %.0f\n", k)
}
