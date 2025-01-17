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
		fmt.Println(" 1.Insert@First,\n 2.Insert@Last,\n 3.Insert@n,\n 4.Remove@n,\n 5.PrintList,\n 0.Exit")
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
			var pos int
			fmt.Println("Enter the Position: ")
			fmt.Scan(&pos)
			fmt.Println("Enter the value to insert")
			fmt.Scan(&item)
			utils.InsertAt(&ll, pos, item)
		case 4:
			var pos int
			fmt.Println("Enter the Position: ")
			fmt.Scan(&pos)
			fmt.Printf("Removing Node at %d\n", pos)
			utils.RemoveAT(&ll, pos)
		case 5:
			utils.PrintList(&ll)
		case 0:
			return
		default:
			fmt.Println("Enter Valid Option....")
		}
	}
}
