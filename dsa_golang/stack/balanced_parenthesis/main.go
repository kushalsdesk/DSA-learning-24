// it checks if a series of parenthesis are balanced or not
// ((()))  balanced
// ((())  not-balanced
package main

import (
	"dsa_golang/utils"
	"fmt"
)

var target string

func main() {
	fmt.Println("Balanced Parenthesis")
	fmt.Println("Enter the series: ")
	fmt.Scan(&target)

	stack := utils.New()

	for _, char := range target {
		top, _ := stack.Peek()
		if top == '(' && char == ')' {
			stack.Pop()
		} else {
			stack.Push(rune(char))
		}
	}

	//as per logic if a series is balanced then the stack should be empty
	if stack.IsEmpty() {
		fmt.Printf("The series %s was balanced\n", target)
	} else {
		fmt.Printf("The series %s was not-balanced\n", target)
	}
}
