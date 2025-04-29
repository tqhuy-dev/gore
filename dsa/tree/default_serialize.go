package tree

import (
	"github.com/s-platform/gore/dsa"
	"github.com/s-platform/gore/utilities"
	"strings"
)

type defaultSerialization[T any] struct {
	emptyCharacter string
	node           *NodeTree[T]
}

func (ds *defaultSerialization[T]) GetNode() *NodeTree[T] {
	return ds.node
}

func (ds *defaultSerialization[T]) Deserialize(result string) {
	arr := strings.Split(result, DefaultSplitCharacter)
	if len(arr) == 0 {
		return
	}
	queue := dsa.InitQueue[*NodeTree[T]]()
	ds.node = &NodeTree[T]{
		Data: utilities.StringParse[T](arr[0]),
	}
	queue.Push(ds.node)
	for index := 1; index < len(arr); index += 2 {
		cursor := queue.Pop()
		if arr[index] != ds.emptyCharacter {
			cursor.NoteLeft = &NodeTree[T]{
				Data: utilities.StringParse[T](arr[index]),
			}
			queue.Push(cursor.NoteLeft)
		}
		if arr[index+1] != ds.emptyCharacter {
			cursor.NoteRight = &NodeTree[T]{
				Data: utilities.StringParse[T](arr[index+1]),
			}
			queue.Push(cursor.NoteRight)
		}
	}
}

func NewDefaultBFSSerialization[T any](emptyCharacter string, node *NodeTree[T]) ISerializationMethod[T] {
	return &defaultSerialization[T]{
		emptyCharacter: emptyCharacter,
		node:           node,
	}
}

func (ds *defaultSerialization[T]) Serialize() (result string, err error) {
	builder := strings.Builder{}
	ds.node.BFS(func(nodeData *NodeTree[T]) bool {

		if nodeData == nil {
			_, err = builder.Write([]byte(ds.emptyCharacter))
		} else {
			_, err = builder.Write([]byte(utilities.ToString(nodeData.Data)))
		}
		_, err = builder.Write([]byte(DefaultSplitCharacter))
		return false
	})
	result = builder.String()
	result = result[0 : len(result)-1]
	return
}
