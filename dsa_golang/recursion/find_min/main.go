package main

import (
	"dsa_golang/utils"
	"fmt"
)

func main() {
	fmt.Println("Finding Min in the Array")
	arr := utils.ProvideArray()

	min := find_min(arr, 0, arr[0])
	fmt.Printf("The minimum value from the array %v is %d\n", arr, min)
}

func find_min(arr []int, start int, current int) int {
	if start > len(arr)-1 {
		return current
	}

	current = min(current, arr[start])

	return find_min(arr, start+1, current)

}
