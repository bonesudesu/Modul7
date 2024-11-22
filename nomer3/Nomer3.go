package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string

	// Input nama klub
	fmt.Print("Masukkan nama klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Masukkan nama klub B: ")
	fmt.Scanln(&klubB)

	// Input skor pertandingan
	for {
		fmt.Print("Masukkan skor klub A: ")
		fmt.Scanln(&skorA)
		fmt.Print("Masukkan skor klub B: ")
		fmt.Scanln(&skorB)

		// Validasi skor
		if skorA < 0 || skorB < 0 {
			break
		}

		// Menentukan pemenang
		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		}
	}

	// Menampilkan daftar pemenang
	fmt.Println("\nDaftar Klub Pemenang:")
	if len(pemenang) == 0 {
		fmt.Println("Tidak ada klub yang menang.")
	} else {
		for _, p := range pemenang {
			fmt.Println(p)
		}
	}
}
