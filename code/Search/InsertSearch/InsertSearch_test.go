package InsertSearch

import (
	"fmt"
	"testing"
)

func TestInsertSearch(t *testing.T) {
	length := 100
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = 1
	}
	// index := InsertSearch(arr, arr[5])
	index := InsertSearch(arr, 1)

	fmt.Println("index = ", index)
}
