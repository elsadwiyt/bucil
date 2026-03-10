package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var ch1, ch2, ch3, ch4 int

	fmt.Scan(&a, &b, &c, &d, &e)

	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	fmt.Scanf("%c%c%c%c", &ch1, &ch2, &ch3, &ch4)

	fmt.Printf("%c%c%c%c\n", ch1+3, ch2+3, ch3+3, ch4+3)
}
