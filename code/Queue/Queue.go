package Queue


const INIT_QUEUE_CAP = 8;

type Queue interface {
	Size() uint
	Cap() uint
	IsEmpty() bool
	// 入队
	EnQueue(newVal interface{}) error
	// 出队
	DeQueue() (interface{}, error)
	Front() (interface{},error)
	Rear() (interface{},error)
	Clear()
}


