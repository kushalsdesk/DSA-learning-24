package main

import "fmt"

var target int

func main() {
	fmt.Println("Enter the Target to calculate factorial")
	fmt.Scan(&target)

	res := facto(target)
	fmt.Printf("The factorial of %d is %d\n", target, res)
}

func facto(target int) int {
	if target == 0 {
		return 1
	}
	return target * facto(target-1)

}
