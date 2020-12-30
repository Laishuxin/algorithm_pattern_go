package MergeSort

import (
	"code/Sort"
	"code/Sort/InsertSort"
)

func MergeSort(arr []int) error {
	if arr == nil {
		return Sort.ErrArrIsNil
	}
	mergeSortWithInsert(arr, 0, len(arr))
	return nil
}


func merge(arr []int, lo int, mi int, hi int) {
	currArr := arr[lo:hi]

	lLen := mi - lo
	lPart := make([]int, lLen)
	for i := 0; i < lLen; i++ {
		lPart[i] = currArr[i]
	}
	rLen := hi - mi
	rPart := arr[mi:]
	var (
		l = 0
		r = 0
		c = 0
	)
	for l < lLen && r < rLen {
		lVal := lPart[l]
		rVal := rPart[r]
		if rVal <= lVal {
			currArr[c] = lVal
			l++
		} else {
			currArr[c] = rVal
			r++
		}
		c++
	}

	for l < lLen {
		currArr[c] = lPart[l]
		l++
		c++
	}
	for r < rLen {
		currArr[c] = rPart[r]
		r++
		c++
	}
	lPart = nil
}

func mergeSort(arr []int, lo int, hi int) {
	if (hi - lo) < 2 {
		return
	}
	mi := (hi + lo) / 2
	mergeSort(arr, lo, mi)
	mergeSort(arr, mi, hi)
	merge(arr, lo, mi, hi)
}

func mergeSortWithInsert(arr []int, lo int, hi int) {
	if (hi - lo) < 2 {
		return
	} else if (hi - lo) < 10 {
		//println("use insert sort")
		_ = InsertSort.InsertSort(arr[lo:hi])
		return
	}

	mi := (hi + lo) / 2
	mergeSortWithInsert(arr, lo, mi)
	mergeSortWithInsert(arr, mi, hi)
	merge(arr, lo, mi, hi)
}
