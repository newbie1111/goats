package stacks

import "goats/container"

/*
Stack is common interface of stack.

Stack must satisfy Push, Pop, Top, and container.Container(Size, Empty, and Clear).
*/
type Stack[Any any] interface {
	Push(value Any)
	Pop() (value Any, err error)
	Top() (value Any, err error)
	container.Container
}
