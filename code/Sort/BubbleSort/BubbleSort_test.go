package BubbleSort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 2, 3, 0, -1, 5}
	BubbleSort(arr)
	fmt.Println(arr)
}

func TestQuickBubbleSort(t *testing.T) {
	arr := []int {0, -1, 2, 5, -5, 100, -50, 50, 50, -50}
	QuickBubbleSort(arr)
	fmt.Println(arr)
}