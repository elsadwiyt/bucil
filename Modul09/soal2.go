package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	var arr [100]int

	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\na. Seluruh isi array:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("\n\nb. Elemen dengan indeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("\n\nc. Elemen dengan indeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}

	var x int
	fmt.Print("\n\nMasukkan nilai x untuk indeks kelipatan x: ")
	fmt.Scan(&x)

	fmt.Println("d. Elemen pada indeks kelipatan", x, ":")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	var idx int
	fmt.Print("\n\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)

	for i := idx; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Println("e. Array setelah penghapusan:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	var jumlah int
	for i := 0; i < n; i++ {
		jumlah += arr[i]
	}
	rata := float64(jumlah) / float64(n)
	fmt.Println("\n\nf. Rata-rata =", rata)

	var selisih, total float64
	for i := 0; i < n; i++ {
		selisih = float64(arr[i]) - rata
		total += selisih * selisih
	}
	std := total / float64(n)
	fmt.Println("g. Variansi =", std)

	var cari, frek int
	fmt.Print("Masukkan bilangan yang dicari frekuensinya: ")
	fmt.Scan(&cari)

	for i := 0; i < n; i++ {
		if arr[i] == cari {
			frek++
		}
	}

	fmt.Println("h. Frekuensi", cari, "=", frek)
}
