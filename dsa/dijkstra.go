package dsa

import (
	"container/heap"
	"fmt"
	"math"
)

type GraphRoute struct {
	vertices map[int]map[int]int
}

type DistanceData struct {
	Dest   int
	Weight int
}

// NewGraph creates a new graph
func NewGraph() *GraphRoute {
	return &GraphRoute{vertices: make(map[int]map[int]int)}
}

func (g *GraphRoute) AddEdge(u, v, weight int) {
	if g.vertices[u] == nil {
		g.vertices[u] = make(map[int]int)
	}
	if g.vertices[v] == nil {
		g.vertices[v] = make(map[int]int)
	}
	g.vertices[u][v] = weight
	g.vertices[v][u] = weight
}

type PriorityQueueItem struct {
	vertex   int
	distance int
}

type PriorityQueue struct {
	Data []*PriorityQueueItem
}

func (pq *PriorityQueue) Len() int           { return len(pq.Data) }
func (pq *PriorityQueue) Less(i, j int) bool { return pq.Data[i].distance < pq.Data[j].distance }
func (pq *PriorityQueue) Swap(i, j int)      { pq.Data[i], pq.Data[j] = pq.Data[j], pq.Data[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*PriorityQueueItem)
	pq.Data = append(pq.Data, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old.Data)
	item := old.Data[n-1]
	pq.Data = old.Data[0 : n-1]
	return item
}

func (g *GraphRoute) Dijkstra(start int) []DistanceData {
	distances := make(map[int]int)
	for vertex := range g.vertices {
		distances[vertex] = math.MaxInt64
	}
	distances[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &PriorityQueueItem{vertex: start, distance: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*PriorityQueueItem)
		currentVertex := current.vertex

		for neighbor, weight := range g.vertices[currentVertex] {
			distance := distances[currentVertex] + weight
			if _, exist := distances[neighbor]; exist == false {
				distances[neighbor] = distance
			} else if distance < distances[neighbor] {
				distances[neighbor] = distance
				heap.Push(pq, &PriorityQueueItem{vertex: neighbor, distance: distance})
			}
		}
	}
	resp := make([]DistanceData, 0)
	for key, value := range distances {
		if value == math.MaxInt64 {
			continue
		}
		resp = append(resp, DistanceData{
			Dest:   key,
			Weight: value,
		})
	}
	return resp
}

func RunGraphRoute() {
	g := NewGraph()
	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 7, 8)
	g.AddEdge(1, 2, 8)
	g.AddEdge(1, 7, 11)
	g.AddEdge(2, 3, 7)
	g.AddEdge(2, 8, 2)
	g.AddEdge(2, 5, 4)
	g.AddEdge(3, 4, 9)
	g.AddEdge(3, 5, 14)
	g.AddEdge(4, 5, 10)
	g.AddEdge(5, 6, 2)
	g.AddEdge(6, 7, 1)
	g.AddEdge(6, 8, 6)
	g.AddEdge(7, 8, 7)

	distances := g.Dijkstra(0)
	for _, element := range distances {
		fmt.Printf("Vertex %d, Distance %d\n", element.Dest, element.Weight)
	}
}
