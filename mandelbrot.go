package main

import "math/cmplx"

const (
	XMAX = 2.4
	YMAX = 1.6
	XMIN = -2.4
	YMIN = -1.6
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
