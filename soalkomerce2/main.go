package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input the number of families : ")
	fmt.Scan(&n)

	// Membuat slice untuk menyimpan jumlah anggota tiap keluarga
	families := make([]int, n)

	fmt.Print("Input the number of members in the family ( separated by a space) : ")

	// Menghitung berapa input yang berhasil dibaca
	inputCount := 0
	for i := 0; i < n; i++ {
		count, err := fmt.Scan(&families[i])
		if err != nil {
			break
		}
		inputCount += count
	}

	// Cek apakah jumlah input sesuai dengan jumlah keluarga
	if inputCount != n {
		fmt.Println("Input must be equal with count of family")
		return
	}

	// Hitung minimum bus yang dibutuhkan
	buses := 0
	used := make([]bool, n)

	// Proses setiap keluarga
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}

		used[i] = true
		remaining := 4 - families[i]

		// Coba pasangkan dengan keluarga lain jika memungkinkan
		if remaining >= 0 {
			paired := false
			for j := i + 1; j < n && !paired; j++ {
				if !used[j] && families[j] <= remaining {
					used[j] = true
					paired = true
				}
			}
		}

		buses++
	}

	fmt.Printf("Minimum bus required is : %d\n", buses)
}
