package main

import (
	"dsa_golang/utils"
	"fmt"
)

func binarySearch(arr []int, start int, end int, number int) int {
	res := -1
	for start < end {
		mid := start + (end-start)/2

		if arr[mid] == number {
			res = mid
			start = mid + 1
		} else if arr[mid] > number {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return res
}

func linearSearch(arr []int, start int, end int, number int) int {
	res := -1
	for i := range arr {
		if arr[i] == number {
			res = i
			continue
		}
	}
	return res
}

var res, number, choice int

func main() {
	fmt.Println("Enter same value for more than once")
	arr := utils.ProvideArray()

	fmt.Printf("The array is %v\n", arr)

	fmt.Println("Which number to find")
	fmt.Scan(&number)

	// trying both binary and linear search to find the number
	fmt.Println("Enter which algo to use 1. Linear , 2. Binary")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		res = linearSearch(arr, 0, len(arr)-1, number)
	case 2:
		res = binarySearch(arr, 0, len(arr)-1, number)
	}

	fmt.Printf("The position of first occurence of %d is %d\n", number, res)
}
