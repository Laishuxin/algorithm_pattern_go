package FibonacciSearch

import (
	"fmt"
	"testing"
)

func Test_newFibonacci(t *testing.T) {
	newFibonacci(10)
}

func TestFibonacciSearch(t *testing.T) {
	// arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	arr := []int{0}

	var index int
	index = FibonacciSearch(arr, 5)
	fmt.Println("index = ", index)
	index = FibonacciSearch(arr, 0)
	fmt.Println("index = ", index)
	index = FibonacciSearch(arr, 9)
	fmt.Println("index = ", index)
}
