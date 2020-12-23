package RingQueue

import (
	"fmt"
	"testing"
)

func TestNewRingQueue(t *testing.T) {
	queue := NewRingQueue(8)
	fmt.Println("queue.ToString()", queue.ToString())
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

func TestRingQueue_EnQueue(t *testing.T) {
	queue:=NewRingQueue(10)
	for i:=0; ; i++{
		queue.EnQueue(fmt.Sprintf("str%v", i))
		fmt.Println(queue.DeQueue())
	}
}