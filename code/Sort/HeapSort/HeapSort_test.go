package HeapSort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int {-1, 5, 4, -2, 0}
	HeapSort(arr)
	fmt.Println(arr)
}