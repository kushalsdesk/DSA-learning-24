package main

import (
	"dsa_golang/utils"
	"fmt"
)

func main() {
	fmt.Println(" Reversing a Queue with stack ")
	arr := utils.ProvideArray()
	queue := utils.NewQueue()
	stack := utils.New()
	for i, _ := range arr {
		queue.Enqueue(arr[i])
	}

	for !queue.IsEmpty() {
		item, _ := queue.Dequeue()
		stack.Push(item)

	}
	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		queue.Enqueue(item)
	}

	for !queue.IsEmpty() {
		fmt.Println("The reversed queue is ")
		item, _ := queue.Dequeue()
		fmt.Printf(" %v\n", item)
	}
}
