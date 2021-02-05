package BinarySearch

import (
	"fmt"
	"testing"
)

func TestBinarySearchLeft(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 3, 4, 5, 6}
	index := -1
	index = BinarySearch(arr, 3)

	fmt.Println("index = ", index)
	index = BinarySearchLeft(arr, 3)
	fmt.Println("index = ", index)

}
