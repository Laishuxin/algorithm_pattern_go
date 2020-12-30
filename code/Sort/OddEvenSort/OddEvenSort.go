package OddEvenSort

import "code/Sort"

// 奇偶排序

func OddEvenSort(arr []int) error {
	if arr == nil {
		return Sort.ErrArrIsNil
	}
	length := len(arr)
	if length < 2 {
		return nil
	}
	isSorted := false
	for !isSorted {
		isSorted = true
		// odd sort
		for i := 1; i < length-1; i += 2 {
			if arr[i] < arr[i+1] {
				Sort.Swap(arr, i, i+1)
				isSorted = false
			}
		}
		// even sort
		for i := 0; i < length-1; i += 2 {
			if arr[i] < arr[i+1] {
				Sort.Swap(arr, i, i+1)
				isSorted = false
			}
		}
	}
	return nil
}
