package tree

import (
	"github.com/s-platform/gore/dsa"
	"github.com/s-platform/gore/utilities"
	"strings"
)

type dfsSerialize[T any] struct {
	emptyCharacter string
	node           *NodeTree[T]
}

func (d *dfsSerialize[T]) Serialize() (result string, err error) {
	arr := make([]string, 0)

	d.node.DFSRecursion(func(nodeData *NodeTree[T]) bool {
		if nodeData == nil {
			arr = append(arr, DefaultEmptyCharacter)
		} else {
			arr = append(arr, utilities.ToString(nodeData.Data))
		}
		return false
	})

	result = strings.Join(arr, DefaultSplitCharacter)
	return
}

func (d *dfsSerialize[T]) Deserialize(result string) {
	arr := strings.Split(result, DefaultSplitCharacter)
	if len(arr) == 0 {
		return
	}
	d.node = &NodeTree[T]{
		Data: utilities.StringParse[T](arr[0]),
	}
	type tmp struct {
		node    *NodeTree[T]
		isLeft  bool
		isRight bool
	}
	stack := dsa.InitStack[tmp]()
	stack.Push(tmp{
		node:    d.node,
		isLeft:  false,
		isRight: false,
	})
	for index := 1; index < len(arr); index++ {
		cursor := stack.Pop()
		var nextElement *NodeTree[T]
		rePush := true

		if arr[index] != DefaultEmptyCharacter {
			nextElement = &NodeTree[T]{
				Data: utilities.StringParse[T](arr[index]),
			}
		}

		if cursor.isLeft == false {
			cursor.node.NoteLeft = nextElement
			cursor.isLeft = true
		} else if cursor.isRight == false {
			cursor.node.NoteRight = nextElement
			cursor.isRight = true
		}

		if cursor.isRight && cursor.isLeft {
			rePush = false
		}
		if rePush {
			stack.Push(cursor)
		}
		if nextElement != nil {
			stack.Push(tmp{
				node:    nextElement,
				isLeft:  false,
				isRight: false,
			})
		}
	}
}

func (d *dfsSerialize[T]) GetNode() *NodeTree[T] {
	return d.node
}

func NewDFSSerialize[T any](emptyCharacter string,
	node *NodeTree[T]) ISerializationMethod[T] {
	return &dfsSerialize[T]{
		emptyCharacter: emptyCharacter,
		node:           node,
	}
}
