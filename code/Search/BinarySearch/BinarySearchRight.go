package BinarySearch

// 二分查找：并且保证找到的数在最右边

func BinarySearchRight(arr []int, data int) int {
	if arr == nil {
		return -1
	}
	index := -1
	lo := 0
	hi := len(arr)
	var mi int
	for lo < hi {
		mi = (lo + hi) / 2
		miVal := arr[mi]
		if miVal < data {
			lo = mi + 1
		} else if data < miVal {
			hi = mi
		} else {
			if mi == (hi-1) || arr[mi] != data {
				index = mi
				break
			}
			lo = mi + 1
		}
	}
	return index
}
