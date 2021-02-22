package MergeSort

import (
	"code/Sort"
)

const (
	threshold = 30
)

func swap(arr []int, n int, m int) {
	temp := arr[n]
	arr[n] = arr[m]
	arr[m] = temp
}

func insertSort(arr []int, lo int, hi int) {
	for i := lo + 1; i < hi; i++ {
		for j := i; j > lo && arr[j] < arr[j-1]; j-- {
			swap(arr, j, j-1)
		}
	}
}


func MergeSort(arr []int) error {
	if arr == nil {
		return Sort.ErrArrIsNil
	}
	mergeSort(arr, 0, len(arr))
	return nil
}

func mergeSort(arr []int, lo int, hi int) {
	if (hi - lo) < 2 {
		return
	} else if hi-lo < threshold {
		insertSort(arr, lo, hi)
	}
	mi := lo + (hi-lo)/2
	mergeSort(arr, lo, mi)
	mergeSort(arr, mi, hi)
	merge(arr, lo, mi, hi)
}

func merge(arr []int, lo int, mi int, hi int) {
	lLength := mi - lo
	lPart := make([]int, lLength, lLength)
	for i := 0; i < lLength; i++ {
		lPart[i] = arr[lo+i]
	}
	l, r, c := 0, mi, lo
	for ; l < lLength && r < hi; c++ {
		if lPart[l] <= arr[r] {
			arr[c] = lPart[l]
			l++
		} else if arr[r] < lPart[l] {
			arr[c] = arr[r]
			r++
		}
	}
	for ; c < hi && l < lLength; c++ {
		arr[c] = lPart[l]
		l++
	}
}
