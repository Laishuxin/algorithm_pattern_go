package TriSearch

// TriSearch 三分查找
func TriSearch(arr []int, target int) int {
	if arr == nil {
		return -1
	}
	left := 0
	right := len(arr) - 1
	mi1 := 0
	mi2 := 0
	for left <= right {
		mi1 = left + int((right-left)/3)
		mi2 = right - int((right-left)/3)
		mi1Val := arr[mi1]
		mi2Val := arr[mi2]
		if target == mi1Val {
			return mi1
		}
		if target == mi2Val {
			return mi2
		}
		if target < mi1Val {
			right = mi1 - 1
		} else if mi2Val < target {
			left = mi2 + 1
		} else {
			left = mi1 + 1
			right = mi2 - 1
		}
	}
	return -1
}
