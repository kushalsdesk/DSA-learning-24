package main

import (
	"dsa_golang/utils"
	"fmt"
)

var number int
var flag bool

func main() {

	arr := utils.ProvideArray()
	fmt.Printf("The primary arrays is %v\n", arr)
	fmt.Println("The number you want to find: ")
	fmt.Scan(&number)
	start, end := 0, len(arr)-1

	for start <= end {
		mid := start + (end-start)/2

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
