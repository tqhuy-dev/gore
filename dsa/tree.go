package dsa

import "fmt"

type TreeSearchType int

const (
	BFS TreeSearchType = 1
	DFS TreeSearchType = 2
)

type HandleNode[T any] func(nodeData T) error

type NodeTree[T any] struct {
	Data      T
	NoteLeft  *NodeTree[T]
	NoteRight *NodeTree[T]
}

func (node *NodeTree[T]) InsertLeft(nodeAdd *NodeTree[T]) {
	node.NoteLeft = nodeAdd
}

func (node *NodeTree[T]) InsertRight(nodeAdd *NodeTree[T]) {
	node.NoteRight = nodeAdd
}

func (node *NodeTree[T]) DFSWithStack() {
	stackNode := InitStack[*NodeTree[T]]()
	stackNode.Push(node)
	for stackNode.Scan() {
		cursorNode := stackNode.Pop()
		if cursorNode != nil {
			fmt.Println(cursorNode.Data)
			if cursorNode.NoteRight != nil {
				stackNode.Push(cursorNode.NoteRight)
			}
			if cursorNode.NoteLeft != nil {
				stackNode.Push(cursorNode.NoteLeft)
			}
		}
	}
}

func (node *NodeTree[T]) DFSRecursion() {
	node.dfs(node)
}

//func (node *NodeTree[T]) GenerateStringTree(data []T, typeof TreeSearchType) string {
//
//}

func (node *NodeTree[T]) dfs(root *NodeTree[T]) {
	if root == nil {
		return
	}
	node.dfs(root.NoteLeft)
	node.dfs(root.NoteRight)
}

func (node *NodeTree[T]) BFS() {
	queue := InitQueue[*NodeTree[T]]()
	queue.Push(node)
	for queue.Scan() {
		nodeCursor := queue.Pop()
		fmt.Println(nodeCursor.Data)
		if nodeCursor.NoteLeft != nil {
			queue.Push(nodeCursor.NoteLeft)
		}
		if nodeCursor.NoteRight != nil {
			queue.Push(nodeCursor.NoteRight)
		}
	}
}

func ExampleTree() {
	nodeA := NodeTree[int]{
		Data: 8,
	}
	nodeA.InsertRight(&NodeTree[int]{
		Data: 14,
	})
	nodeA.InsertLeft(&NodeTree[int]{
		Data: 13,
	})
	nodeB := NodeTree[int]{
		Data: 7,
	}
	nodeB.InsertRight(&NodeTree[int]{
		Data: 12,
	})
	nodeB.InsertLeft(&NodeTree[int]{
		Data: 11,
	})
	nodeC := NodeTree[int]{
		Data: 5,
	}
	nodeC.InsertRight(&nodeA)
	nodeC.InsertLeft(&nodeB)

	nodeC.BFS()
}
