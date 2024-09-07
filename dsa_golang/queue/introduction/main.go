package main

import (
	"dsa_golang/utils"
	"fmt"
)

func main() {
	fmt.Println("Operations with QUEUE")
	queue := utils.NewQueue()

	arr := utils.ProvideArray()

	for i, _ := range arr {
		queue.Enqueue(arr[i])
	}

	fmt.Print("The items are: ")
	for !queue.IsEmpty() {
		item, _ := queue.Dequeue()
		fmt.Printf(" %v\n", item)
	}

}
