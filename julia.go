package main

func juliaSet(cr float64, ci float64, maxIterations int, juliaC complex128) int {
	iterations := 0
	var temp float64
	for cr*cr+ci*ci < 4 && iterations < maxIterations {
		temp = cr*cr - ci*ci + real(juliaC)
		ci = 2.0*cr*ci + imag(juliaC)
		cr = temp
		iterations++
	}

	return iterations
}
