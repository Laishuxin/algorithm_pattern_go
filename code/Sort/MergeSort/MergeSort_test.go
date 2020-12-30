package MergeSort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{0, 1, -1, 5, 0, 100, -50}
	MergeSort(arr)
	fmt.Println(arr)
}

func TestMergeSortWithInsert(t *testing.T) {
	length := 30
	arr := make([]int, 0, length)
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Int()%length)
	}
	MergeSort(arr)
	fmt.Println(arr)
}
