package main

import (
	"dsa_golang/utils"
	"fmt"
)

func main() {
	var target, result string

	fmt.Println("Remove all adjacent characters")
	fmt.Println("Enter the string: ")
	fmt.Scan(&target)

	stack := utils.New()

	for _, char := range target {
		top, _ := stack.Peek()
		if top == char {
			stack.Pop()
		} else {
			stack.Push(rune(char))
		}

	}

	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		result = string(item.(rune)) + result
	}

	fmt.Printf("The Finished String is %s\n", result)
}
