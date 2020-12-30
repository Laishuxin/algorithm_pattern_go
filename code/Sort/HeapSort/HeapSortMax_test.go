package HeapSort

import (
	"fmt"
	"testing"
)

func TestHeapSortMax(t *testing.T) {
	arr := []int{1, 2, -1, 5, -5, 0}
	HeapSortMax(arr)
	fmt.Println(arr)
}