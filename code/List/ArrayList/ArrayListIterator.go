package ArrayList

import "errors"

type ArrayListIterator struct {
	list         *ArrayList
	currentIndex uint
}

func (it *ArrayListIterator) HasNext() bool {
	return it.currentIndex < it.list.Size()
}

func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("has not next")
	}
	value, err := it.list.Get(it.currentIndex);
	it.currentIndex++
	return value, err
}

func (it *ArrayListIterator) Remove() {
	it.currentIndex--
	_ = it.list.Delete(it.currentIndex)
}

func (it *ArrayListIterator) GetIndex() uint {
	return it.currentIndex
}
