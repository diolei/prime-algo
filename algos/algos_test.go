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
