package slicequeue

import (
	"errors"
	"goats/queues"
)

type SliceQueue[Any any] []Any

/*
New method returns SliceQueue (Defined type of []Any).

SliceQueue is implemented using Generics,
it is necessary to specify type when using this method.
*/
func New[Any any]() queues.Queue[Any] {
	queue := make(SliceQueue[Any], 0)
	return &queue
}

/*
Pop method returns front of value of queue and nil if there are any elements on the queue,
and removes front element from the queue.

However, if queue is empty, it returns initial value of the type and error.
*/
func (queue *SliceQueue[Any]) Push(value Any) {
	(*queue) = append((*queue), value)
}

func (queue *SliceQueue[Any]) Pop() (value Any, err error) {
	front, err := queue.Front()

	if err == nil {
		(*queue)[0] = *new(Any)
		(*queue) = (*queue)[1:]
	}

	return front, err
}

/*
front method returns front of value of queue and nil if there are any elements on the queue.

However, if queue is empty, it returns initial value of the type and error.
*/
func (queue *SliceQueue[Any]) Front() (value Any, err error) {
	if queue.Empty() {
		return *new(Any), errors.New("queue is empty")
	} else {
		return (*queue)[0], nil
	}
}

func (queue *SliceQueue[Any]) Size() int {
	return len(*queue)
}

func (queue *SliceQueue[Any]) Empty() bool {
	return queue.Size() == 0
}

func (queue *SliceQueue[Any]) Clear() {
	(*queue) = make(SliceQueue[Any], 0)
}

func (queue *SliceQueue[Any]) Rotate(n int) {
	if !queue.Empty() {
		rotateIndex := n % queue.Size()

		if n < 0 {
			rotateIndex += queue.Size()
		}

		(*queue) = append((*queue)[rotateIndex:], (*queue)[:rotateIndex]...)
	}
}
