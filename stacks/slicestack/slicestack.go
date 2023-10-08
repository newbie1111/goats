package slicestack

import (
	"errors"
	"goats/stacks"
)

type SliceStack[Any any] []Any

/*
New method returns SliceStack (Defined type of []Any).

SliceStack is implemented using Generics,
it is necessary to specify type when using this method.
*/
func New[Any any]() stacks.Stack[Any] {
	stack := make(SliceStack[Any], 0)
	return &stack
}

func (stack *SliceStack[Any]) Push(value Any) { *stack = append(*stack, value) }

/*
Pop method returns top of value of stack and nil if there are any elements on the stack,
and removes top element from the stack.

However, if stack is empty, it returns initial value of the type and error.
*/
func (stack *SliceStack[Any]) Pop() (value Any, err error) {
	top, err := stack.Top()

	if err == nil {
		(*stack)[stack.Size()-1] = *new(Any)
		(*stack) = (*stack)[:stack.Size()-1]
	}

	return top, err
}

/*
Top method returns top of value of stack and nil if there are any elements on the stack.

However, if stack is empty, it returns initial value of the type and error.
*/
func (stack *SliceStack[Any]) Top() (value Any, err error) {
	if stack.Empty() {
		return *new(Any), errors.New("stack is empty")
	} else {
		return (*stack)[stack.Size()-1], nil
	}
}

func (stack *SliceStack[Any]) Size() int { return len((*stack)) }

func (stack *SliceStack[Any]) Empty() bool { return stack.Size() == 0 }

func (stack *SliceStack[Any]) Clear() { (*stack) = make([]Any, 0) }
