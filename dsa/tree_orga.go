package dsa

import (
	"fmt"
	"github.com/tqhuy-dev/gore/utilities"
)

type Graph interface {
	BFS(start int) []TreeNode
	DFS(start int) []TreeNode
}

type TreeNode struct {
	Data  int
	Level int
	Root  int
}

type dfsData struct {
	VisitedArray []TreeNode
	mapLevel     map[int]int
}

type TreeOrg struct {
	vertices map[int][]int
}

func NewTreeOrg() *TreeOrg {
	return &TreeOrg{vertices: make(map[int][]int)}
}

func (g *TreeOrg) AddEdge(v, w int) {
	g.vertices[v] = append(g.vertices[v], w)
}

func GetNodeChildrenWithRoot(data []TreeNode, value int) []TreeNode {
	result := make([]TreeNode, 0)

	queue := make([]int, 0)
	queue = append(queue, value)
	for len(queue) > 0 {
		v := queue[0]
		nodes := utilities.Filter(data, func(item TreeNode, index int) bool {
			return item.Root == v && item.Data != item.Root
		})
		result = append(result, nodes...)
		for _, node := range nodes {
			queue = append(queue, node.Data)
		}
		queue = queue[1:]

	}

	return result
}

type Node struct {
	Data     int
	Children []*Node
}

func BuildTree(entries []TreeNode) *Node {
	nodes := make(map[int]*Node)

	// Create nodes for all entries
	for _, e := range entries {
		nodes[e.Data] = &Node{Data: e.Data}
	}

	var root *Node

	// Link nodes based on the root field
	for _, e := range entries {
		if e.Root == 0 {
			root = nodes[e.Data] // Identify root node
		} else {
			parent := nodes[e.Root]
			parent.Children = append(parent.Children, nodes[e.Data])
		}
	}

	return root
}

func PrintTree(node *Node, prefix string, isLast bool) {
	if node == nil {
		return
	}

	// Print the current node
	fmt.Println(prefix + "|-" + fmt.Sprint(node.Data))

	// New prefix for children
	newPrefix := prefix
	if isLast {
		newPrefix += "   "
	} else {
		newPrefix += "|  "
	}

	// Print all children
	for i, child := range node.Children {
		PrintTree(child, newPrefix, i == len(node.Children)-1)
	}
}

func PrintOrganization(data []TreeNode, value int) {
	currentNode, exist := utilities.Find(data, func(item TreeNode) bool {
		return item.Data == value
	})
	if !exist {
		return
	}
	nodes := GetNodeChildrenWithRoot(data, currentNode.Data)
	nodes = append([]TreeNode{currentNode}, nodes...)
	tree := BuildTree(nodes)
	PrintTree(tree, "", true)
}

func (g *TreeOrg) BFS(start int) []TreeNode {
	visited := make(map[int]bool)
	queue := []int{start}
	mapLevel := make(map[int]int)
	mapLevel[start] = 0
	visited[start] = true
	visitedArray := make([]TreeNode, 0)
	visitedArray = append(visitedArray, TreeNode{
		Data:  start,
		Level: 0,
		Root:  0,
	})

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		for _, adjacent := range g.vertices[vertex] {
			mapLevel[adjacent] = mapLevel[vertex] + 1
			visitedArray = append(visitedArray, TreeNode{
				Data:  adjacent,
				Level: mapLevel[adjacent],
				Root:  vertex,
			})
			if !visited[adjacent] {
				visited[adjacent] = true
				queue = append(queue, adjacent)
			}
		}
	}
	return visitedArray
}

// DFSUtil is a utility function used by DFS
func (g *TreeOrg) DFSUtil(v int, visited map[int]bool, data *dfsData) {
	// Mark the current node as visited
	visited[v] = true

	// Recur for all the vertices adjacent to this vertex
	for _, adjacent := range g.vertices[v] {
		data.mapLevel[adjacent] = data.mapLevel[v] + 1
		data.VisitedArray = append(data.VisitedArray, TreeNode{
			Data:  adjacent,
			Level: data.mapLevel[adjacent],
			Root:  v,
		})
		if !visited[adjacent] {
			g.DFSUtil(adjacent, visited, data)
		}
	}
}

// DFS performs the Depth-First Search
func (g *TreeOrg) DFS(start int) []TreeNode {
	// Mark all the vertices as not visited
	visited := make(map[int]bool)

	tmp := dfsData{
		VisitedArray: make([]TreeNode, 0),
		mapLevel:     make(map[int]int),
	}
	tmp.VisitedArray = append(tmp.VisitedArray, TreeNode{
		Data:  start,
		Level: 0,
		Root:  0,
	})
	tmp.mapLevel[start] = 0
	// Call the recursive helper function to print DFS traversal
	g.DFSUtil(start, visited, &tmp)
	return tmp.VisitedArray
}

func RunTreeOrg() {
	g := NewTreeOrg()
	g.AddEdge(0, 3)
	g.AddEdge(3, 27)
	g.AddEdge(3, 23)
	g.AddEdge(0, 4)
	g.AddEdge(4, 1)
	g.AddEdge(4, 7)
	g.AddEdge(4, 15)
	g.AddEdge(15, 2)
	g.AddEdge(15, 17)
	g.AddEdge(27, 127)
	g.AddEdge(27, 125)

	fmt.Println("Breadth-First Search starting from vertex 2:")
	nodeData := g.DFS(0)
	PrintOrganization(nodeData, 3)
}
