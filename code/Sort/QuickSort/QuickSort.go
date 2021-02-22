package QuickSort

import (
	"code/Sort"
	"math/rand"
)

func QuickSort(arr []int) {
	if arr == nil {
		return
	}
	quickSort(arr, 0, len(arr))
}

func insertSort(arr []int, lo int, hi int) {
	if hi <= lo+1 {
		return
	}
	for i := lo + 1; i < hi; i++ {
		for j := i; j >= lo && arr[j] < arr[j-1]; j-- {
			Sort.Swap(arr, j, j-1)
		}
	}
}
func quickSort(arr []int, lo int, hi int) {
	if hi <= lo+1 {
		return
	}
	pivotIndex := partition(arr, lo, hi)
	quickSort(arr, lo, pivotIndex)
	quickSort(arr, pivotIndex+1, hi)
}

func partition(arr []int, lo int, hi int) int {
	Sort.Swap(arr, lo, rand.Intn(hi-lo)+lo)
	pivot := arr[lo]
	left, right := lo, hi-1
	for left < right {
		for /*left < right &&*/ pivot < arr[right] {
			right--
		}
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			Sort.Swap(arr, left, right)
		}
	}
	Sort.Swap(arr, lo, left)
	return left
}
