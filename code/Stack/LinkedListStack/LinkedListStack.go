package LinkedListStack

import (
	"code/List"
	"code/List/LinkedList"
	"code/Stack"
)

type LinkedListStack struct {
	linkedList *LinkedList.LinkedList
	capacity   uint
}

func NewLinkedListStack(capacity uint) *LinkedListStack {
	stack := new(LinkedListStack)
	stack.linkedList = LinkedList.NewLinkedList()
	if capacity >= Stack.INIT_CAPACITY {
		stack.capacity = capacity
	} else {
		stack.capacity = Stack.INIT_CAPACITY
	}
	return stack
}

func (stack *LinkedListStack) Size() uint {
	return stack.linkedList.Size()
}

func (stack *LinkedListStack) Capacity() uint {
	return stack.capacity
}

func (stack *LinkedListStack) IsFull() bool {
	return stack.Size() >= stack.capacity
}

func (stack *LinkedListStack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *LinkedListStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	list := stack.linkedList
	res, _ := list.Get(list.Size() - 1)
	_ = list.Delete(list.Size() - 1)
	return res
}

func (stack *LinkedListStack) Push(value interface{}) {
	if !stack.IsFull() {
		stack.linkedList.Append(value)
	}
}

func (stack *LinkedListStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	res, _ := stack.linkedList.Get(stack.linkedList.Size() - 1)
	return res
}

func (stack *LinkedListStack) Iterator() List.Iterator {
	return stack.linkedList.Iterator()
}
