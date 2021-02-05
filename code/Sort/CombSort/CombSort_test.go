package CombSort

import (
	"fmt"
	"sort"
	"testing"
)

func TestCombSort(t *testing.T) {
	// length := 63
	// arr := make([]int, length, length)
	// for i := 0; i < length; i++ {
	// 	arr[i] = rand.Int() % length
	// }
	arr := []int{0, 1, 4, 3, 4, 5, 6}
	fmt.Println("arr is sorted? ", sort.IntsAreSorted(arr))
	CombSort(arr)
	fmt.Println("arr is sorted? ", sort.IntsAreSorted(arr))
}
