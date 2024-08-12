package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Printf("The maxtrix is %d\n", matrix)
	fmt.Println("Traversing the Matrix for every elements")

	for i, row := range matrix {
		for j, val := range row {
			fmt.Printf("The element at position [%d][%d] is %d\n", i, j, val)
		}
	}
}
