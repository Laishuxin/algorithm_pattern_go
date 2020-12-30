package QuickSort

import (
	"fmt"
	"testing"
)

func TestQuickSortSimplicityVersion(t *testing.T) {
	arr := []int{5, 1, -3, 10, 2}
	sortedArr := QuickSortSimplicityVersion(arr)
	fmt.Println(arr)
	fmt.Println(sortedArr)

}
