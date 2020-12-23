package LinkedList

import (
	"code/List"
	"fmt"
)

// 双向链表

type LinkedList struct {
	head *Node // 头指针
	rear *Node // 尾指针
	size uint  // 链表的大小
}

func NewLinkedList() *LinkedList {
	list := new(LinkedList)
	return list
}

func (list *LinkedList) Size() uint {
	return list.size
}

func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *LinkedList) Get(index uint) (interface{}, error) {
	node := list.locateNode(index)
	if node == nil {
		return nil, List.ErrIndexOutOfRange
	}
	return node.data, nil
}

func (list *LinkedList) Set(index uint, newVal interface{}) error {
	node := list.locateNode(index)
	if node == nil {
		return List.ErrIndexOutOfRange
	}
	node.data = newVal
	return nil
}

func (list *LinkedList) Insert(index uint, newVal interface{}) error {
	size := list.size
	if index >= size {
		return List.ErrIndexOutOfRange
	}
	if index != 0 && index != size-1 {
		newNode := NewNode(newVal)
		currNode := list.locateNode(index)
		newNode.pPrev = currNode
		newNode.pNext = currNode.pNext
		currNode.pNext.pPrev = newNode
		currNode.pNext = newNode
		list.size++
	} else {
		if index == 0 {
			list.Push(newVal)
		} else {
			list.Append(newVal)
		}
	}
	return nil
}

func (list *LinkedList) Push(newVal interface{}) {
	newNode := NewNode(newVal)
	if list.IsEmpty() {
		list.head = newNode
		list.rear = newNode
		list.size++
		return
	}
	newNode.pNext = list.head
	list.head.pPrev = newNode
	list.head = newNode
	list.size++
}

func (list *LinkedList) Append(newVal interface{}) {
	newNode := NewNode(newVal)
	if list.IsEmpty() {
		list.head = newNode
		list.rear = newNode
		list.size++
		return
	}
	newNode.pPrev = list.rear
	list.rear.pNext = newNode
	list.rear = newNode
	list.size++
}

// TODO: done it
func (list *LinkedList) Delete(index uint) error {
	size := list.size
	if index >= size {
		return List.ErrIndexOutOfRange
	}
	if index != 0 && index != size-1 {
		currNode := list.locateNode(index)
		currNode.pPrev.pNext = currNode.pNext
		currNode.pNext.pPrev = currNode.pPrev
		currNode = nil
	} else {
		if index == 0 {
			return list.Shift()
		} else {
			return list.Pop()
		}
	}
	return nil
}

func (list *LinkedList) Shift() error {
	if list.IsEmpty() {
		return List.ErrEmptyList
	}
	if list.size > 1 {
		newHeadNode := list.head.pNext
		newHeadNode.pPrev = nil
		list.head = newHeadNode
		list.size--
		return nil
	}
	list.head = nil
	list.rear = nil
	list.size--
	return nil
}

func (list *LinkedList) Pop() error {
	if list.IsEmpty() {
		return List.ErrEmptyList
	}
	if list.size > 1 {
		newRearNode := list.rear.pPrev
		newRearNode.pNext = nil
		list.rear = newRearNode
		list.size--
		return nil
	}
	list.head = nil
	list.rear = nil
	list.size--
	return nil
}

func (list *LinkedList) Clear() {
	for list.Pop() == nil {
	}
}

func (list *LinkedList) String() string {
	if list.size == 0 {
		return ""
	}
	str := ""
	for currNode := list.head; currNode != nil; currNode = currNode.pNext {
		str += fmt.Sprintf("%v", currNode.data)
	}
	return str
}

func (list *LinkedList) Iterator() List.Iterator {
	return NewLinkedListIterator(list)
}

// ==== 私有方法 ====
// 根据索引定位到指定节点
func (list *LinkedList) locateNode(index uint) *Node {
	if index >= list.size {
		return nil
	}
	currNode := list.head
	for i := uint(0); i < index; i++ {
		currNode = currNode.pNext
	}
	return currNode
}
