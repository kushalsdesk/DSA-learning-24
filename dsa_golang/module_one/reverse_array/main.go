package main

import "fmt"

func reverseArray(arr []int) []int {

	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--

	}
	return arr
}

func main() {

	array := []int{1, 2, 3, 4, 5, 10, 8, 18, 7, 6}
	fmt.Printf("The array before operation: %d\n", array)
	fmt.Printf("The reversed array is %d\n", reverseArray(array))
}
