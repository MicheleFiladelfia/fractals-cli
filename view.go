package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/charmbracelet/lipgloss"
)

func (m *model) View() string {
	pixels := make([]string, 0)

	var wg sync.WaitGroup

	//generating fractal colored "image" by using concurrency
	for y := m.yStart; y < m.height+m.yStart; y++ {

		//formula for imaginary part
		i := (YMIN + (YMAX-YMIN)*((float64(y))/float64(m.height))) / m.zoom

		wg.Add(1)
		pixels = append(pixels, "")

		go func(y int) {
			defer wg.Done()
			switch m.state {
			case MandelbrotSet:
				calculateMandelbrot(m, pixels, y, i)
			case JuliaSet:
				calculateJuliaSet(m, pixels, y, i)
			}
		}(y)

	}
	wg.Wait()

	for i := 0; i < m.height; i++ {
		if runtime.GOOS != "windows" && i != 0 {
			fmt.Print("\n")
		}

		fmt.Print(pixels[i])
	}

	fmt.Print("\x1b[H")

	return ""
}

func calculateMandelbrot(m *model, pixels []string, y int, imag float64) {
	for x := m.xStart; x < m.width+m.xStart; x++ {
		//formula for real part
		r := (XMIN + (XMAX-XMIN)*((float64(x))/float64(m.width))) / m.zoom

		iterations := mandelbrot(r, imag, m.maxIterations)

		if iterations == m.maxIterations {
			//if current pixel is the set
			pixels[y-m.yStart] += ((lipgloss.NewStyle().SetString(" ").Background(lipgloss.Color("#000000"))).String())
		} else {
			//else set proper color to current pixel based on the number of iteration
			pixels[y-m.yStart] += (lipgloss.NewStyle().SetString(" ").Background(lipgloss.Color(m.palette[iterations%len(m.palette)]))).String()
		}
	}
}

func calculateJuliaSet(m *model, pixels []string, y int, imag float64) {
	for x := m.xStart; x < m.width+m.xStart; x++ {
		// Formula for real part
		r := (XMIN + (XMAX-XMIN)*((float64(x))/float64(m.width))) / m.zoom

		// Use a constant value for the Julia set
		juliaC := complex(-0.7, 0.27015)

		iterations := juliaSet(r, imag, m.maxIterations, juliaC)

		if iterations == m.maxIterations {
			// If current pixel is in the set
			pixels[y-m.yStart] += ((lipgloss.NewStyle().SetString(" ").Background(lipgloss.Color("#000000"))).String())
		} else {
			// Set proper color to current pixel based on the number of iterations
			pixels[y-m.yStart] += (lipgloss.NewStyle().SetString(" ").Background(lipgloss.Color(m.palette[iterations%len(m.palette)]))).String()
		}
	}
}
