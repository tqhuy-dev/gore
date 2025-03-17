package algo

import (
	"container/heap"
	"fmt"
	"math"
)

type NodeAStar struct {
	x, y    int
	g, h, f float64
	parent  *NodeAStar
	index   int
}

type PriorityAStarQueue []*NodeAStar

func (pq PriorityAStarQueue) Len() int           { return len(pq) }
func (pq PriorityAStarQueue) Less(i, j int) bool { return pq[i].f < pq[j].f }
func (pq PriorityAStarQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}
func (pq *PriorityAStarQueue) Push(x interface{}) {
	n := x.(*NodeAStar)
	n.index = len(*pq)
	*pq = append(*pq, n)
}
func (pq *PriorityAStarQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// Heuristic function (Manhattan Distance)
func heuristic(a, b NodeAStar) float64 {
	return math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
}

// AStar algorithm implementation
func AStar(grid [][]int, start, goal NodeAStar) []NodeAStar {
	rows, cols := len(grid), len(grid[0])
	directions := []struct{ dx, dy int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	openSet := make(PriorityAStarQueue, 0)
	heap.Init(&openSet)

	start.g, start.h = 0, heuristic(start, goal)
	start.f = start.h
	heap.Push(&openSet, &start)

	visited := make(map[[2]int]bool)
	nodeMap := map[[2]int]*NodeAStar{{start.x, start.y}: &start}

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*NodeAStar)

		// Goal reached
		if current.x == goal.x && current.y == goal.y {
			path := []NodeAStar{}
			for current != nil {
				path = append([]NodeAStar{*current}, path...)
				current = current.parent
			}
			return path
		}

		visited[[2]int{current.x, current.y}] = true

		// Explore neighbors
		for _, d := range directions {
			nx, ny := current.x+d.dx, current.y+d.dy

			if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] == 1 || visited[[2]int{nx, ny}] {
				continue
			}

			newG := current.g + 1
			neighbor, exists := nodeMap[[2]int{nx, ny}]

			if !exists {
				neighbor = &NodeAStar{x: nx, y: ny}
				nodeMap[[2]int{nx, ny}] = neighbor
			}

			if !exists || newG < neighbor.g {
				neighbor.g = newG
				neighbor.h = heuristic(*neighbor, goal)
				neighbor.f = neighbor.g + neighbor.h
				neighbor.parent = current

				if !exists {
					heap.Push(&openSet, neighbor)
				}
			}
		}
	}
	return nil
}

// PrintPath prints the found path
func PrintPath(path []NodeAStar) {
	if path == nil {
		fmt.Println("No path found!")
		return
	}
	fmt.Println("Shortest Path:")
	for _, node := range path {
		fmt.Printf("(%d, %d) -> ", node.x, node.y)
	}
	fmt.Println("Goal")
}

func RunAStar() {
	grid := [][]int{
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}

	start := NodeAStar{x: 0, y: 0}
	goal := NodeAStar{x: 4, y: 4}

	path := AStar(grid, start, goal)
	PrintPath(path)
}
