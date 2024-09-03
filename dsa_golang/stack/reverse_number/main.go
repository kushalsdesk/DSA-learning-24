package main

import (
	"dsa_golang/utils"
	"fmt"
)

func main() {
	var originalTarget, target int
	result := 0
	stack := utils.New()

	fmt.Println("Reverse a number")
	fmt.Println("Enter a number to reverse")
	fmt.Scan(&target)

	// Save the original value of target
	originalTarget = target

	// Push digits onto the stack
	for target != 0 {
		stack.Push(target % 10)
		target /= 10
	}

	power := 1
	index := 0
	for !stack.IsEmpty() {
		item, _ := stack.Pop()

		// Update power based on the index
		for i := 0; i < index; i++ {
			power *= 10
		}

		// Update the result with the current item
		result = result + (item.(int) * power)
		index++

		// Reset power for the next iteration
		power = 1

		fmt.Printf("result yet %d and index %d after outer loop:\n", result, index)
	}

	fmt.Printf("The reverse of %d is %d:\n", originalTarget, result)
}
