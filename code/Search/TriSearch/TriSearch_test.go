package TriSearch

import (
	"fmt"
	"testing"
)

func TestTriSearch(t *testing.T) {
	length := 6
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = i
	}
	fmt.Println("index = ", TriSearch(arr, 0))
	fmt.Println("index = ", TriSearch(arr, 2))
	fmt.Println("index = ", TriSearch(arr, 5))
	fmt.Println("index = ", TriSearch(arr, 50))
	fmt.Println("index = ", TriSearch(arr, 6))
}
