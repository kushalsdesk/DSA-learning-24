package main

import (
	"fmt"
)

func main() {

	matrix := [][]int{
		{7, 8, 9},
		{4, 5, 6},
		{1, 2, 3},
	}

	fmt.Printf("The matrix is %v\n", matrix)

	//The principal Diagonal happens when the row_value == col_value
	fmt.Println("The Prinicipal Diagonal axis is: ")

	for i, row := range matrix {
		for j, val := range row {
			if i == j {
				fmt.Println(val)
			}
		}
	}

	// Secondary Diagonal happens when the row_index + col_index == matrix_size-1
	// so for simplication we can have col_index  = (matrix_size-1) - row_index
	fmt.Println("The Secondary Diagonal axis is: ")

	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i][len(matrix)-1-i])
	}

}
