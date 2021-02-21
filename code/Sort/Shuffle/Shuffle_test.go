package Sort

import (
	"fmt"
	"sort"
	"testing"
)

func _shellSort(arr []int) {
	length := len(arr)
	gap := 1
	for gap < length/3 {
		gap = 3*gap + 1
	}
	for gap >= 1 {
		for i := gap; i < length; i += gap {
			for j := i; j >= gap && arr[j] < arr[j-gap]; j -= gap {
				temp := arr[j]
				arr[j] = arr[j-gap]
				arr[j-gap] = temp
			}
		}
		gap = gap / 3
	}
}
func TestShuffle(t *testing.T) {
	length := 20
	arr := make([]int, length, length)
	for k, _ := range arr {
		arr[k] = k
	}
	fmt.Printf("sort.IntsAreSorted(arr) = %+v\n", sort.IntsAreSorted(arr))
	Shuffle(arr)
	fmt.Printf("arr = %+v\n", arr)
	_shellSort(arr)
	fmt.Printf("arr = %+v\n", arr)

}
