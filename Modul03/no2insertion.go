package main

import "fmt"

const NMAX = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [NMAX]Buku

func inputBuku(A *DaftarBuku, n *int) {
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Scan(
			&A[i].id,
			&A[i].judul,
			&A[i].penulis,
			&A[i].penerbit,
			&A[i].eksemplar,
			&A[i].tahun,
			&A[i].rating,
		)
	}
}

func sortBuku(A *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := A[i]
		j := i - 1

		for j >= 0 && A[j].rating < temp.rating {
			A[j+1] = A[j]
			j--
		}

		A[j+1] = temp
	}
}

func main() {
	var A DaftarBuku
	var n, cari int

	inputBuku(&A, &n)

	sortBuku(&A, n)

	fmt.Println("Buku Terfavorit:")
	fmt.Println(A[0].judul, A[0].penulis, A[0].penerbit, A[0].tahun)

	fmt.Println("5 Buku Rating Tertinggi:")
	for i := 0; i < n && i < 5; i++ {
		fmt.Println(A[i].judul, "-", A[i].rating)
	}

	fmt.Scan(&cari)

	kiri, kanan := 0, n-1
	ketemu := false

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if A[tengah].rating == cari {
			fmt.Println("Data Buku:")
			fmt.Println(
				A[tengah].judul,
				A[tengah].penulis,
				A[tengah].penerbit,
				A[tengah].tahun,
				A[tengah].eksemplar,
				A[tengah].rating,
			)
			ketemu = true
			break
		} else if cari > A[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}
