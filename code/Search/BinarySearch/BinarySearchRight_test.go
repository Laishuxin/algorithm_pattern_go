package BinarySearch

import (
	"fmt"
	"testing"
)

func TestBinarySearchRight(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 3, 4, 5, 6}
	index := -1
	index = BinarySearch(arr, 3)

	fmt.Println("index = ", index)
	index = BinarySearchRight(arr, 3)
	fmt.Println("index = ", index)
}
