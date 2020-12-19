package ArrayList

import (
	"errors"
	"fmt"
)

func Show() {
	fmt.Println("ArrayList.show()")
}

const INIT_CAPACITY int = 8
const INDEX_OUT_OF_RNAGE_ERR_MSG string = "index out of range!!"

type List interface {
	Size() int
	Capacity() int
	Get(index int) (interface{}, error)
	Set(index int, newVal interface{}) error
	Insert(index int, newVal interface{}) error
	Append(newVal interface{})
	Clear()
	Delete(index int) error
	String() string
	Iterator() Iterator
}


func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator)
	it.currentIndex = 0
	it.list = list
	return it
}

type ArrayList struct {
	dataStore []interface{}
	size      int
}

func NewArrayList() *ArrayList {
	arrayList := new(ArrayList)
	arrayList.dataStore = make([]interface{}, 0, INIT_CAPACITY)
	arrayList.size = 0
	return arrayList
}

func (list *ArrayList) Size() int {
	return list.size
}

func (list *ArrayList) Capacity() int {
	return cap(list.dataStore)
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New(INDEX_OUT_OF_RNAGE_ERR_MSG)
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Set(index int, newVal interface{}) error {
	if index < 0 || index >= list.size {
		return errors.New(INDEX_OUT_OF_RNAGE_ERR_MSG)
	}
	list.dataStore[index] = newVal
	return nil
}

func (list *ArrayList) Insert(index int, newVal interface{}) error {
	if index < 0 || index >= list.size {
		return errors.New(INDEX_OUT_OF_RNAGE_ERR_MSG)
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
	if list.size >= cap(list.dataStore) {
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
	list.dataStore = make([]interface{}, 0, INIT_CAPACITY)
	list.size = 0
}

func (list *ArrayList) Delete(index int) error {
	if index < 0 || index >= list.size {
		return errors.New(INDEX_OUT_OF_RNAGE_ERR_MSG)
	}

	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.size--
	return nil
}

func (list *ArrayList) String() string {
	return fmt.Sprintf("%v", list.dataStore)
}
