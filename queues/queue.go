package queues

import "goats/container"

/*
Queue is common interface of queue.

Queue must satisfy Push, Pop, Front, and container.Container(Size, Empty, and Clear).
*/
type Queue[Any any] interface {
	Push(value Any)
	Pop() (value Any, err error)
	Front() (value Any, err error)
	container.Container
}
