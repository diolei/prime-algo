package algos

import (
	"math"
)

type Point struct {
	x int
	y int
}


type BinaryNode struct {
    value int
    left *BinaryNode
    right *BinaryNode
}


func LinearSearch(haystack []int, needle int) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}

func BinarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack) - 1
	for high >= low {
		middle := (low + high) / 2
		if needle == haystack[middle] {
			return true
		} else if needle < haystack[middle] {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return false
}

func twoCrystalBalls(breaks []bool) int {
	jump_amount := math.Floor(math.Sqrt(float64(len(breaks))))

	i := int(jump_amount)
	for ; i < len(breaks); i += int(jump_amount) {
		if breaks[i] {
			break
		}
	}
	i -= int(jump_amount)

	for j := 0; j <= int(jump_amount) && i < len(breaks); i, j = i+1, j+1 {
		if breaks[i] {
			return i
		}
	}
	return -1
}

func bubbleSort(array []int) {
	// keep track of last index
	for i := 0; i < len(array); i++ {
		// -1 to not go out of bounds, - i to discard last element
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j+1], array[j] = array[j], array[j+1]
			}
		}
	}
}

var dir = [][]int{
	{-1, 0}, // Left
	{1, 0},  // Right
	{0, -1}, // Up
	{0, 1},  // Down
}

func mazeSolver(maze []string, wall string, start Point, end Point) []Point {
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}
	path := make([]Point, 0)

	_, path = WalkMaze(maze, wall, start, end, seen, path)
	return path
}

func WalkMaze(maze []string, wall string, current Point, end Point, seen [][]bool, path []Point) (bool, []Point) {
	// 1. Off the map
	if current.x < 0 || current.x >= len(maze[0]) || current.y < 0 || current.y >= len(maze) {
		return false, path
	}
	// 2. On a wall
	if string(maze[current.y][current.x]) == wall {
		return false, path
	}
	// 3. End reached
	if current.x == end.x && current.y == end.y {
		path = append(path, end)
		return true, path
	}
	// 4. Seen
	if seen[current.y][current.x] {
		return false, path
	}

	// Pre
	seen[current.y][current.x] = true
	path = append(path, current)

	// Recurse
	for i := 0; i < len(dir); i++ {
		x, y := dir[i][0], dir[i][1]
		if ok, updated_path := WalkMaze(maze, wall, Point{x: current.x + x, y: current.y + y}, end, seen, path); ok {
			return true, updated_path
		}
	}

	// Post
	path = path[:len(path)-1]

	return false, path
}

// Partition: Fetches the pivot index and moves items to right or left
func partition(array []int, low, high int) ([]int, int) {
	pivot := array[high]
	index := low
	for i := low; i < high; i++ {
		if array[i] < pivot {
			array[index], array[i] = array[i], array[index]
			index++
		}
	}
	array[index], array[high] = array[high], array[index]
	return array, index
}

// Quick Sort: Calls the partition function and recurses
func qs(array []int, low, high int) []int {
	if low < high {
		var pivot_index int
		array, pivot_index = partition(array, low, high)
		array = qs(array, low, pivot_index-1)
		array = qs(array, pivot_index+1, high)
	}
	return array
}

func QuickSort(array []int) {
	qs(array, 0, len(array)-1)
}


func WalkBTPreOrder(current *BinaryNode, path *[]int) {
    if current == nil {
        return
    }
    // Pre order traversal
    *path = append(*path, current.value) // 1. Visit node
    WalkBTPreOrder(current.left, path) // 2. Walk left
    WalkBTPreOrder(current.right, path) // 3. Walk right
}

func PreOrderSearch(head *BinaryNode) []int {
    path := make([]int, 0)
    WalkBTPreOrder(head, &path)
    return path
}

func WalkBTInOrder(current *BinaryNode, path *[]int) {
    if current == nil {
        return
    }

    // In order traversal
    WalkBTInOrder(current.left, path) //  1. Walk left
    *path = append(*path, current.value) // 2. Visit node
    WalkBTInOrder(current.right, path) // 3. Walk right
}

func InOrderSearch(head *BinaryNode) []int {
    path := make([]int, 0)
    WalkBTInOrder(head, &path)
    return path
}
func WalkBTPostOrder(current *BinaryNode, path *[]int) {
    if current == nil {
        return
    }

    // Post order traversal
    WalkBTPostOrder(current.left, path) // 1. Walk left
    WalkBTPostOrder(current.right, path) // 2. Walk right
    *path = append(*path, current.value) // 3. Visit node
}

func PostOrderSearch(head *BinaryNode) []int {
    path := make([]int, 0)
    WalkBTPostOrder(head, &path)
    return path
}
