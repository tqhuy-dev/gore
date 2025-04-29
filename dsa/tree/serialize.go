package tree

type ISerializationMethod[T any] interface {
	Serialize() (result string, err error)
	Deserialize(result string)
	GetNode() *NodeTree[T]
}
