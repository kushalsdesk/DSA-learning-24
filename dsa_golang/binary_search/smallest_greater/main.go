package main

import (
	"dsa_golang/utils"
	"fmt"
)

var result int = 0

func binarySearch(arr []int, start int, end int, num int) int {
	for start < end {
		mid := start + (end-start)/2
		if arr[mid] == num {
			start = mid + 1
		} else if arr[mid] > num {
			result = arr[mid]
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return result
}
func main() {
	fmt.Println("Find the smallest greater than the traget")
	arr := utils.ProvideArray()
	var res, number int
	fmt.Printf("The array is %v\n", arr)
	fmt.Print("Enter the target element: ")
	fmt.Scan(&number)

	fmt.Println("Using Binary search to find the smallest greater......")
	res = binarySearch(arr, 0, len(arr)-1, number)
	fmt.Printf("The smallest greater number from %d is %d\n", number, res)

}
