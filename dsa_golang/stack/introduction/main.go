package main

import (
	"dsa_golang/utils"
	"fmt"
)

var choice int

func main() {
	stack := utils.Stack{}
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

		fmt.Scan(&choice)

		switch choice {
		case 1:
			var value int
			fmt.Println("Enter the Value to push : ")
			fmt.Scan(&value)
			stack.Push(value)
			fmt.Printf("%d Pushed \n", value)
		case 2:
			fmt.Println("Popping Element")
			popped := stack.Pop()
			if popped != -1 {

				fmt.Printf("%d popped\n", popped)
			} else {
				fmt.Println("Stack Empty")
			}
		case 3:
			fmt.Println("Peeking to stack")
			peek := stack.Peek()
			if peek != -1 {
				fmt.Printf("%d is topmost\n", peek)
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
