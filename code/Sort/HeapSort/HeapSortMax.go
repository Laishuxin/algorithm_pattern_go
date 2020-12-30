package HeapSort

import "code/Sort"

func HeapSortMax(arr []int) error {
	if arr == nil {
		return Sort.ErrArrIsNil
	}
	length := len(arr)
	if length <= 1 {
		return nil
	}

	// build heap
	for i := (length - 2) / 2; i >= 0; i-- {
		downAdjust(arr, i, length)
	}

	for i := length - 1; i > 0; i-- {
		Sort.Swap(arr, 0, i)
		downAdjust(arr, 0, i)
	}
	return nil
}

// 最大堆辅助函数
func downAdjust(arr []int, parentIndex int, length int) {
	temp := arr[parentIndex]
	childIndex := 2*parentIndex + 1
	for childIndex < length {
		if childIndex+1 < length && arr[childIndex] < arr[childIndex+1] {
			childIndex++
		}
		if temp >= arr[childIndex] {
			break
		}
		arr[parentIndex] = arr[childIndex]
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}
	arr[parentIndex] = temp
}
