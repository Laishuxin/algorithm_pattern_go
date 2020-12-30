package InsertSort

import "code/Sort"

func InsertSort(arr []int) error {
	if arr == nil {
		return Sort.ErrArrIsNil
	}
	size := len(arr)
	for i := 1; i < size; i++ {
		var j int
		temp := arr[i]
		for j = i - 1; j >= 0 && arr[j] < temp; j-- {
			//Sort.Swap(arr, j, j+1)
			arr[j+1] = arr[j]
		}
		arr[j + 1] = temp
	}
	return nil
}
