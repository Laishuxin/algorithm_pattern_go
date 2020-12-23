package RingQueue

import "code/Queue"

// 环形队列
type RingQueue1 struct {
	dataStore []interface{}
	front     uint
	rear      uint
	capacity  uint
}

func NewRingQueue1(capacity uint) *RingQueue1 {
	queue := new(RingQueue1)
	if capacity < Queue.INIT_QUEUE_CAP {
		capacity = Queue.INIT_QUEUE_CAP
	}
	queue.dataStore = make([]interface{}, capacity, capacity)
	queue.capacity = capacity
	queue.front = 0
	queue.rear = 0
	return queue
}

func (queue *RingQueue1) Size() uint {
	if queue.rear >= queue.front {
		return queue.rear - queue.front
	}
	return queue.front - queue.rear
}

func (queue *RingQueue1) Cap() uint {
	return queue.capacity
}

func (queue *RingQueue1) IsEmpty() bool {
	return queue.front == queue.rear
}

func (queue *RingQueue1) IsFull() bool {
	return queue.Size() + 1 == queue.capacity
}

func (queue *RingQueue1) EnQueue(newVal interface{}) error {
	if queue.IsFull() {
		return Queue.ErrFullQueue
	}
	queue.dataStore[queue.rear] = newVal
	queue.rear++
	if queue.rear >= queue.capacity {
		queue.rear = 0
	}
	return nil
}

func (queue *RingQueue1) DeQueue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, Queue.ErrEmptyQueue
	}
	data := queue.dataStore[queue.front]
	queue.dataStore[queue.front] = nil
	queue.front++
	if queue.front >= queue.capacity {
		queue.front = 0
	}

	return data, nil
}

func (queue *RingQueue1) Front() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, Queue.ErrEmptyQueue
	}
	return queue.dataStore[queue.front], nil
}

func (queue *RingQueue1) Rear() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, Queue.ErrEmptyQueue
	}
	return queue.dataStore[queue.rear], nil
}

func (queue *RingQueue1) Clear() {
	capacity := queue.capacity
	for i := uint(0); i < capacity; i++ {
		queue.dataStore[i] = nil
	}
	queue.front = 0
	queue.rear = 0
}
