package QuickSort

import "math/rand"

func QuickSortSimplicityVersion(arr []int) []int {
	if arr == nil {
		return nil
	}
	length := len(arr)
	if length <= 1 {
		return []int{arr[0]}
	}

	pivot := arr[rand.Int()%length]
	less := make([]int, 0, 0)
	great := make([]int, 0, 0)
	equal := make([]int, 0, 0)
	equal = append(equal, pivot)
	for i := 1; i < length; i++ {
		curr := arr[i]
		if curr < pivot {
			less = append(less, curr)
		} else if curr > pivot {
			great = append(great, curr)
		} else {
			equal = append(equal, curr)
		}
	}
	if len(less) > 0 {
		less = QuickSortSimplicityVersion(less)
	}
	if len(great) > 0 {
		great = QuickSortSimplicityVersion(great)
	}
	return append(append(less, equal...), great...)
}
