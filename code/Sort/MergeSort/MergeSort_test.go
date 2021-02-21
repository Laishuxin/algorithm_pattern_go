package MergeSort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	length := 30000
	arr := make([]int, 0, length)
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Int()%length)
	}
	fmt.Printf("sort.IntsAreSorted(arr) = %+v\n", sort.IntsAreSorted(arr))
	fmt.Printf("size =  = %+v\n", length)

	MergeSort(arr)
	fmt.Printf("sort.IntsAreSorted(arr) = %+v\n", sort.IntsAreSorted(arr))
}
