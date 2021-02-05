package CombSort

import "code/Sort"

// 梳子排序：冒泡排序 + 希尔排序的思想
func CombSort(arr []int) error {
	if arr == nil {
		return Sort.ErrArrIsNil
	}
	length := len(arr)
	gap := length
	const SHRINK = 0.8
	isSorted := false
	for gap > 1 || !isSorted {
		gap = int(float64(gap) * SHRINK) // gap = gap / 1.3
		isSorted = true
		for i := 0; i+gap < length; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				isSorted = false
			}
		}
	}
	return nil
}
