package utils

import "fmt"

// create `Stack` structure
type Stack struct {
	items []interface{}
}

// Creates an Struc and returns a pointer to it
func New() *Stack {
	return &Stack{
		items: []interface{}{},
	}
}

// NOTE: Push method to add the element
func (s *Stack) Push(item interface{}) { // interface{} says that incoming parameter can be any type
	fmt.Printf("Pushing %v in stack\n", item)
	s.items = append(s.items, item) // Appending will naturally put it at last
}

// NOTE: Pop method to remove the element & Reciever is with Pointer as need direct operation on the memory
func (s *Stack) Pop() (interface{}, bool) {
	fmt.Println("Popping from the Stack")
	if len(s.items) == 0 {
		return nil, true
	}
	// getting the last element from the items array
	poppedItem := s.items[len(s.items)-1]

	// removing the last element from the items array
	s.items = s.items[:len(s.items)-1]
	return poppedItem, false
}

// NOTE: Peek() method sends the last element
func (s *Stack) Peek() (interface{}, bool) {
	fmt.Println("Peeking from the Stack")
	if len(s.items) == 0 {
		return nil, true
	}
	//Sending the element at length-1 index which is the last
	return s.items[len(s.items)-1], false
}

// NOTE: IsEmpty() method helps to validate the stack before custom operations
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// NOTE: Size() method will send the Size of stack directly
func (s *Stack) Size() int {
	return len(s.items)
}
