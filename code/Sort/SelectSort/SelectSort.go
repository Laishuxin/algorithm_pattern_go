package SelectSort

import "errors"

//func swap(i1 *interface{}, i2 *interface{}) {
//	var temp *interface{}
//	temp = i1
//	i1 = i2
//	i2 = temp
//}
func swap(arr []interface{}, index1 int, index2 int) {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
}

func SelectSort(arr []interface{}, cmpLess func(*interface{}, *interface{}) bool) error {
	if arr == nil {
		return errors.New("select sort: nil arr")
	}
	size := len(arr)
	for i := 0; i < size; i++ {
		maxIndex := i
		for j := i + 1; j < size; j++ {
			if cmpLess(&arr[maxIndex], &arr[j]) {
				maxIndex = j
			}
		}
		if maxIndex != i {
			swap(arr, maxIndex, i)
		}
	}
	return nil
}
