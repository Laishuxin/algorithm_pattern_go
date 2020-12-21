package ArrayListStack

import (
	"code/Stack"
	"fmt"
	"testing"
)

func TestNewArrayListStack(t *testing.T) {
	var stack Stack.Stack
	stack = NewArrayListStack(8)
	fmt.Println("stack.size = ", stack.Size())
	fmt.Println("stack.capcity = ", stack.Capacity())
	stack.Push("str1")
	stack.Push("str2")
	stack.Push("str3")
	stack.Push("str4")
	stack.Push("str5")
	stack.Push("str6")
	stack.Push("str7")
	stack.Push("str8")
	stack.Push("str9")

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}

}

func TestArrayListStack_Iterator(t *testing.T) {
	stack := NewArrayListStack(8)
	fmt.Println("stack.size = ", stack.Size())
	fmt.Println("stack.capcity = ", stack.Capacity())
	stack.Push("str1")
	stack.Push("str2")
	stack.Push("str3")
	stack.Push("str4")
	stack.Push("str5")
	stack.Push("str6")
	stack.Push("str7")
	stack.Push("str8")
	stack.Push("str9")

	it := stack.Iterator()
	for (it.HasNext()) {
		fmt.Println(it.Next())
	}

	fmt.Println("stack.Size() = ", stack.Size())
}