package ShellSort

import "code/Sort"

func ShellSort(arr []int) error {
	if arr == nil {
		return Sort.ErrArrIsNil
	}
	if len(arr) <= 1 {
		return nil
	}
	shellSort(arr)
	return nil
}

func shellSort(arr []int) {
	length := len(arr)
	for step:=length; step >= 1; step /= 2 {
		// 对步长轴每一个元素进行处理
		for i := 0; i < step; i++ {
			shellSortJumpInsertSort(arr, i, length, step)
		}
	}
}

func shellSortJumpInsertSort(arr []int, lo int, hi int, step int) {
	for i := lo + step; i < hi; i += step {
		var j int
		temp := arr[i]
		for j = i - step; j >= lo && arr[j] < temp; j -= step {
			arr[j+step] = arr[j]
		}
		arr[j+step] = temp
	}
}
