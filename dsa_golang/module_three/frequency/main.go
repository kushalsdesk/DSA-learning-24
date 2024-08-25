package main

import (
	"dsa_golang/utils"
	"fmt"
)

// Helper function for binary search
func binarySearch(arr []int, target int, findFirst bool) int {
	low, high := 0, len(arr)-1
	result := -1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			result = mid
			if findFirst {
				high = mid - 1 // Continue searching in the left half for the first occurrence
			} else {
				low = mid + 1 // Continue searching in the right half for the last occurrence
			}
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}

// Function to find the frequency of a target number
func findFrequency(arr []int, target int) int {
	first := binarySearch(arr, target, true)
	if first == -1 {
		return 0 // Target not found
	}
	last := binarySearch(arr, target, false)
	if last == -1 {
		return 0 // Target not found, should not happen if first is valid
	}
	return last - first + 1
}
func linearSearch(arr []int, start int, end int, number int) int {
	if start > end {
		return 0
	}
	if arr[start] == number {
		return 1 + linearSearch(arr, start+1, end, number)
	}

	return linearSearch(arr, start+1, end, number)

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
		result = findFrequency(arr, number)
	default:
		fmt.Println("Invalid choice. Please enter 1 for linear search or 2 for binary search.")
		return
	}

	if result != -1 {
		fmt.Printf("The number %d is present %d times\n", number, result)
	} else {
		fmt.Println("The number is not present in the array")
	}

}
