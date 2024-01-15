package datastruc

import "errors"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

type Stack struct {
	head   *Node
	length int
}

func (q *Queue) PeekQueue() (int, error) {
	if q.head == nil {
		return 0, errors.New("Queue is empty")
	}
	return q.head.value, nil
}

func (q *Queue) Deque() (int, error) {
	if q.head == nil {
		return 0, errors.New("Queue is empty")
	}

	q.length--

	current := q.head
	q.head = q.head.next
	return current.value, nil
}

func (q *Queue) Enqueue(data int) {
	q.length++
	new_node := &Node{value: data}
	if q.head == nil {
		q.head = new_node
		q.tail = new_node
	} else {
		q.tail.next = new_node
		q.tail = new_node
	}

}

func (s *Stack) PeekStack() (int, error) {
	if s.head == nil {
		return 0, errors.New("Stack is empty")
	}
	return s.head.value, nil
}

func (s *Stack) Push(data int) {
	s.length++
	new_node := &Node{value: data}
	if s.head == nil {
		s.head = new_node
	}

	new_node.next = s.head
	s.head = new_node
}

func (s *Stack) Pop() (int, error) {
	if s.head == nil {
		return 0, errors.New("Stack is empty")
	}
	s.length--

	current := s.head
	s.head = current.next
	return current.value, nil
}
