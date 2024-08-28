package utils

import "fmt"

var arrSize int

func ProvideArray() []int {
	counter := 0
	fmt.Println("Enter the size of the array: ")
	fmt.Scan(&arrSize)

	a := make([]int, arrSize)
	for counter < arrSize {
		fmt.Printf("Enter the Value for position %d\n", counter)
		fmt.Scan(&a[counter])
		counter++
	}
	return a

}
