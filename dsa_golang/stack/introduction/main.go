package main

import (
	"dsa_golang/utils"
	"fmt"
	"strconv"
)

var input string

func main() {
	stack := utils.New()
	fmt.Println("Welcome to Stack Operations")
	for {
		fmt.Println("\nStack Operations Menu:")
		fmt.Println("1. Push")
		fmt.Println("2. Pop")
		fmt.Println("3. Peek")
		fmt.Println("4. Check if Empty")
		fmt.Println("5. Display Size")
		fmt.Println("6. Exit")
		fmt.Printf("Choose an option: ")
		fmt.Println()

		fmt.Scanln(&input)
		// Try to convert the input to an integer
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 6.")
			continue // Skip to the next iteration of the loop
		}
		switch choice {
		case 1:
			var value string
			fmt.Println("Enter the Value to push(any type) : ")
			fmt.Scan(&value)
			stack.Push(value)
			fmt.Printf("%v Pushed \n", value)
		case 2:
			fmt.Println("Popping Element")
			popped, empty := stack.Pop()
			if !empty {

				fmt.Printf("%v popped\n", popped)
			} else {
				fmt.Println("Stack Empty")
			}
		case 3:
			fmt.Println("Peeking to stack")
			peek, empty := stack.Peek()
			if !empty {
				fmt.Printf("%v is topmost\n", peek)
			} else {
				fmt.Println("Stack already empty")
			}
		case 4:
			if stack.IsEmpty() {
				fmt.Println("Stack is Empty")
			} else {
				fmt.Println("Stack not Empty")
			}
		case 5:
			fmt.Printf("The stack size is: %d\n", stack.Size())
		case 6:
			fmt.Println("exiting...")
			return
		default:
			fmt.Println("Enter Valid Option")
			continue

		}
	}
}
