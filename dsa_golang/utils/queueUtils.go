package utils

import "fmt"

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		items: []interface{}{},
	}
}

func (q *Queue) Enqueue(item interface{}) {
	fmt.Printf("Enqueuing %v in queue\n", item)
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (interface{}, bool) {
	fmt.Println("Dequeuing from Queue")
	if len(q.items) == 0 {
		return nil, true
	}
	removedItem := q.items[0] //grabbing the first element
	q.items = q.items[1:]     //building without the first element
	return removedItem, false
}

func (q *Queue) Peek() (interface{}, bool) {
	fmt.Println("Peeking in the Queue")
	if len(q.items) == 0 {
		return nil, true
	}
	return q.items[0], false

}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
func (q *Queue) Size() int {
	return len(q.items)
}
