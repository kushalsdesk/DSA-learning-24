package main

import (
	"dsa_golang/utils"
	"fmt"
)

var choice int

func main() {
	fmt.Println("Working with Different Operations with LinkedList")
	ll := utils.Newlist()
	for true {
		fmt.Println(" 1.Insert@First,\n 2.Insert@Last,\n 3.PrintList,\n 5.Exit")
		fmt.Println("Enter Option:")
		fmt.Scan(&choice)
		var item string
		switch choice {
		case 1:
			fmt.Println("Enter the value to insert")
			fmt.Scan(&item)
			utils.InsertFirst(&ll, item)
		case 2:
			fmt.Println("Enter the value to insert")
			fmt.Scan(&item)
			utils.InsertLast(&ll, item)
		case 3:
			utils.PrintList(&ll)
		case 5:
			return
		default:
			fmt.Println("Enter Valid Option....")
		}
	}
}
