package ArrayQueue

import (
	"code/Queue"
)

type ArrayQueue struct {
	dataStore []interface{}
	size      uint
}

func NewArrayQueue(initCap uint) *ArrayQueue {
	queue := new(ArrayQueue)
	if (initCap < Queue.INIT_QUEUE_CAP) {
		initCap = Queue.INIT_QUEUE_CAP
	}
	queue.dataStore = make([]interface{}, 0, initCap)
	queue.size = 0
	return queue
}

func (queue *ArrayQueue) Size() uint {
	return queue.size
}

func (queue *ArrayQueue) Cap() uint {
	return uint(cap(queue.dataStore))
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *ArrayQueue) EnQueue(newVal interface{}) error {
	queue.dataStore = append(queue.dataStore, newVal)
	queue.size++
	return nil
}

func (queue *ArrayQueue) DeQueue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, Queue.ErrEmptyQueue
	}
	head := queue.dataStore[0]
	if queue.size > 1 {
		queue.dataStore = queue.dataStore[1:]
	} else {
		queue.dataStore = make([]interface{}, 0, Queue.INIT_QUEUE_CAP)
	}
	queue.size--
	return head, nil
}

func (queue *ArrayQueue) Front() (interface{}, error) {
	if (queue.IsEmpty()) {
		return nil, Queue.ErrEmptyQueue
	}
	return queue.dataStore[0], nil
}

func (queue *ArrayQueue) Rear() (interface{}, error) {
	if (queue.IsEmpty()) {
		return nil, Queue.ErrEmptyQueue
	}
	return queue.dataStore[queue.size - 1], nil
}

func (queue *ArrayQueue) Clear() {
	queue.dataStore = nil
	queue.dataStore = make([]interface{}, 0, Queue.INIT_QUEUE_CAP)
	queue.size = 0
}
