package main

import (
	"dsa_golang/utils"
	"fmt"
)

func main() {
	var target int
	fmt.Println("Getting a Stack in FIFO with queue")
	fmt.Println("Enter how many items:")
	fmt.Scan(&target)
	stack := utils.New()
	stack2 := utils.New()

	var item string
	for target != 0 {
		fmt.Println("Enter item to push into ")
		fmt.Scan(&item)
		stack.Push(item)
		target--
	}
	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		stack2.Push(item)
	}

	fmt.Println("Again printing the elements in FIFO manner")
	for !stack2.IsEmpty() {
		item, _ := stack2.Pop()
		fmt.Println(" ", item)
	}

}
