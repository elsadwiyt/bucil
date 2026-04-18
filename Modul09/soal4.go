package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

type tab struct {
	tb tabel
	n  int
}

func isiArray(t *tab, n *int) {
	var ch string
	*n = 0

	fmt.Print("Teks : ")
	for {
		fmt.Scan(&ch)

		if ch == "." || *n == NMAX {
			break
		}

		t.tb[*n] = []rune(ch)[0]
		*n++
	}
	t.n = *n
}

func balikanArray(t *tab, n int) {
	for i := 0; i < n/2; i++ {
		t.tb[i], t.tb[n-1-i] = t.tb[n-1-i], t.tb[i]
	}
}

func palindrome(t tab, n int) bool {
	for i := 0; i < n/2; i++ {
		if t.tb[i] != t.tb[n-1-i] {
			return false
		}
	}
	return true
}

func cetak(t tab, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t.tb[i])
	}
	fmt.Println()
}

func main() {
	var x tab
	var n int

	isiArray(&x, &n)

	fmt.Print("Teks : ")
	cetak(x, n)

	balikanArray(&x, n)

	fmt.Print("Reverse teks : ")
	cetak(x, n)

	if palindrome(x, n) {
		fmt.Println("Palindrom ? true")
	} else {
		fmt.Println("Palindrom ? false")
	}
}
