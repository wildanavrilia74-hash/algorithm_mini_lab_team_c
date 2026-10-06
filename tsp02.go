package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan jumlah kota: ")
	fmt.Scan(&n)

	// Matriks jarak antar kota
	distance := make([][]int, n)

	for i := 0; i < n; i++ {
		distance[i] = make([]int, n)
	}

	// Input jarak
	fmt.Println("\nMasukkan jarak antar kota:")

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				distance[i][j] = 0
			} else if i < j {
				fmt.Printf(
					"Jarak kota %d ke kota %d: ",
					i+1,
					j+1,
				)

				fmt.Scan(&distance[i][j])

				// Jarak dibuat dua arah
				distance[j][i] = distance[i][j]
			}
		}
	}

	// Menampilkan matriks
	fmt.Println("\nMatriks Jarak:")

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%5d", distance[i][j])
		}
		fmt.Println()
	}

	// Array untuk menandai kota yang sudah dikunjungi
	visited := make([]bool, n)

	// Mulai dari kota pertama
	current := 0
	visited[current] = true

	// Menyimpan rute
	route := []int{current}

	totalDistance := 0

	// Mencari kota terdekat
	for len(route) < n {

		nextCity := -1
		minDistance := int(^uint(0) >> 1)

		for i := 0; i < n; i++ {

			if !visited[i] &&
				distance[current][i] < minDistance {

				minDistance = distance[current][i]
				nextCity = i
			}
		}

		// Jika tidak ada kota yang bisa dikunjungi
		if nextCity == -1 {
			break
		}

		// Tandai kota sebagai sudah dikunjungi
		visited[nextCity] = true

		// Tambahkan jarak
		totalDistance +=
			distance[current][nextCity]

		// Tambahkan kota ke rute
		route = append(route, nextCity)

		// Pindah ke kota berikutnya
		current = nextCity
	}

	// Kembali ke kota awal
	totalDistance += distance[current][0]

	route = append(route, 0)

	// Menampilkan hasil
	fmt.Println("\n================================")
	fmt.Println("HASIL TSP")
	fmt.Println("================================")

	fmt.Print("Rute: ")

	for i, city := range route {

		fmt.Printf("Kota %d", city+1)

		if i < len(route)-1 {
			fmt.Print(" -> ")
		}
	}

	fmt.Println()

	fmt.Println(
		"Total jarak:",
		totalDistance,
	)
}
