package List

const INIT_CAPACITY = 8

type List interface {
	Size() uint
	Capacity() uint
	Get(index uint) (interface{}, error)
	Set(index uint, newVal interface{}) error
	Insert(index uint, newVal interface{}) error
	Append(newVal interface{})
	Clear()
	Delete(index uint) error
	String() string
	Iterator() Iterator
}