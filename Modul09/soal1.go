package main

import "fmt"

type Titik struct {
	x, y int
}

type Lingkaran struct {
	cx, cy int
	r      int
}

func didalam(c Lingkaran, p Titik) bool {
	dx := p.x - c.cx
	dy := p.y - c.cy

	return dx*dx+dy*dy <= c.r*c.r
}

func main() {
	var c1, c2 Lingkaran
	var p Titik

	fmt.Scan(&c1.cx, &c1.cy, &c1.r)

	fmt.Scan(&c2.cx, &c2.cy, &c2.r)

	fmt.Scan(&p.x, &p.y)

	dalam1 := didalam(c1, p)
	dalam2 := didalam(c2, p)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
