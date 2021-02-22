package MergeSort

const cutoff = 20

// Bottom up merge sort replaces recursion with iteration.
func BottomUpMergeSort(arr []int) {
	if arr == nil {
		return
	}
	bottomUpMergeSort(arr, 0, len(arr))
}

func bottomUpMergeInsertSort(arr []int, lo int, hi int) {
	for i := lo + 1; i < hi; i++ {
		for j := i; j > lo && arr[j] < arr[j-1]; j-- {
			swap(arr, j, j-1)
		}
	}
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func bottomUpMergeSort(arr []int, lo int, hi int) {
	if hi-lo <= 1 {
		return
	} else if hi-lo <= cutoff {
		bottomUpMergeInsertSort(arr, lo, hi)
		return
	}

	length := hi - lo

	for i := 1; i < length; i <<= 1 {
		for j := lo; j < hi-i; j += i + i {
			bottomUpMerge(arr, j, j+i, min(j+i+i, length))
		}
	}
}

func bottomUpMerge(arr []int, lo int, mi int, hi int) {
	lLength := mi - lo
	lPart := make([]int, lLength, lLength)
	for i := 0; i < lLength; i++ {
		lPart[i] = arr[lo+i]
	}
	l, r, c := 0, mi, lo
	for ; l < lLength && r < hi; c++ {
		if lPart[l] <= arr[r] {
			arr[c] = lPart[l]
			l++
		} else {
			arr[c] = arr[r]
			r++
		}
	}
	// merge equal
	for ; c < hi && l < lLength; c++ {
		arr[c] = lPart[l]
		l++
	}
}
