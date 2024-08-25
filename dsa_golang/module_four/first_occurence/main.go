package main

import (
	"dsa_golang/utils"
	"fmt"
)

var number, res int

func binarySearch(arr []int, start int, end int, number int) int {
	res := -1
	for start < end {

		mid := start + (end-start)/2

		if arr[mid] == number {
			res = mid
			end = mid - 1
		} else if arr[mid] > number {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return res
}
func main() {
	fmt.Println("Enter same value for more than once")
	arr := utils.ProvideArray()

	fmt.Printf("The array is %v\n", arr)

	fmt.Println("Which number to find")
	fmt.Scan(&number)
	//trying only binary search to seek first_occurence

	res = binarySearch(arr, 0, len(arr)-1, number)
	fmt.Printf("The position of first occurence of %d is %d\n", number, res)

}
