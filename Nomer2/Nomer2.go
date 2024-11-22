package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	// Membuat array dengan kapasitas n
	arr := make([]int, n)

	// Mengisi array
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan keseluruhan isi array")
		fmt.Println("2. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen dengan indeks genap")
		fmt.Println("4. Tampilkan elemen dengan indeks kelipatan x")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata")
		fmt.Println("7. Tampilkan standar deviasi")
		fmt.Println("8. Tampilkan frekuensi dari suatu bilangan")
		fmt.Println("9. Keluar")

		var pilihan int
		fmt.Print("Pilih opsi (1-9): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("Isi array:", arr)

		case 2:
			fmt.Print("Elemen dengan indeks ganjil: ")
			for i := 1; i < len(arr); i += 2 {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()

		case 3:
			fmt.Print("Elemen dengan indeks genap: ")
			for i := 0; i < len(arr); i += 2 {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()

		case 4:
			var x int
			fmt.Print("Masukkan nilai x: ")
			fmt.Scan(&x)
			fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
			for i := 0; i < len(arr); i += x {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()

		case 5:
			var indeks int
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&indeks)
			if indeks >= 0 && indeks < len(arr) {
				arr = append(arr[:indeks], arr[indeks+1:]...)
				fmt.Println("Elemen pada indeks", indeks, "telah dihapus.")
				fmt.Println("Isi array setelah penghapusan:", arr)
			} else {
				fmt.Println("Indeks tidak valid.")
			}

		case 6:
			if len(arr) > 0 {
				var total int
				for _, v := range arr {
					total += v
				}
				rataRata := float64(total) / float64(len(arr))
				fmt.Println("Rata-rata:", rataRata)
			} else {
				fmt.Println("Array kosong.")
			}

		case 7:
			if len(arr) > 0 {
				var total, totalKuadrat float64
				for _, v := range arr {
					total += float64(v)
					totalKuadrat += float64(v * v)
				}
				rataRata := total / float64(len(arr))
				varians := (totalKuadrat / float64(len(arr))) - (rataRata * rataRata)
				stdDev := math.Sqrt(varians)
				fmt.Println("Standar deviasi:", stdDev)
			} else {
				fmt.Println("Array kosong.")
			}

		case 8:
			var bilangan int
			fmt.Print("Masukkan bilangan untuk dihitung frekuensinya: ")
			fmt.Scan(&bilangan)
			frekuensi := 0
			for _, v := range arr {
				if v == bilangan {
					frekuensi++
				}
			}
			fmt.Printf("Frekuensi dari %d: %d\n", bilangan, frekuensi)

		case 9:
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
