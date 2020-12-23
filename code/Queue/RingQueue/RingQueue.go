package RingQueue

import (
	"code/Queue"
	"fmt"
)

// 环形队列
type RingQueueInterface interface {
	Size() int
	Cap() int
	Front() (front interface{}, err error)
	Rear() (rear interface{}, err error)
	IsEmpty() bool
	IsFull() bool
	Clear()
	EnQueue(newVal interface{}) (err error)
	DeQueue() (data interface{}, err error)
}

type RingQueue struct {
	dataStore []interface{}
	capacity  int
	// 队头
	front int
	// 队尾
	rear int
}

func NewRingQueue(capacity int) *RingQueue {
	queue := new(RingQueue)
	if capacity < Queue.INIT_QUEUE_CAP {
		capacity = Queue.INIT_QUEUE_CAP
	}
	queue.dataStore = make([]interface{}, capacity, capacity)
	queue.capacity = capacity
	queue.front = 0
	queue.rear = 0
	return queue
}

func (queue *RingQueue) Size() int {
	return (queue.front - queue.rear + queue.capacity) % queue.capacity
}

func (queue *RingQueue) Cap() int {
	return queue.capacity
}

func (queue *RingQueue) Front() (front interface{}, err error) {
	if queue.IsEmpty() {
		return nil, Queue.ErrEmptyQueue
	}
	return queue.dataStore[queue.front], nil
}

func (queue *RingQueue) Rear() (rear interface{}, err error) {
	if queue.IsEmpty() {
		return nil, Queue.ErrEmptyQueue
	}
	return queue.dataStore[queue.rear], nil
}

func (queue *RingQueue) IsEmpty() bool {
	return queue.front == queue.rear
}

func (queue *RingQueue) IsFull() bool {
	return (queue.rear+1)%queue.capacity == queue.front
}

func (queue *RingQueue) Clear() {
	capacity := queue.capacity
	for i := 0; i < capacity; i++ {
		queue.dataStore[i] = nil
	}
	queue.front = 0
	queue.rear = 0
}

func (queue *RingQueue) EnQueue(newVal interface{}) (err error) {
	if queue.IsFull() {
		return Queue.ErrEmptyQueue
	}
	queue.dataStore[queue.rear] = newVal
	queue.rear = (queue.rear + 1) % queue.capacity
	return nil
}

func (queue *RingQueue) DeQueue() (data interface{}, err error) {
	if queue.IsEmpty() {
		return nil, Queue.ErrEmptyQueue
	}
	data = queue.dataStore[queue.front]
	queue.dataStore[queue.front] = nil;
	queue.front = (queue.front + 1) % queue.capacity
	return data, nil
}

// TODO: 修改
func (queue *RingQueue) ToString() string {
	return fmt.Sprintf("%v", queue.dataStore)
}
