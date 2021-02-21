package ShellSort

func swap(arr []int, n int, m int) {
	temp := arr[n]
	arr[n] = arr[m]
	arr[m] = temp
}
func ShellSort(arr []int) {
	if arr == nil {
		return
	}
	length := len(arr)
	h := 1
	for h < length/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < length; i += h {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				swap(arr, j, j-h)
			}
		}
		h = h / 3
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
