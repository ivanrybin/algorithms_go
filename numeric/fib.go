package numeric

// Fib returns nth Fibonacci number O(1) / O(n) (mem / time).
func Fib(n uint) uint64 {
	if n < 2 {
		return uint64(n)
	}
	prev, curr := uint64(0), uint64(1)
	for i := uint(2); i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
