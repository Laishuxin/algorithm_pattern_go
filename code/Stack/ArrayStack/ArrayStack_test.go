package ArrayStack

import (
	"code/Stack"
	"fmt"
	"testing"
)

func TestNewArrayStack(t *testing.T) {
	//stack := NewArrayStack(5)
	var stack Stack.Stack = NewArrayStack(5)
	fmt.Println("size = ", stack.Size())
	fmt.Println("capacity = ", stack.Capacity())
	stack.Push("1")
	fmt.Println("top =(1) ", stack.Top())
	stack.Push("2")
	fmt.Println("top =(2) ", stack.Top())
	stack.Push("3")
	fmt.Println("top =(3) ", stack.Top())
	stack.Push("4")
	fmt.Println("top =(4) ", stack.Top())
	stack.Push("5")
	stack.Push("6")
	stack.Push("7")
	stack.Push("8")
	stack.Push("9")
	fmt.Println("top =(8) ", stack.Top())

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}