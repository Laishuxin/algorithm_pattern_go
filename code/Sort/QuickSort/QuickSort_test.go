package QuickSort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	length := 600
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Int() % length
	}

	fmt.Printf("quicksort\n")
	fmt.Printf("arr1 = %+v\n", arr)
	fmt.Println("arr is sorted ? ", sort.IntsAreSorted(arr))
	QuickSort(arr)
	fmt.Printf("arr1 = %+v\n", arr)
	fmt.Println("arr is sorted ? ", sort.IntsAreSorted(arr))
}
