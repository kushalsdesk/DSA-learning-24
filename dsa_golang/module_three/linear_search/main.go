package main

import "fmt"

var number int
var arrSize int

func provideArray() []int {
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

func main() {
	a := provideArray()
	fmt.Printf("The Primary Slice is %v\n", a)
	fmt.Println("Enter the Number to find from the array")
	fmt.Scan(&number)

	for i, val := range a {
		if val == number {
			fmt.Printf("The number %d is present at index %d\n", val, i)
			return
		}
	}
}
