package FibonacciSearch

type fibonacci struct {
	prev int
	curr int
}

// [1, 1, 2, 3, 5, 8, 13]
func newFibonacci(n int) *fibonacci {
	fib := new(fibonacci)
	fib.curr = 1
	// fib.prev = 0
	for fib.curr < n {
		fib.next()
	}
	return fib
}

func (fib *fibonacci) get() int {
	return fib.curr
}

func (fib *fibonacci) back() int {
	fib.prev = fib.curr - fib.prev
	fib.curr = fib.curr - fib.prev
	return fib.curr
}

func (fib *fibonacci) next() int {
	fib.curr = fib.curr + fib.prev
	fib.prev = fib.curr - fib.prev
	return fib.curr
}

// FibonacciSearch search target via fibonacci search.
// suppose sortedArr is increasing.
func FibonacciSearch(sortedArr []int, target int) int {
	// ----------------------------------------------------------------
	// boundary value judgement
	if sortedArr == nil {
		return -1
	}
	lo := 0
	hi := len(sortedArr)
	if target < sortedArr[lo] || target > sortedArr[hi-1] {
		return -1
	}
	// ----------------------------------------------------------------
	fib := newFibonacci(hi - lo)
	for lo < hi {
		for hi-lo < fib.get() {
			fib.back()
		}
		mi := lo + fib.get() - 1
		miVal := sortedArr[mi]
		if target < miVal {
			hi = mi
		} else if miVal < target {
			lo = mi + 1
		} else {
			return mi
		}
	}

	return -1
}
