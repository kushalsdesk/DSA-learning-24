package main

import (
	"dsa_golang/utils"
	"fmt"
)

var number int

func main() {
	a := utils.ProvideArray()
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
