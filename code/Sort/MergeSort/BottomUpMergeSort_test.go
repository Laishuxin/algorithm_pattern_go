package MergeSort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestBottomUpMergeSort(t *testing.T) {
	length := 150
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Int() % length
	}
	fmt.Printf("sort.IntsAreSorted(arr) = %+v\n", sort.IntsAreSorted(arr))
	BottomUpMergeSort(arr)
	fmt.Printf("sort.IntsAreSorted(arr) = %+v\n", sort.IntsAreSorted(arr))

}
