package container

/*
Container is common interface of data structures.

Container must satisfy the three methods Size, Empty, and Clear.

Size returns number of elements stored in the data structure.

Empty returns wheather the data structure is empty or not.

Clear empties the data structure.
*/
type Container interface {
	Size() int
	Empty() bool
	Clear()
}
