package tree

import (
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
