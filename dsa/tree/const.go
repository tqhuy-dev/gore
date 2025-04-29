package tree

type SearchType int

const (
	BFS                   SearchType = 1
	DFS                   SearchType = 2
	DefaultEmptyCharacter            = "#"
	DefaultSplitCharacter            = ","
)
