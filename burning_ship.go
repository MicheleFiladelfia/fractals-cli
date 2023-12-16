package main

import "math"

func burningShip(cr float64, ci float64, maxIterations int) int {
	var (
		zr   float64
		zi   float64
		temp float64
	)
	iterations := 0

	for zr*zr+zi*zi < 4 && iterations < maxIterations {
		temp = zr*zr - zi*zi + cr
		zi = math.Abs(2.0*zr*zi) + ci
		zr = temp
		iterations++
	}

	return iterations
}
