package List

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()
	GetIndex() uint
}

type Iterable interface {
	Iterator() Iterator
}
