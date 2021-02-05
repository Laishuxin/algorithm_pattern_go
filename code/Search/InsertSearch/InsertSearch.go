package InsertSearch

// 插值查询

func InsertSearch(arr []int, data int) int {
	if arr == nil {
		return -1
	}
	left := 0
	right := len(arr) - 1
	var mi int
	for left <= right {
		r := arr[right] - arr[left]
		if r == 0 {
			if data == arr[left] {
				return left
			}
			return -1
		}
		mi = left + (right-left)*(data-arr[left])/r
		if mi < left || mi > right {
			return -1
		}
		miVal := arr[mi]
		if data < miVal {
			right = mi - 1
		} else if data > miVal {
			left = mi + 1
		} else {
			return mi
		}
	}
	return -1
}
