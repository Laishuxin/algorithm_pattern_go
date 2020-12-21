package Stack

const INIT_CAPACITY = 8
type Stack interface {
	Size() uint
	Capacity() uint
	IsFull() bool
	IsEmpty() bool
	Pop() interface{}
	Push(value interface{})
	Top() interface{}
}
