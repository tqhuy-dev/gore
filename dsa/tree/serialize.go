package tree

import (
	"github.com/s-platform/gore/utilities"
	"strings"
)

type ISerializationTree interface {
	SerializeBFS() (result string, err error)
}

type defaultSerialization[T any] struct {
	emptyCharacter string
	node           *NodeTree[T]
}

func NewDefaultSerialization[T any](emptyCharacter string, node *NodeTree[T]) ISerializationTree {
	return &defaultSerialization[T]{
		emptyCharacter: emptyCharacter,
		node:           node,
	}
}

func (ds *defaultSerialization[T]) SerializeBFS() (result string, err error) {
	builder := strings.Builder{}
	ds.node.BFS(func(nodeData *NodeTree[T]) bool {

		if nodeData == nil {
			_, err = builder.Write([]byte(ds.emptyCharacter))
		} else {
			_, err = builder.Write([]byte(utilities.ToString(nodeData.Data)))
		}
		_, err = builder.Write([]byte(","))
		return false
	})
	result = builder.String()
	return
}
