package main

const (
	XMAX = 2.4
	YMAX = 1.6
	XMIN = -2.4
	YMIN = -1.6
)

func mandelbrot(cr float64, ci float64 , maxIterations int) int {
  var zr float64
  var zi float64
  iterations:=0
  var temp float64
  for zr * zr + zi * zi < 4 && iterations < maxIterations {
    temp = zr * zr - zi * zi + cr
    zi = 2.0 * zr * zi + ci
    zr = temp
    iterations++
  }

  return iterations
}
