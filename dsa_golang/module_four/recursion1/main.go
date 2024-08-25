package main

import (
	"dsa_golang/utils"
	"fmt"
)

// LinearSearch recursively searches for an element in a sorted array.
func linearSearch(arr []int, start int, end int, number int) int {
	if start > end { // Base case for not found
		return -1
	}
	if arr[start] == number {
		return start
	}
	return linearSearch(arr, start+1, end, number)
}

// BinarySearch recursively searches for an element in a sorted array using binary search.
func binarySearch(arr []int, start int, end int, number int) int {
	if start > end { // Base case for not found
		return -1
	}
	mid := start + (end-start)/2
	if arr[mid] == number {
		return mid
	} else if arr[mid] > number {
		return binarySearch(arr, start, mid-1, number)
	} else {
		return binarySearch(arr, mid+1, end, number)
	}
}

func main() {
	fmt.Println("Sorted array must be present for binary search to work")
	arr := utils.ProvideArray()
	fmt.Printf("The array is %v\n", arr)

	fmt.Println("Enter the value to search: ")
	var number int
	fmt.Scan(&number)

	fmt.Println("Enter which algorithm to use: 1. Linear, 2. Binary")
	var choice int
	fmt.Scan(&choice)

	var result int
	switch choice {
	case 1:
		fmt.Println("Welcome to linear search")
		result = linearSearch(arr, 0, len(arr)-1, number)
	case 2:
		fmt.Println("Welcome to binary search")
		result = binarySearch(arr, 0, len(arr)-1, number)
	default:
		fmt.Println("Invalid choice. Please enter 1 for linear search or 2 for binary search.")
		return
	}

	if result != -1 {
		fmt.Printf("The number %d is at index: %d\n", number, result)
	} else {
		fmt.Println("The number is not present in the array")
	}
}
