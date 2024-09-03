package main

import (
	"dsa_golang/utils"
	"fmt"
)

// var target, result string

func main() {
	target, result := "", ""
	fmt.Println("Reverse a String")
	fmt.Println("Enter a String to Reverse: ")
	fmt.Scan(&target)
	stack := utils.New()
	for _, char := range target {
		stack.Push(char)
	}
	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		if char, ok := item.(rune); ok {
			result += string(char)
		}
	}

	fmt.Printf("The Reversed String is: %s\n", result)
}
