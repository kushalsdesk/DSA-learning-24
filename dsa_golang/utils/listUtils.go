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

func InsertAt(head **Node, pos int, value interface{}) {
	fmt.Printf("Inserting at %v\n", pos)
	if pos == 1 {
		InsertFirst(head, value)
		return
	}
	if *head == nil {
		fmt.Println("The list is empty")
		return
	}
	current := *head
	newNode := &Node{Value: value}
	loop := 1

	for loop < pos-1 && current.Next != nil {
		current = current.Next
		loop++
	}

	if loop == pos-1 {
		newNode.Next = current.Next
		current.Next = newNode
	} else {
		fmt.Println("Position out of bounds")
	}
}
func RemoveAT(head **Node, pos int) {
	if *head == nil {
		fmt.Println("Already empty")
		return
	}
	if pos == 1 {
		*head = (*head).Next
		return
	}
	current := *head
	loop := 1
	for loop < pos-1 && current.Next != nil {
		current = current.Next
		loop++
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	} else {
		fmt.Println("Index Out of bounds")
	}
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
