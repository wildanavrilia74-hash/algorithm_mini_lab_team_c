package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{5, 2, 8, 1, 4}

	fmt.Println("Before:", arr)

	bubbleSort(arr)

	fmt.Println("After:", arr)
}
