package datastruc

import "testing"

func TestPeekEmptyQueue(t *testing.T) {
	q := &Queue{}
	_, err := q.PeekQueue()
	if err == nil || err.Error() != "Queue is empty" {
		t.Errorf("Expected 'Queue is empty' error, got %v", err)
	}
}

func TestDequeEmptyQueue(t *testing.T) {
	q := &Queue{}
	_, err := q.Deque()
	if err == nil || err.Error() != "Queue is empty" {
		t.Errorf("Expected 'Queue is empty' error, got %v", err)
	}
}

func TestEnqueueAndPeek(t *testing.T) {
	q := &Queue{}
	q.Enqueue(5)
	val, err := q.PeekQueue()
	if err != nil || val != 5 {
		t.Errorf("Expected 5, got %v", val)
	}
}

func TestEnqueueAndDeque(t *testing.T) {
	q := &Queue{}
	q.Enqueue(5)
	val, err := q.Deque()
	if err != nil || val != 5 {
		t.Errorf("Expected 5, got %v", val)
	}
}

func TestMultipleEnqueueAndDeque(t *testing.T) {
	q := &Queue{}
	q.Enqueue(5)
	q.Enqueue(10)
	q.Enqueue(15)

	val_1, _ := q.Deque()
	val_2, _ := q.Deque()
	val_3, _ := q.Deque()

	if val_1 != 5 || val_2 != 10 || val_3 != 15 {
		t.Errorf("Expected 5, 10, 15, got %v, %v, %v", val_1, val_2, val_3)
	}
}

func TestPeekEmptyStack(t *testing.T) {
	s := &Stack{}
	_, err := s.PeekStack()
	if err == nil || err.Error() != "Stack is empty" {
		t.Errorf("Expected 'Stack is empty' error, got %v", err)
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := &Stack{}
	_, err := s.Pop()
	if err == nil || err.Error() != "Stack is empty" {
		t.Errorf("Expected 'Stack is empty' error, got %v", err)
	}
}

func TestPushAndPeek(t *testing.T) {
	s := &Stack{}
	s.Push(5)
	val, err := s.PeekStack()
	if err != nil || val != 5 {
		t.Errorf("Expected 5, got %v", val)
	}
}

func TestPushAndPop(t *testing.T) {
	s := &Stack{}
	s.Push(5)
	val, err := s.Pop()
	if err != nil || val != 5 {
		t.Errorf("Expected 5, got %v", val)
	}
}

func TestMultiplePushAndPop(t *testing.T) {
	s := &Stack{}
	s.Push(5)
	s.Push(10)
	s.Push(15)

	val_1, _ := s.Pop()
	val_2, _ := s.Pop()
	val_3, _ := s.Pop()

	if val_1 != 15 || val_2 != 10 || val_3 != 5 {
		t.Errorf("Expected 15, 10, 5, got %v, %v, %v", val_1, val_2, val_3)
	}
}
