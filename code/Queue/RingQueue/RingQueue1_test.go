package RingQueue

import (
	"fmt"
	"testing"
)

func TestNewRingQueue1(t *testing.T) {
	queue := NewRingQueue1(8)
	fmt.Println("queue.size = ", queue.Size())
	fmt.Println("queue.isEmpty() = ", queue.IsEmpty())
	queue.EnQueue("3")
	queue.EnQueue("4")
	queue.EnQueue("5")
	queue.EnQueue("6")
	queue.EnQueue("7")
	queue.EnQueue("8")
	queue.EnQueue("9")
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	queue.EnQueue("1000")
	queue.EnQueue("2000")
	queue.EnQueue("3000")
	queue.EnQueue("4000")
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
}