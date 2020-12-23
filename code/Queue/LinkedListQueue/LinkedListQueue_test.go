package LinkedListQueue

import (
	"fmt"
	"testing"
)

func TestNewLinkedListQueue(t *testing.T) {
	queue:= NewLinkedListQueue()
	queue.EnQueue("1")
	queue.EnQueue("2")
	queue.EnQueue("3")
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
}