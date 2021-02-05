package BinarySearch

import (
	"code/Sort/QuickSort"
	"fmt"
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var index int
	arr := []int{-1, 5, 4, -2, 0}
	// QuickSort.QuickSort(arr)
	sort.Ints(arr)
	fmt.Println("arr is sorted? ", sort.IntsAreSorted((arr)))
	fmt.Println(arr)
	data := 4
	index = BinarySearch(arr, data)
	fmt.Println("index = ", index)
}

func Test_binarySearchTest(t *testing.T) {
	length := 100
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		// arr[i] = rand.Int() % length
		arr[i] = i
	}
	QuickSort.QuickSortIncreasingOrder(arr)
	fmt.Println("arr is sorted? ", sort.IntsAreSorted((arr)))
	data := arr[5]
	fmt.Println("index = (5?)", binarySearchTest(arr, data))
}
