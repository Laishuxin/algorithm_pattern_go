package LinkedList

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList()
	fmt.Println(list)
}

func TestLinkedList_Push(t *testing.T) {
	list := NewLinkedList()
	list.Push("1")
	list.Push("2")
	list.Push("3")
	list.Push("4")
	list.Push("5")
	list.Push("6")
	fmt.Println(list.String())
}

func TestLinkedList_Append(t *testing.T) {
	list := NewLinkedList()
	list.Push("1")
	list.Push("2")
	list.Push("3")
	list.Push("4")
	list.Push("5")
	list.Push("6")
	fmt.Println(list.String())
}

func TestLinkedList_Pop(t *testing.T) {
	list := NewLinkedList()
	list.Push("1")
	list.Push("2")
	list.Push("3")
	list.Push("4")
	list.Push("5")
	list.Push("6")
	fmt.Println(list.String())
	list.Clear()
	fmt.Println(list.String())
}

func TestLinkedList_Shift(t *testing.T) {
	list := NewLinkedList()
	list.Push("1")
	list.Push("2")
	list.Push("3")
	list.Push("4")
	list.Push("5")
	list.Push("6")
	fmt.Println(list.String())
	list.Shift()
	list.Shift()
	list.Shift()
	list.Shift()
	list.Shift()
	list.Append("2")
	fmt.Println(list.String())
}

func TestLinkedList_Insert(t *testing.T) {
	list := NewLinkedList()
	list.Append("1")
	list.Append("2")
	list.Append("3")
	list.Append("4")
	fmt.Println(list.String())
	list.Insert(0, "0")
	list.Insert(4, "5")
	list.Insert(2, "2.1")
	fmt.Println(list.String())
}

func TestLinkedList_Delete(t *testing.T) {
	list := NewLinkedList()
	list.Append("1")
	list.Append("2")
	list.Append("3")
	list.Append("4")
	list.Append("5")
	fmt.Println(list.String())
	list.Delete(2)
	fmt.Println(list.String())
}