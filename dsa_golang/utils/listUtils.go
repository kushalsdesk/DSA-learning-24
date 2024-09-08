package utils

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

func Newlist() *Node {
	return nil
}

func InsertFirst(head **Node, value interface{}) {
	fmt.Println("Inserting Value at First")
	newNode := &Node{Value: value}
	newNode.Next = *head // Point to the current head
	*head = newNode      // Set new node as the new head
}

func InsertLast(head **Node, value interface{}) {
	fmt.Println("Inserting Value at last")
	newNode := &Node{Value: value}
	if *head == nil {
		*head = newNode
		return
	}
	current := *head
	for current.Next != nil {
		current = current.Next // Traverse to the end of the list
	}
	current.Next = newNode // Append the new node at the end
}

func PrintList(head **Node) {
	if *head == nil {
		fmt.Println("The list is already empty")
		return
	}
	current := *head
	for current != nil {
		fmt.Print(current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}
