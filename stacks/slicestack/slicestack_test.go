package slicestack

import (
	"errors"
	"testing"
)

func TestStackPushPop(t *testing.T) {
	stack := New[int]()
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if determinedValue := stack.Empty(); determinedValue != true {
		t.Errorf("Got %v expected %v", determinedValue, true)
	}

	for i, v := range list {
		stack.Push(v)
		determinedValue, err := stack.Pop()
		if determinedValue != list[i] || err != nil {
			t.Errorf("Got %v, %v expected %v, %v", determinedValue, err, list[i], nil)
		}
		if determinedValue = determinedValue + 1; determinedValue != list[i]+1 || err != nil {
			t.Errorf("Got %v, %v expected %v, %v", determinedValue, err, list[i]+1, nil)
		}
	}

	if determinedValue := stack.Empty(); determinedValue != true {
		t.Errorf("Got %v expected %v", determinedValue, true)
	}

	for _, v := range list {
		stack.Push(v)
	}

	for i := range list {
		determinedValue, err := stack.Pop()
		if determinedValue != list[len(list)-1-i] || err != nil {
			t.Errorf("Got %v, %v expected %v, %v", determinedValue, err, list[len(list)-1-i], nil)
		}
		if determinedValue = determinedValue + 1; determinedValue != list[len(list)-1-i]+1 || err != nil {
			t.Errorf("Got %v, %v expected %v, %v", determinedValue, err, list[len(list)-1-i]+1, nil)
		}
	}

	if determinedValue := stack.Empty(); determinedValue != true {
		t.Errorf("Got %v expected %v", determinedValue, true)
	}
}

func TestStackSize(t *testing.T) {
	stack := New[int]()
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if determinedValue := stack.Empty(); determinedValue != true {
		t.Errorf("Got %v expected %v", determinedValue, true)
	}

	for _, v := range list {
		stack.Push(v)
	}

	if determinedValue := stack.Size(); determinedValue != len(list) {
		t.Errorf("Got %v expected %v", determinedValue, len(list))
	}
}

func TestStackClear(t *testing.T) {
	stack := New[int]()
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if determinedValue := stack.Empty(); determinedValue != true {
		t.Errorf("Got %v expected %v", determinedValue, true)
	}

	for _, v := range list {
		stack.Push(v)
	}

	if determinedValue := stack.Empty(); determinedValue != false {
		t.Errorf("Got %v expected %v", determinedValue, true)
	}

	stack.Clear()

	if determinedValue := stack.Empty(); determinedValue != true {
		t.Errorf("Got %v expected %v", determinedValue, true)
	}
}

func TestStackTop(t *testing.T) {
	stack := New[int]()

	stack.Push(1)

	if determinedValue, err := stack.Top(); determinedValue != 1 || err != nil {
		t.Errorf("Got %v, %v expected %v, %v", determinedValue, err, 1, nil)
	}

	if determinedValue, err := stack.Pop(); determinedValue != 1 || err != nil {
		t.Errorf("Got %v, %v expected %v, %v", determinedValue, err, 1, nil)
	}

	if determinedValue := stack.Empty(); determinedValue != true {
		t.Errorf("Got %v expected %v", determinedValue, true)
	}

	if determinedValue, err := stack.Top(); determinedValue != 0 || err == nil {
		t.Errorf("Got %v, %v expected %v, %v", determinedValue, err, 0, errors.New("stack is empty"))
	}

	if determinedValue, err := stack.Pop(); determinedValue != 0 || err == nil {
		t.Errorf("Got %v, %v expected %v, %v", determinedValue, err, 0, errors.New("stack is empty"))
	}

}
