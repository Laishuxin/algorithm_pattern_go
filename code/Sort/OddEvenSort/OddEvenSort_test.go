package OddEvenSort

import (
	"fmt"
	"testing"
)

func TestOddEvenSort(t *testing.T) {
	//arr := []int{-1, 5, -3, 0}
	arr := []int{-1, 5, -3, 0, 1}
	OddEvenSort(arr)
	fmt.Println(arr)

}