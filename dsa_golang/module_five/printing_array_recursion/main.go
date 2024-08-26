// will include both the forward and backward printing in the same file
package main

import (
	"dsa_golang/utils"
	"fmt"
)

var choice int

// Print Forward
func forwardPrint(arr []int, counter int) {
	if counter > len(arr)-1 {
		return
	}
	fmt.Println(arr[counter])

	forwardPrint(arr, counter+1)
}

// Print Backward
func backwardprint(arr []int, size int) {
	if size < 0 {
		return
	}
	fmt.Println(arr[size])
	backwardprint(arr, size-1)
}

func main() {
	fmt.Println(" Both the forward and backward array printing")
	arr := utils.ProvideArray()
	fmt.Println("Choose Printing style: 1. forward, 2. backward")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("Printing Forward as usual")
		forwardPrint(arr, 0)
	case 2:
		fmt.Println("Printing Backward as usual")
		backwardprint(arr, len(arr)-1)
	default:
		fmt.Println("Enter valid options to execute")
	}

}
