package SelectSort

import (
	"fmt"
	"testing"
)

func cmpLess(i1 *interface{}, i2 *interface{}) bool {
	n1 := (*i1).(int)
	n2 := (*i2).(int)
	return n1 < n2
}

func TestSelectSort(t *testing.T) {
	arr := []interface{}{1, 3, 2, 0, -1, 5, 4}
	err := SelectSort(arr, cmpLess)
	if (err == nil) {
		fmt.Println(arr)
	} else {
		fmt.Println("there are some error: ", err)
	}
}
