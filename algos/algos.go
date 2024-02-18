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
	left  *BinaryNode
	right *BinaryNode
}

type WeightedAdjacencyMatrix [][]int

type GraphEdge struct {
	to     int
	weight int
}

type WeightedAdjacencyList [][]GraphEdge

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
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
	WalkBTPreOrder(current.left, path)   // 2. Walk left
	WalkBTPreOrder(current.right, path)  // 3. Walk right
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
	WalkBTInOrder(current.left, path)    //  1. Walk left
	*path = append(*path, current.value) // 2. Visit node
	WalkBTInOrder(current.right, path)   // 3. Walk right
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
	WalkBTPostOrder(current.left, path)  // 1. Walk left
	WalkBTPostOrder(current.right, path) // 2. Walk right
	*path = append(*path, current.value) // 3. Visit node
}

func PostOrderSearch(head *BinaryNode) []int {
	path := make([]int, 0)
	WalkBTPostOrder(head, &path)
	return path
}

func BFS(head *BinaryNode, needle int) bool {
	var q []*BinaryNode = []*BinaryNode{head}

	for len(q) > 0 {
		current := q[0]
		q = q[1:] // Dequeue the front node

		if current.value == needle {
			return true
		}

		// Enqueue children nodes
		if current.left != nil {
			q = append(q, current.left)
		}
		if current.right != nil {
			q = append(q, current.right)
		}
	}
	return false
}

func CompareBT(a *BinaryNode, b *BinaryNode) bool {
	// Structural check
	if a == nil && b == nil {
		return true
	}

	// Structural check
	if a == nil || b == nil {
		return false
	}

	// Value check
	if a.value != b.value {
		return false
	}

	return CompareBT(a.left, b.left) && CompareBT(a.right, b.right)
}

func SearchDFS(current *BinaryNode, needle int) bool {
	if current == nil {
		return false
	} else if current.value == needle {
		return true
	} else if current.value < needle {
		return SearchDFS(current.right, needle)
	} else {
		return SearchDFS(current.left, needle)
	}
}

func DFS(head *BinaryNode, needle int) bool {
	return SearchDFS(head, needle)
}

func MatrixBFS(graph *WeightedAdjacencyMatrix, source int, needle int) []int {
	if source == needle {
		return []int{source}
	}

	seen := make([]bool, len(*graph))
	previous := make([]int, len(*graph))
	for i := range previous {
		previous[i] = -1
	}

	seen[source] = true
	q := make([]int, 0)
	q = append(q, source)

	for len(q) != 0 {
		current := q[0]
		q = q[1:]

		if current == needle {
			break
		}

		adj := (*graph)[current]
		for i := 0; i < len(adj); i++ {
			if adj[i] == 0 || seen[i] {
				continue
			}
			seen[i] = true
			previous[i] = current
			q = append(q, i)
		}
	}

	if previous[needle] == -1 {
		return nil
	}

	out := make([]int, 0)
	current := needle
	for current != -1 {
		out = append([]int{current}, out...)
		current = previous[current]
	}

	return out
}

func WalkMatrixDFS(graph *WeightedAdjacencyList, current int, needle int, seen []bool, path *[]int) bool {

	if seen[current] {
		return false
	}

	seen[current] = true

	// Pre
	*path = append(*path, current)

	if current == needle {
		return true
	}

	// Recurse
	list := (*graph)[current]
	for i := 0; i < len(list); i++ {
		edge := list[i]
		if WalkMatrixDFS(graph, edge.to, needle, seen, path) {
			return true
		}
	}

	// Post
	*path = (*path)[:len(*path)-1]

	return false
}

func MatrixDFS(graph *WeightedAdjacencyList, source int, needle int) []int {
	seen := make([]bool, len(*graph))
	path := new([]int)

	WalkMatrixDFS(graph, source, needle, seen, path)

	if len(*path) == 0 {
		return nil
	}

	return *path
}

func hasUnvisited(seen []bool, distances []int) bool {
	for i, s := range seen {
		if !s && distances[i] < int(math.Inf(1)) {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, distances []int) int {
	index := -1
	lowest_distance := int(math.Inf(1))

	for i := 0; i < len(seen); i++ {
		if seen[i] == true {
			continue
		}

		if lowest_distance > distances[i] {
			lowest_distance = distances[i]
			index = i
		}

	}
	return index
}

func Dijkstra(source int, sink int, array *WeightedAdjacencyList) []int {
	seen := make([]bool, len(*array))
	distances := make([]int, len(*array))
	previous := make([]int, len(*array))
	for i := range distances {
		distances[i] = int(math.Inf(1))
	}

	for i := range previous {
		distances[i] = -1
	}

	distances[source] = 0

	for hasUnvisited(seen, distances) {
		current := getLowestUnvisited(seen, distances)
		seen[current] = true

		adj := (*array)[current]
		for i := 0; i < len(adj); i++ {
			edge := adj[i]
			if seen[edge.to] {
				continue
			}
			distance := distances[current] + edge.weight
			if distance < distances[edge.to] {
				distances[edge.to] = distance
				previous[edge.to] = current
			}
		}

	}
	out := make([]int, 0)
	current := sink
	for current != -1 {
		out = append([]int{current}, out...)
		current = previous[current]
	}
	return out
}
