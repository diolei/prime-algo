package algos

import (
	"reflect"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	// Test case where the needle is present in the haystack
	haystack := []int{1, 2, 3, 4, 5}
	needle := 3
	result := LinearSearch(haystack, needle)
	if result != true {
		t.Errorf("LinearSearch(%v, %v) = %v; want true", haystack, needle, result)
	}

	// Test case where the needle is not present in the haystack
	haystack = []int{1, 2, 4, 5}
	needle = 3
	result = LinearSearch(haystack, needle)
	if result != false {
		t.Errorf("LinearSearch(%v, %v) = %v; want false", haystack, needle, result)
	}

	// Test case where the haystack is empty
	haystack = []int{}
	needle = 3
	result = LinearSearch(haystack, needle)
	if result != false {
		t.Errorf("LinearSearch(%v, %v) = %v; want false", haystack, needle, result)
	}
}

func TestBinarySearch(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5}
	needle := 3
	result := BinarySearch(haystack, needle)
	if result != true {
		t.Errorf("BinarySearch(%v, %v) = %v; want true", haystack, needle, result)
	}

	haystack = []int{1, 2, 4, 5}
	needle = 3
	result = BinarySearch(haystack, needle)
	if result != false {
		t.Errorf("BinarySearch(%v, %v) = %v; want false", haystack, needle, result)
	}

	haystack = []int{}
	needle = 3
	result = BinarySearch(haystack, needle)
	if result != false {
		t.Errorf("BinarySearch(%v, %v) = %v; want false", haystack, needle, result)
	}
}

func TestTwoCrystalBalls(t *testing.T) {
	// Test case where the ball breaks at some point
	breaks := []bool{false, false, false, true, true, true, true, true}
	result := twoCrystalBalls(breaks)
	expected := 3
	if result != expected {
		t.Errorf("twoCrystalBalls(%v) = %v; want %v", breaks, result, expected)
	}

	// Test case where the ball doesn't break
	breaks = []bool{false, false, false, false, false, false, false}
	result = twoCrystalBalls(breaks)
	expected = -1
	if result != expected {
		t.Errorf("twoCrystalBalls(%v) = %v; want %v", breaks, result, expected)
	}

	// Test case where the ball breaks at the last drop
	breaks = []bool{false, false, false, false, false, false, true}
	result = twoCrystalBalls(breaks)
	expected = 6
	if result != expected {
		t.Errorf("twoCrystalBalls(%v) = %v; want %v", breaks, result, expected)
	}
}

func TestBubbleSort(t *testing.T) {
	// Test case with unsorted array
	input := []int{3, 2, 1}
	bubbleSort(input)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("BubbleSort(%v) = %v; want %v", input, input, expected)
	}

	// Test case with partially sorted array
	input = []int{5, 1, 4, 2, 8}
	bubbleSort(input)
	expected = []int{1, 2, 4, 5, 8}
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("BubbleSort(%v) = %v; want %v", input, input, expected)
	}

	// Test case with already sorted array
	input = []int{1, 2, 3, 4, 5}
	bubbleSort(input)
	expected = []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("BubbleSort(%v) = %v; want %v", input, input, expected)
	}

	// Test case with an empty array
	input = []int{}
	bubbleSort(input)
	expected = []int{}
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("BubbleSort(%v) = %v; want %v", input, input, expected)
	}
}

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	expected := []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}

	start := Point{x: 10, y: 0}
	end := Point{x: 1, y: 5}
	wall := "x"
	result := mazeSolver(maze, wall, start, end)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed. Expected %v, got %v", expected, result)
	}
}

func TestQuickSort(t *testing.T) {
	// Test case with unsorted array
	input := []int{3, 2, 1}
	QuickSort(input)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("QuickSort(%v) = %v; want %v", input, input, expected)
	}

	// Test case with partially sorted array
	input = []int{5, 1, 4, 2, 8}
	QuickSort(input)
	expected = []int{1, 2, 4, 5, 8}
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("QuickSort(%v) = %v; want %v", input, input, expected)
	}

	// Test case with already sorted array
	input = []int{1, 2, 3, 4, 5}
	QuickSort(input)
	expected = []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("QuickSort(%v) = %v; want %v", input, input, expected)
	}

	// Test case with an empty array
	input = []int{}
	QuickSort(input)
	expected = []int{}
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("QuickSort(%v) = %v; want %v", input, input, expected)
	}
}

