package Sort

import "math/rand"

// Shuffle an array in place
// shuffle algorithm using mathematical method: uniform distribution
func Shuffle(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		r := rand.Int() % i
		temp := arr[i]
		arr[i] = arr[r]
		arr[r] = temp
	}
}