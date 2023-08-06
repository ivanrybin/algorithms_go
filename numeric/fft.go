package numeric

// FFT matrix version of FFT (inefficient by memory due a lot of duplicated subarrays).
func FFT(a []complex128, w complex128) []complex128 {
	if len(a) == 1 {
		return a
	}
	n := len(a)
	a0 := make([]complex128, n/2)
	a1 := make([]complex128, n/2)
	for i := 0; i < n; i += 2 {
		a0[i/2], a1[i/2] = a[i], a[i+1]
	}
	s0 := FFT(a0, w*w)
	s1 := FFT(a1, w*w)
	r := make([]complex128, n)
	wi := complex(1, 0)
	for i := 0; i < n/2; i++ {
		r[i] = s0[i] + wi*s1[i]
		r[i+n/2] = s0[i] - wi*s1[i]
		wi *= w
	}
	return r
}
