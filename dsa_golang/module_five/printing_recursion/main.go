package main

import "fmt"

var target, start int = 0, 1
var choice int

func decreasing(target int) {
	if target == 0 {
		return
	}
	fmt.Println(target)
	decreasing(target - 1)
}
func increasing(start, target int) {
	if start > target {
		return
	}
	fmt.Println(start)
	increasing(start+1, target)
}
func main() {
	fmt.Println("Enter the target number : ")
	fmt.Scan(&target)
	fmt.Println("Enter options to print: 1. Forward, 2. Backward")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("Printing the numbers in increasing order")
		increasing(start, target)
	case 2:
		fmt.Println("Printing all the numbers in decreasing order")
		decreasing(target)
	default:
		fmt.Println("Enter valid options to Execute")
	}

}
