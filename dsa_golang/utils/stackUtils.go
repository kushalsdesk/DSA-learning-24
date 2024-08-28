package utils

import "fmt"

type Stack struct {
	items []int
}

// NOTE: Push method to add the element
func (s *Stack) Push(value int) {
	fmt.Printf("Pushing %d in stack\n", value)
	s.items = append(s.items, value) // Appending will naturally put it at last
}

// NOTE: Pop method to remove the element
// Reciever is with Pointer as need direct operation on the memory
func (s *Stack) Pop() int {
	fmt.Println("Popping from the Stack")
	if len(s.items) == 0 {
		return -1
	}
	// getting the last element from the items array
	poppedItem := s.items[len(s.items)-1]
	// removing the last element from the items array
	s.items = s.items[:len(s.items)-1]
	return poppedItem
}

// NOTE: Peek() method sends the last element
func (s *Stack) Peek() int {
	fmt.Println("Popping from the Stack")
	if len(s.items) == 0 {
		return -1
	}
	//Sending the element at length-1 index which is the last
	return s.items[len(s.items)-1]
}

// NOTE: IsEmpty() method helps to validate the stack before custom operations
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// NOTE: Size() method will send the Size of stack directly
func (s *Stack) Size() int {
	return len(s.items)
}
