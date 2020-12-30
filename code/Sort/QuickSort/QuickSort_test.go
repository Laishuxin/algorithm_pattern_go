package QuickSort

import (
	"fmt"
	"testing"
)

func TestQuickSort1(t *testing.T) {
	arr := []int{-1, 4, 1993, 2, -4, 5, -1}
	QuickSort(arr)
	fmt.Println(arr)
}
