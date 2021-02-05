package BinarySearch

import (
	"fmt"
	"testing"
)

func TestBinarySearchFirstEq(t *testing.T) {
	arr := []int{1, 2, 5, 6, 9, 9, 10, 10, 10, 11}
	index := -1
	index = BinarySearch(arr, 7)

	fmt.Println("index = ", index)
	index = BinarySearchFirstGe(arr, 7)
	fmt.Println("index = ", index)

}
