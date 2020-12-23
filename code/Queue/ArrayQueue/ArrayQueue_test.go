package ArrayQueue

import (
	"fmt"
	"testing"
)

func TestNewArrayQueue(t *testing.T) {
	queue := NewArrayQueue(10)
	queue.EnQueue("str1")
	queue.EnQueue("str2")
	queue.EnQueue("str3")
	queue.EnQueue("str4")
	queue.EnQueue("str4")
	queue.EnQueue("str4")
	fmt.Println("size = ", queue.Size())
	fmt.Println("cap = ", queue.Cap())
	queue.Clear()
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
}