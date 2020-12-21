package ArrayListStack

import (
	"code/List"
	"code/Stack"
)
import "code/List/ArrayList"

type ArrayListStack struct {
	arrayList *ArrayList.ArrayList
}

func NewArrayListStack(capacity uint) *ArrayListStack {
	stack := new(ArrayListStack)
	if capacity >= Stack.INIT_CAPACITY {
		stack.arrayList = ArrayList.NewArrayList(capacity)
	} else {
		stack.arrayList = ArrayList.NewArrayList(Stack.INIT_CAPACITY)
	}
	return stack
}

func (stack *ArrayListStack) Size() uint {
	return stack.arrayList.Size()
}

func (stack *ArrayListStack) Capacity() uint {
	return stack.arrayList.Capacity()
}

func (stack *ArrayListStack) IsFull() bool {
	return stack.Size() >= stack.Capacity()
}

func (stack *ArrayListStack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *ArrayListStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	list := stack.arrayList
	res, _ := list.Get(list.Size() - 1)
	_ = list.Delete(list.Size() - 1)
	return res
}

func (stack *ArrayListStack) Push(value interface{}) {
	if !stack.IsFull() {
		stack.arrayList.Append(value)
	}
}

func (stack *ArrayListStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	res, _ := stack.arrayList.Get(stack.arrayList.Size() - 1)
	return res
}

func (stack *ArrayListStack) Iterator() List.Iterator {
	return stack.arrayList.Iterator()
}
