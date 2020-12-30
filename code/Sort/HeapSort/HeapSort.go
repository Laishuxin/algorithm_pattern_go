package HeapSort

import "code/Sort"

func HeapSort(arr []int) error {
	if arr == nil {
		return Sort.ErrArrIsNil
	}
	length := len(arr)
	if length <= 1 {
		return nil
	}
	// 构建最小堆
	for i := (length - 2) / 2; i >= 0; i-- {
		_downAdjust(arr, i, length)
	}

	for i := length - 1; i > 0; i-- {
		Sort.Swap(arr, 0, i)
		_downAdjust(arr, 0, i)
	}
	return nil
}

func _downAdjust(arr []int, parentIndex int, length int) {
	temp := arr[parentIndex]
	childIndex := 2*parentIndex + 1
	for childIndex < length {
		if childIndex+1 < length && arr[childIndex+1] < arr[childIndex] {
			childIndex++
		}
		if temp <= arr[childIndex] {
			break
		}
		arr[parentIndex] = arr[childIndex]
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}
	arr[parentIndex] = temp
}
