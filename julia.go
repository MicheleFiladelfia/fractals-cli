package main

func julia(cr float64, ci float64, maxIterations int, bailout float64, juliaC complex128) int {
	iterations := 0
	var temp float64
	for cr*cr+ci*ci < bailout && iterations < maxIterations {
		temp = cr*cr - ci*ci + real(juliaC)
		ci = 2.0*cr*ci + imag(juliaC)
		cr = temp
		iterations++
	}

	return iterations
}
