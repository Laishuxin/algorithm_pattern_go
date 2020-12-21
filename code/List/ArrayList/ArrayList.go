package ArrayList

import (
	"code/List"
	"errors"
	"fmt"
)

type ArrayList struct {
	dataStore []interface{}
	size      uint
}

func NewArrayList(initialCapacity uint) *ArrayList {
	arrayList := new(ArrayList)
	if (initialCapacity < List.INIT_CAPACITY) {
		initialCapacity = List.INIT_CAPACITY
	}
	arrayList.dataStore = make([]interface{}, 0, initialCapacity)
	arrayList.size = 0
	return arrayList
}

func (list *ArrayList) Iterator() List.Iterator {
	it := new(ArrayListIterator)
	it.currentIndex = 0
	it.list = list
	return it
}

func (list *ArrayList) Size() uint {
	return list.size
}

func (list *ArrayList) Capacity() uint {
	return uint(cap(list.dataStore))
}

func (list *ArrayList) Get(index uint) (interface{}, error) {
	if index >= list.size {
		return nil, errors.New("index out of range!")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Set(index uint, newVal interface{}) error {
	if index >= list.size {
		return errors.New("index out of range!")
	}
	list.dataStore[index] = newVal
	return nil
}

func (list *ArrayList) Insert(index uint, newVal interface{}) error {
	if index >= list.size {
		return errors.New("index out of range!")
	}
	list.ensureCap()
	list.dataStore = list.dataStore[:list.Size()+1]
	for i := list.Size(); i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newVal
	list.size++
	return nil
}

func (list *ArrayList) ensureCap() {
	if list.size >= uint(cap(list.dataStore)) {
		newDataStore := make([]interface{}, list.size, 2*list.size)
		copy(newDataStore, list.dataStore)
		list.dataStore = nil
		list.dataStore = newDataStore
	}
}

func (list *ArrayList) Append(newVal interface{}) {
	list.dataStore = append(list.dataStore, newVal)
	list.size++
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, List.INIT_CAPACITY)
	list.size = 0
}

func (list *ArrayList) Delete(index uint) error {
	if index >= list.size {
		return errors.New("index out of range!")
	}

	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.size--
	return nil
}

func (list *ArrayList) String() string {
	return fmt.Sprintf("%v", list.dataStore)
}
