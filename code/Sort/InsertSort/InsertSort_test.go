package InsertSort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInsertSort(t *testing.T) {
	//arr := []int{0, 1, 2, 3, -1, 5}
	//arr := []int{1, 5}
	length:= 100
	arr := make([]int, 0, length)
	for i:=0; i < length; i++ {
		arr = append(arr, rand.Intn(length))
	}
	fmt.Println(arr)
	InsertSort(arr)
	fmt.Println(arr)
}