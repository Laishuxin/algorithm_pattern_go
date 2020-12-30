package QuickSort

import (
	"code/Sort"
	"math/rand"
)

func QuickSort(arr []int) error {
	if arr == nil {
		return Sort.ErrArrIsNil
	}
	quickSortMin(arr, 0, len(arr))
	return nil
}

func quickSortMin(arr []int, lo int, hi int) {
	if hi-lo < 2 {
		return
	}
	l := lo
	r := hi - 1
	pivotIndex := lo
	// ensure pivot is random
	Sort.Swap(arr, rand.Intn(r-l)+l, pivotIndex)
	pivot := arr[pivotIndex]
	for l < r {
		for l < r && arr[r] < pivot {
			r--
		}
		for l < r && pivot <= arr[l] {
			l++
		}
		if l < r {
			Sort.Swap(arr, l, r)
		}
	}
	Sort.Swap(arr, pivotIndex, l)
	quickSortMin(arr, lo, l)
	quickSortMin(arr, l+1, hi)
}

func quickSortMax(arr []int, lo int, hi int) {
	if hi-lo < 2 {
		return
	}
	l := lo
	r := hi - 1
	pivotIndex := lo
	Sort.Swap(arr, rand.Intn(r-l)+l, pivotIndex)
	pivot := arr[pivotIndex]
	for l < r {
		for l < r && pivot < arr[r] {
			r--
		}
		for l < r && arr[l] <= pivot {
			l++
		}
		if l < r {
			Sort.Swap(arr, l, r)
		}
	}
	Sort.Swap(arr, pivotIndex, l)
	quickSortMax(arr, lo, l)
	quickSortMax(arr, l+1, hi)
}
