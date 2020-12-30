package ShellSort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestShellSort(t *testing.T) {
	length:=20
	arr:=make([]int, 0, length)
	for i:=0; i < length; i++ {
		arr =  append(arr, rand.Int()%length)
	}
	fmt.Println(arr)
	ShellSort(arr)
	fmt.Println(arr)
}

func Test_shellSortJumpInsertSort(t *testing.T) {
	length:=6
	arr:=make([]int, 0, length)
	for i:=0; i < length; i++ {
		arr =  append(arr, rand.Int()%length)
	}
	fmt.Println(arr)
	shellSortJumpInsertSort(arr, 0, length, 1)
	fmt.Println(arr)
}