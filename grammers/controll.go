package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%d", numbers[i])
	}

	fmt.Println()

	// for-range
	for _, v := range numbers {
		fmt.Printf("%d", v)
	}

	fmt.Println()
}
