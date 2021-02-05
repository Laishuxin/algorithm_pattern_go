package QuickSort

import (
	"code/Sort"
	"math/rand"
)

const SORT_THRESHOLD = 4

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func insertSort(arr []int, lo int, hi int) {
	for i := lo + 1; i < hi; i++ {
		temp := arr[i]
		var j int
		for j = i - 1; j >= lo && temp < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
}

func QuickSortIncreasingOrder(arr []int) error {
	if arr == nil {
		return Sort.ErrArrIsNil
	}
	quickSort(arr, 0, len(arr))
	return nil
}

func quickSort(arr []int, lo int, hi int) {
	if hi-lo < SORT_THRESHOLD {
		insertSort(arr, lo, hi)
		return
	}
	pivotIndex := partition(arr, lo, hi)
	quickSort(arr, lo, pivotIndex)
	quickSort(arr, pivotIndex+1, hi)
}

func partition(arr []int, lo int, hi int) int {
	left := lo
	right := hi - 1
	pivotIndex := left
	swap(arr, pivotIndex, rand.Int()%(right-left)+left)
	pivot := arr[pivotIndex]
	for left < right {
		for left < right && pivot < arr[right] {
			right--
		}
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			swap(arr, left, right)
		}
	}
	swap(arr, pivotIndex, left)
	return left
}
