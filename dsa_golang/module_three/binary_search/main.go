package main

import "fmt"

var arrSize int
var number int
var flag bool

func provideArray() []int {
	counter := 0
	fmt.Println("Enter the size of the array: ")
	fmt.Scan(&arrSize)

	fmt.Println("Remember the arrays needs to sorted for binary search. so push values in order")
	a := make([]int, arrSize)
	for counter < arrSize {
		fmt.Printf("Enter the Value for position %d\n", counter)
		fmt.Scan(&a[counter])
		counter++
	}
	return a
}
func main() {

	arr := provideArray()
	fmt.Printf("The primary arrays is %v\n", arr)
	fmt.Println("The number you want to find: ")
	fmt.Scan(&number)
	start, end := 0, len(arr)-1

	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == number {
			fmt.Printf("The number %d is located at index %d\n", number, mid)
			flag = true
			break
		} else if arr[mid] < number {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	if !flag {
		fmt.Println("The number is not present in the array")
	}
}
