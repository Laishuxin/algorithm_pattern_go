package BubbleSort

import "code/Sort"

func BubbleSort(arr []int) error {
	if arr == nil {
		return Sort.ErrArrIsNil
	}
	size := len(arr)
	isSorted := false
	for !isSorted {
		isSorted = true
		for i := 0; i < size-1; i++ {
			if arr[i] < arr[i+1] {
				Sort.Swap(arr, i, i+1)
				isSorted = false
			}
		}
		size--
	}
	return nil
}

func QuickBubbleSort(arr []int) error {
	if arr == nil {
		return Sort.ErrArrIsNil
	}
	lastIndex := len(arr) - 1
	var rBound int
	isSorted := false

	for !isSorted {
		isSorted = true
		rBound = lastIndex
		for i := 0; i < rBound; i++ {
			if arr[i] < arr[i+1] {
				Sort.Swap(arr, i, i+1)
				isSorted = false
				rBound = i
			}
		}
	}

	return nil
}
