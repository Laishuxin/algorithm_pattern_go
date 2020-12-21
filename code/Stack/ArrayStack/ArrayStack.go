package ArrayStack

import "code/Stack"

type ArrayStack struct {
	dataStore []interface{}
	capacity  uint
	size      uint
}


func NewArrayStack(capacity uint) *ArrayStack {
	stack := new(ArrayStack)
	if capacity >= Stack.INIT_CAPACITY {
		stack.capacity = capacity
	} else {
		stack.capacity = Stack.INIT_CAPACITY
	}
	stack.size = 0
	stack.dataStore = make([]interface{}, 0, stack.capacity)

	return stack
}

func (stack *ArrayStack) Size() uint {
	return stack.size
}

func (stack *ArrayStack) Capacity() uint {
	return stack.capacity
}

func (stack *ArrayStack) IsFull() bool {
	return stack.size >= stack.capacity
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *ArrayStack) Pop() interface{} {
	if (stack.IsEmpty()) {
		return nil
	}
	stack.size--
	return stack.dataStore[stack.size]
}

func (stack *ArrayStack) Push(value interface{}) {
	if (!stack.IsFull()) {
		stack.dataStore = append(stack.dataStore, value)
		stack.size++
	}
}

func (stack *ArrayStack) Top() interface{} {
	if (stack.IsEmpty()) {
		return nil
	}
	return stack.dataStore[stack.size - 1]
}
