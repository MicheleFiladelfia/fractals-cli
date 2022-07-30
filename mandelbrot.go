package main

import "math/cmplx"

const (
	xMax = 2.4
	yMax = 1.6
	xMin = -2.4
	yMin = -1.6
)

func mandelbrot(c complex128, maxIteration int) int {
	var z complex128
	var it int

	for cmplx.Abs(z) <= 2 && it < maxIteration {
		z = z*z + c
		it++
	}

	return it
}
