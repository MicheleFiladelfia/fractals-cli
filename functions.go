package main
import "math"

func mandelbrot(cr float64, ci float64, maxIterations int, bailout float64) int {
	var zr float64
	var zi float64
	iterations := 0
	var temp float64
	for zr*zr+zi*zi < bailout && iterations < maxIterations {
		temp = zr*zr - zi*zi + cr
		zi = 2.0*zr*zi + ci
		zr = temp
		iterations++
	}

	return iterations
}


func julia(cr float64, ci float64, maxIterations int, bailout float64) int {
	iterations := 0
	var temp float64
	for cr*cr+ci*ci < bailout && iterations < maxIterations {
		temp = cr*cr - ci*ci - 0.7
		ci = 2.0*cr*ci + 0.27015
		cr = temp
		iterations++
	}

	return iterations
}


func burningShip(cr float64, ci float64, maxIterations int, bailout float64) int {
	var (
		zr   float64
		zi   float64
		temp float64
	)
	iterations := 0

	for zr*zr+zi*zi < bailout && iterations < maxIterations {
		temp = zr*zr - zi*zi + cr
		zi = math.Abs(2.0*zr*zi) + ci
		zr = temp
		iterations++
	}

	return iterations
}
