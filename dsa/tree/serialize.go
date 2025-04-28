package tree

import (
	"github.com/s-platform/gore/utilities"
	"strings"
)

type ISerializationMethod interface {
	Serialize() (result string, err error)
	Deserialize(result string)
}

type defaultSerialization[T any] struct {
	emptyCharacter string
	node           *NodeTree[T]
}

func (ds *defaultSerialization[T]) Deserialize(result string) {
	//TODO implement me
	panic("implement me")
}

func NewDefaultBFSSerialization[T any](emptyCharacter string, node *NodeTree[T]) ISerializationMethod {
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
		_, err = builder.Write([]byte(","))
		return false
	})
	result = builder.String()
	return
}