func TestTraversals(t *testing.T) {
	root := &BinaryNode{value: 1}
	root.left = &BinaryNode{value: 2}
	root.right = &BinaryNode{value: 3}
	root.left.left = &BinaryNode{value: 4}
	root.left.right = &BinaryNode{value: 5}
	root.right.left = &BinaryNode{value: 6}
	root.right.right = &BinaryNode{value: 7}

	// Expected results for pre-order traversal
	expectedPreOrder := []int{1, 2, 4, 5, 3, 6, 7}
	// Expected results for in-order traversal
	expectedInOrder := []int{4, 2, 5, 1, 6, 3, 7}
	// Expected results for post-order traversal
	expectedPostOrder := []int{4, 5, 2, 6, 7, 3, 1}

	if !reflect.DeepEqual(PreOrderSearch(root), expectedPreOrder) {
		t.Errorf("PreOrderSearch failed")
	}
	if !reflect.DeepEqual(InOrderSearch(root), expectedInOrder) {
		t.Errorf("InOrderSearch failed")
	}
	if !reflect.DeepEqual(PostOrderSearch(root), expectedPostOrder) {
		t.Errorf("PostOrderSearch failed")
	}
}

func TestBFS(t *testing.T) {
	// Test case where the needle is present in the haystack
	t.Run("NeedlePresent", func(t *testing.T) {
		root := &BinaryNode{value: 5}
		root.left = &BinaryNode{value: 3}
		root.right = &BinaryNode{value: 7}
		root.left.left = &BinaryNode{value: 2}
		root.left.right = &BinaryNode{value: 4}
		root.right.left = &BinaryNode{value: 6}
		root.right.right = &BinaryNode{value: 8}

		result := BFS(root, 6)
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})

	// Test case where the needle is not present in the haystack
	t.Run("NeedleNotPresent", func(t *testing.T) {
		root := &BinaryNode{value: 5}
		root.left = &BinaryNode{value: 3}
		root.right = &BinaryNode{value: 7}
		root.left.left = &BinaryNode{value: 2}
		root.left.right = &BinaryNode{value: 4}
		root.right.left = &BinaryNode{value: 6}
		root.right.right = &BinaryNode{value: 8}

		result := BFS(root, 9)
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})

	// Test case where the haystack is empty
	t.Run("EmptyTree", func(t *testing.T) {
		root := &BinaryNode{}

		result := BFS(root, 1)
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})
}

func TestDFS(t *testing.T) {
	// Test case where the needle is present in the haystack
	t.Run("NeedlePresent", func(t *testing.T) {
		root := &BinaryNode{value: 5}
		root.left = &BinaryNode{value: 3}
		root.right = &BinaryNode{value: 7}
		root.left.left = &BinaryNode{value: 2}
		root.left.right = &BinaryNode{value: 4}
		root.right.left = &BinaryNode{value: 6}
		root.right.right = &BinaryNode{value: 8}

		result := DFS(root, 6)
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})

	// Test case where the needle is not present in the haystack
	t.Run("NeedleNotPresent", func(t *testing.T) {
		root := &BinaryNode{value: 5}
		root.left = &BinaryNode{value: 3}
		root.right = &BinaryNode{value: 7}
		root.left.left = &BinaryNode{value: 2}
		root.left.right = &BinaryNode{value: 4}
		root.right.left = &BinaryNode{value: 6}
		root.right.right = &BinaryNode{value: 8}

		result := DFS(root, 9)
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})

	// Test case where the haystack is empty
	t.Run("EmptyTree", func(t *testing.T) {
		root := &BinaryNode{}

		result := DFS(root, 1)
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})
}

func TestMatrixBFS(t *testing.T) {
	tests := []struct {
		name         string
		graph        WeightedAdjacencyMatrix
		source       int
		needle       int
		expectedPath []int
	}{
		{
			name: "Direct Path",
			graph: WeightedAdjacencyMatrix{
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
				{0, 0, 0, 0},
			},
			source:       0,
			needle:       2,
			expectedPath: []int{0, 1, 2},
		},
		{
			name: "No Path",
			graph: WeightedAdjacencyMatrix{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			source:       0,
			needle:       2,
			expectedPath: nil,
		},
		{
			name: "Self Loop",
			graph: WeightedAdjacencyMatrix{
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			source:       0,
			needle:       0,
			expectedPath: []int{0},
		},
		{
			name: "Indirect Path",
			graph: WeightedAdjacencyMatrix{
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
				{0, 0, 0, 0},
			},
			source:       0,
			needle:       3,
			expectedPath: []int{0, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := MatrixBFS(&tt.graph, tt.source, tt.needle)
			if !equalIntSlices(path, tt.expectedPath) {
				t.Errorf("MatrixBFS = %v, expected %v", path, tt.expectedPath)
			}
		})
	}
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMatrixDFS(t *testing.T) {
	graph := WeightedAdjacencyList{
		{{to: 1, weight: 1}, {to: 2, weight: 2}},
		{{to: 3, weight: 3}},
		{{to: 4, weight: 4}},
		{{to: 5, weight: 5}},
		{{to: 6, weight: 6}},
		{{to: 7, weight: 7}},
		{},
	}

	source := 0
	needle := 5

	path := MatrixDFS(&graph, source, needle)
	if path == nil {
		t.Errorf("Expected a non-nil path, but got nil")
	} else {
		t.Logf("Found path: %v", path)
	}
}
