package LinkedList

import "code/List"

type LinkedListIterator struct {
	list         *LinkedList
	currentIndex uint
}

func NewLinkedListIterator(list *LinkedList) *LinkedListIterator {
	return &LinkedListIterator{list: list}
}

func (it *LinkedListIterator) HasNext() bool {
	return it.currentIndex < it.list.size
}

func (it *LinkedListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, List.ErrHasNotNext
	}
	nextData := it.list.locateNode(it.currentIndex).data
	it.currentIndex++
	return nextData, nil
}

func (it *LinkedListIterator) Remove() {
	it.currentIndex--
	_ = it.list.Delete(it.currentIndex)
}

func (it *LinkedListIterator) GetIndex() uint {
	return it.currentIndex
}
