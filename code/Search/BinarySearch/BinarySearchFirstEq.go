package BinarySearch

// 二分查找：并且保证找到索引在第一个大于等于该数

func BinarySearchFirstGe(arr []int, data int) int {
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
		} else {
			if mi == 0 || arr[mi-1] < data {
				index = mi
				break
			} else {
				hi = mi
			}
		}
	}
	return index
}
