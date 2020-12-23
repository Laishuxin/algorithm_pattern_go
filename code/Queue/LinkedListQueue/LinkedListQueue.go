package LinkedListQueue

import (
	"code/List/LinkedList"
	"code/Queue"
)

type LinkedListQueue struct {
	linkedList *LinkedList.LinkedList
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{linkedList: LinkedList.NewLinkedList()}
}

func (queue *LinkedListQueue) Size() uint {
	return queue.linkedList.Size()
}

func (queue *LinkedListQueue) Cap() uint {
	return queue.Size()
}

func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.Size() == 0
}

func (queue *LinkedListQueue) EnQueue(newVal interface{}) error {
	queue.linkedList.Append(newVal)
	return nil
}

func (queue *LinkedListQueue) DeQueue() (interface{}, error) {
	front, err := queue.linkedList.Front()
	if err != nil {
		return nil, Queue.ErrEmptyQueue
	}
	queue.linkedList.Shift()
	return front, nil
}

func (queue *LinkedListQueue) Front() (interface{}, error) {
	front, err := queue.linkedList.Front()
	if err != nil {
		return nil, Queue.ErrEmptyQueue
	}
	return front, nil
}

func (queue *LinkedListQueue) Rear() (interface{}, error) {
	rear, err := queue.linkedList.Rear()
	if err != nil {
		return nil, Queue.ErrEmptyQueue
	}
	return rear, nil
}

func (queue *LinkedListQueue) Clear() {
	queue.linkedList.Clear()
}
