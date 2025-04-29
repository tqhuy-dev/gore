package tree

import (
	"fmt"
	"github.com/s-platform/gore/dsa"
)

type HandleNode[T any] func(nodeData T) bool

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
	stackNode := dsa.InitStack[*NodeTree[T]]()
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

func (node *NodeTree[T]) Serialization(serialize ISerializationMethod[T]) (result string, err error) {
	if serialize == nil {
		serialize = NewDefaultBFSSerialization[T](DefaultEmptyCharacter, node)
	}
	return serialize.Serialize()
}

func (node *NodeTree[T]) Deserialization(result string, serialize ISerializationMethod[T]) {
	if serialize == nil {
		serialize = NewDefaultBFSSerialization[T](DefaultEmptyCharacter, nil)
	}
	serialize.Deserialize(result)
	tmp := serialize.GetNode()
	node.Data = tmp.Data
	node.NoteLeft = tmp.NoteLeft
	node.NoteRight = tmp.NoteRight
}

func (node *NodeTree[T]) dfs(root *NodeTree[T]) {
	if root == nil {
		return
	}
	node.dfs(root.NoteLeft)
	node.dfs(root.NoteRight)
}

func (node *NodeTree[T]) BFS(handle HandleNode[*NodeTree[T]]) {
	queue := dsa.InitQueue[*NodeTree[T]]()
	queue.Push(node)
	for queue.Scan() {
		nodeCursor := queue.Pop()
		isStop := handle(nodeCursor)
		if isStop {
			return
		}
		if nodeCursor != nil {
			queue.Push(nodeCursor.NoteLeft)
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

	nodeD := NodeTree[int]{
		Data: 18,
	}

	nodeA.InsertLeft(&NodeTree[int]{
		Data:     13,
		NoteLeft: &nodeD,
	})
	nodeB := NodeTree[int]{
		Data: 7,
	}

	nodeB.InsertLeft(&NodeTree[int]{
		Data: 11,
	})
	nodeC := NodeTree[int]{
		Data: 5,
	}
	nodeC.InsertRight(&nodeA)
	nodeC.InsertLeft(&nodeB)

	result, _ := nodeC.Serialization(nil)
	fmt.Println(result)
}

func ExampleSerialize() {
	ex := "5,7,8,11,#,13,14,#,#,18,#,#,#,#,#"
	nodeA := NodeTree[int64]{}
	nodeA.Deserialization(ex, nil)
	nodeA.BFS(func(nodeData *NodeTree[int64]) bool {
		if nodeData != nil {
			fmt.Println(nodeData.Data)
		}
		return false
	})
}
