package datastruc

import (
	"errors"
)

type Node struct {
	value    int
	next     *Node
	previous *Node
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

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type BinaryNode struct {
	value int
	left  *BinaryNode
	right *BinaryNode
}

type MinHeap struct {
	length int
	data   []int
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

func (d *DoublyLinkedList) Prepend(data int) {
	new_node := &Node{value: data}

	d.length++
	if d.head == nil {
		d.head = new_node
		d.tail = new_node
		return
	}
	new_node.next = d.head
	d.head.previous = new_node
	d.head = new_node
}

func (d *DoublyLinkedList) InsertAt(data int, index int) error {
	if index > d.length || index < 0 {
		return errors.New("Invalid index")
	}
	if index == d.length {
		d.Append(data)
		return nil
	} else if index == 0 {
		d.Prepend(data)
		return nil
	}
	current := d.head
	for i := 0; i < index; i++ {
		if current != nil {
			current = current.next
		} else {
			return errors.New("Current node is nil")
		}
	}
	new_node := &Node{value: data}
	new_node.next = current
	new_node.previous = current.previous
	if current.previous != nil {
		current.previous.next = new_node
	}
	current.previous = new_node
	d.length++
	return nil
}

func (d *DoublyLinkedList) Append(data int) {
	new_node := &Node{value: data}

	d.length++
	if d.tail == nil {
		d.head, d.tail = new_node, new_node
		return
	}
	new_node.next = d.tail
	d.tail.next = new_node
	d.tail = new_node
}

func (d *DoublyLinkedList) Remove(data int) (int, error) {
	current := d.head
	for i := 0; i < d.length; i++ {
		if current.value == data {
			break
		}
		current = current.next
	}

	if current == nil {
		return -1, errors.New("Error")
	}

	d.length--
	if current.previous != nil {
		current.previous.next = current.next
	}

	if current.next != nil {
		current.next.previous = current.previous
	}

	if current == d.head {
		d.head = current.next
	}

	if current == d.tail {
		d.tail = current.previous
	}

	current.next, current.previous = nil, nil
	return current.value, nil
}

func (h *MinHeap) Insert(value int) {
	if h.length == len(h.data) {
		newSize := 2 * len(h.data)
		if newSize == 0 {
			newSize = 1
		}
		newData := make([]int, newSize)
		copy(newData, h.data)
		h.data = newData
	}

	h.data[h.length] = value
	h.HeapifyUp(h.length)
	h.length++
}

func (h *MinHeap) Delete() int {
	if h.length == 0 {
		return -1
	}

	out := h.data[0]
	h.length--
	h.data[0] = h.data[h.length]
	h.HeapifyDown(0)

	if cap(h.data)/2 > h.length {
		newData := make([]int, h.length, h.length)
		copy(newData, h.data[:h.length])
		h.data = newData
	}
	return out
}

func (h *MinHeap) Parent(index int) int {
	return ((index - 1) / 2)
}

func (h *MinHeap) LeftChild(index int) int {
	return ((index * 2) + 1)
}

func (h *MinHeap) RightChild(index int) int {
	return ((index * 2) + 2)
}

func (h *MinHeap) HeapifyDown(index int) {
	for {
		left := h.LeftChild(index)
		right := h.RightChild(index)
		smallest := index

		if left < h.length && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < h.length && h.data[right] < h.data[smallest] {
			smallest = right
		}

		if smallest != index {
			h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
			index = smallest
		} else {
			break
		}
	}
}

func (h *MinHeap) HeapifyUp(index int) {
	for index > 0 {
		parent := h.Parent(index)
		if h.data[parent] <= h.data[index] {
			break
		}
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

type LRUNode struct {
	value interface{}
	next  *LRUNode
	prev  *LRUNode
}

type LRU struct {
	length        int
	head          *LRUNode
	tail          *LRUNode
	lookup        map[interface{}]*LRUNode
	reverseLookup map[*LRUNode]interface{}
	capacity      int
}

func createNodeLRU(value interface{}) *LRUNode {
	return &LRUNode{value: value}
}

func NewLRU(capacity int) *LRU {
	lru := &LRU{
		length:        0,
		head:          nil,
		tail:          nil,
		lookup:        make(map[interface{}]*LRUNode),
		reverseLookup: make(map[*LRUNode]interface{}),
		capacity:      capacity,
	}
	return lru
}

func (lru *LRU) UpdateLRU(key interface{}, value interface{}) {
	// Check existence
	// If not in cache, insert it:
	// - Check capacity and evict if full
	// Otherwise, update to front of list
	// Update value
	node, exists := lru.lookup[key]
	if !exists {
		node = createNodeLRU(value)
		lru.length++
		lru.prepend(node)
		lru.trimCache()

		lru.lookup[key] = node
		lru.reverseLookup[node] = key
	} else {
		lru.detach(node)
		lru.prepend(node)
		node.value = value
	}

}

func (lru *LRU) GetLRU(key interface{}) interface{} {
	// Check the cache for existence
	node, exists := lru.lookup[key]
	if !exists {
		return nil
	}

	// Update found value and move it to front
	lru.detach(node)
	lru.prepend(node)

	// Return found value or nil if it doesn't exist
	return node.value
}

func (lru *LRU) detach(node *LRUNode) {
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if lru.head == node {
		lru.head = lru.head.next
	}
	if lru.tail == node {
		lru.tail = lru.tail.prev
	}

	node.next, node.prev = nil, nil
}

func (lru *LRU) prepend(node *LRUNode) {
	if lru.head == nil {
		lru.head, lru.tail = node, node
		return
	}

	node.next = lru.head
	lru.head.prev = node
	lru.head = node
}

func (lru *LRU) trimCache() {
	if lru.length <= lru.capacity {
		return
	}

	tmp_tail := lru.tail
	lru.detach(tmp_tail)

	key := lru.reverseLookup[tmp_tail]
	delete(lru.lookup, key)
	delete(lru.reverseLookup, tmp_tail)
	lru.length--
}
