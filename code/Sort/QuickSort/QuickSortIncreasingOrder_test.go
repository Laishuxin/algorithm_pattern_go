package QuickSort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func Test_insertSort(t *testing.T) {
	length := 100
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Int() % 100
	}
	fmt.Println("arr is sorted? ", sort.IntsAreSorted(arr))
	QuickSortIncreasingOrder(arr)
	fmt.Println("arr is sorted? ", sort.IntsAreSorted(arr))
}
