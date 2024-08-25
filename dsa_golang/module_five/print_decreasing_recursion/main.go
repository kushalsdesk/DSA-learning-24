package main

import "fmt"

func main() {
	var target int
	fmt.Println("Printing the rest  in Decreasing order")
	fmt.Println("Enter the target number: ")
	fmt.Scan(&target)

	fmt.Println("Printing all the numbers in decreasing order")
	decreasing(target)
}

func decreasing(target int) {
	if target == 0 {
		return
	}
	fmt.Println(target)
	decreasing(target - 1)
}
