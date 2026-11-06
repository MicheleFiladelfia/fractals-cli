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

func tricorn(cr float64, ci float64, maxIterations int, bailout float64) int {
	var zr float64
	var zi float64
	iterations := 0
	var temp float64
	for zr*zr+zi*zi < bailout && iterations < maxIterations {
		temp = zr*zr - zi*zi + cr
		zi = -2.0*zr*zi + ci // Complex conjugate squared
		zr = temp
		iterations++
	}

	return iterations
}

func celticMandelbrot(cr float64, ci float64, maxIterations int, bailout float64) int {
	var zr float64
	var zi float64
	iterations := 0
	var temp float64
	for zr*zr+zi*zi < bailout && iterations < maxIterations {
		temp = math.Abs(zr*zr-zi*zi) + cr
		zi = 2.0*zr*zi + ci
		zr = temp
		iterations++
	}

	return iterations
}

func buffalo(cr float64, ci float64, maxIterations int, bailout float64) int {
	var zr float64
	var zi float64
	iterations := 0
	var temp float64
	for zr*zr+zi*zi < bailout && iterations < maxIterations {
		temp = math.Abs(zr*zr-zi*zi) + cr
		zi = -2.0*math.Abs(zr)*zi + ci
		zr = temp
		iterations++
	}

	return iterations
}

func multibrot3(cr float64, ci float64, maxIterations int, bailout float64) int {
	var zr float64
	var zi float64
	iterations := 0
	var temp float64
	for zr*zr+zi*zi < bailout && iterations < maxIterations {
		temp = zr*(zr*zr-3.0*zi*zi) + cr
		zi = zi*(3.0*zr*zr-zi*zi) + ci
		zr = temp
		iterations++
	}

	return iterations
}

func newton(cr float64, ci float64, maxIterations int, bailout float64) int {
	zr := cr
	zi := ci
	tolerance := 1e-4

	for iter := 0; iter < maxIterations; iter++ {
		// z^3 - 1 with Newton-Raphson: z_{n+1} = (2*z + 1/z^2) / 3
		// z^2 = (zr^2 - zi^2) + i*(2*zr*zi)
		r2 := zr*zr - zi*zi
		i2 := 2.0 * zr * zi
		magSq2 := r2*r2 + i2*i2
		if magSq2 == 0 {
			break
		}

		// 1 / z^2 = (r2 - i*i2) / magSq2
		invR := r2 / magSq2
		invI := -i2 / magSq2

		// next z = (2/3)*z + (1/3)*(1/z^2)
		nextR := (2.0*zr + invR) / 3.0
		nextI := (2.0*zi + invI) / 3.0

		diffR := nextR - zr
		diffI := nextI - zi
		if diffR*diffR+diffI*diffI < tolerance {
			return iter
		}

		zr = nextR
		zi = nextI
	}

	return maxIterations
}

func perpendicularBurningShip(cr float64, ci float64, maxIterations int, bailout float64) int {
	var zr, zi, temp float64
	iterations := 0
	for zr*zr+zi*zi < bailout && iterations < maxIterations {
		temp = zr*zr - zi*zi + cr
		zi = -2.0*zr*math.Abs(zi) + ci
		zr = temp
		iterations++
	}
	return iterations
}

func celticBurningShip(cr float64, ci float64, maxIterations int, bailout float64) int {
	var zr, zi, temp float64
	iterations := 0
	for zr*zr+zi*zi < bailout && iterations < maxIterations {
		temp = math.Abs(zr*zr-zi*zi) + cr
		zi = 2.0*math.Abs(zr*zi) + ci
		zr = temp
		iterations++
	}
	return iterations
}

func multibrot4(cr float64, ci float64, maxIterations int, bailout float64) int {
	var zr, zi, temp float64
	iterations := 0
	for zr*zr+zi*zi < bailout && iterations < maxIterations {
		r2 := zr * zr
		i2 := zi * zi
		temp = r2*r2 - 6.0*r2*i2 + i2*i2 + cr
		zi = 4.0*zr*zi*(r2-i2) + ci
		zr = temp
		iterations++
	}
	return iterations
}

func lambda(cr float64, ci float64, maxIterations int, bailout float64) int {
	zr := 0.5
	zi := 0.0
	iterations := 0
	// z = c * z * (1 - z)
	for zr*zr+zi*zi < bailout && iterations < maxIterations {
		// w = 1 - z = (1 - zr) - i*zi
		omzr := 1.0 - zr
		omzi := -zi

		// p = z * (1 - z)
		pr := zr*omzr - zi*omzi
		pi := zr*omzi + zi*omzr

		// z = c * p
		nextR := cr*pr - ci*pi
		nextI := cr*pi + ci*pr

		zr = nextR
		zi = nextI
		iterations++
	}
	return iterations
}

func sharkFin(cr float64, ci float64, maxIterations int, bailout float64) int {
	var zr, zi, temp float64
	iterations := 0
	for zr*zr+zi*zi < bailout && iterations < maxIterations {
		temp = zr*zr - zi*zi + cr
		zi = 2.0*math.Abs(zr)*math.Abs(zi) + ci
		zr = temp
		iterations++
	}
	return iterations
}



