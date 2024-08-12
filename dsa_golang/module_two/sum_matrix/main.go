package main

import "fmt"

func main() {

	mat_one := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	mat_two := [][]int{
		{7, 8, 9},
		{4, 5, 6},
		{1, 2, 3},
	}

	mat_total := make([][]int, len(mat_one))
	for i := range mat_total {
		mat_total[i] = make([]int, len(mat_one[i]))
	}

	fmt.Printf("The Two matrices for addition are %v ,%v\n", mat_one, mat_two)

	// as the matrices are of the same size it helps to add and program the addition

	for i, row := range mat_one {
		for j := range row {
			mat_total[i][j] = mat_one[i][j] + mat_two[i][j]
		}
	}

	fmt.Printf("The Total of the added matrices %v\n", mat_total)
}
