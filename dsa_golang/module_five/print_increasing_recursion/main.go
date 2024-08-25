package main

import "fmt"

var target, start int = 0, 1

func main() {
	fmt.Println("Enter the target number : ")
	fmt.Scan(&target)

	fmt.Println("Printing the numbers in increasing order")

	increasing(start, target)
}

func increasing(start, target int) {
	if start > target {
		return
	}
	fmt.Println(start)
	increasing(start+1, target)
}
