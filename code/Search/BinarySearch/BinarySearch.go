package BinarySearch

import (
	"fmt"
)

// BinarySearch : search data in a descending order array.
func BinarySearch(arr []int, data int) int {
	if arr == nil {
		return -1
	}
	low := 0
	hig := len(arr)
	var mid int
	for low < hig {
		mid = (low + hig) / 2
		midVal := arr[mid]
		if data < midVal {
			hig = mid
		} else if midVal < data {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func binarySearchTest(arr []int, data int) int {
	if arr == nil {
		return -1
	}
	lo := 0
	hi := len(arr)
	var mi int
	times := 0
	for lo < hi {
		times++
		fmt.Println("search data use times: ", times)
		mi = (lo + hi) / 2
		miVal := arr[mi]
		if data < miVal {
			hi = mi
		} else if data > miVal {
			lo = mi + 1
		} else {
			return mi
		}
	}
	return -1
}
