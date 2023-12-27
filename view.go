package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/charmbracelet/lipgloss"
)

func (m *model) View() string {
	pixels := make([]string, 0)

	var mu sync.Mutex
	var wg sync.WaitGroup

	//generating fractal colored "image" by using concurrency
	for y := m.yStart; y < m.height+m.yStart; y++ {

		//formula for imaginary part
		i := (YMIN + (YMAX-YMIN)*((float64(y))/float64(m.height))) / m.zoom

		wg.Add(1)
		pixels = append(pixels, "")

		go func(y int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			for x := m.xStart; x < m.width+m.xStart; x++ {
				//formula for real part
				r := (XMIN + (XMAX-XMIN)*((float64(x))/float64(m.width))) / m.zoom
				var iterations int

				if m.state == MandelbrotSet {
					iterations = mandelbrot(r, i, m.maxIterations)
				} else if m.state == JuliaSet {
					// Use a constant value for the Julia set
					juliaC := complex(-0.7, 0.27015)
					iterations = julia(r, i, m.maxIterations, juliaC)
				} else {
					iterations = burningShip(r, i, m.maxIterations)
				}

				if iterations == m.maxIterations {
					//if current pixel is the set
					pixels[y-m.yStart] += ((lipgloss.NewStyle().SetString(" ").Background(lipgloss.Color("#000000"))).String())
				} else {
					//else set proper color to current pixel based on the number of iteration
					pixels[y-m.yStart] += (lipgloss.NewStyle().SetString(" ").Background(lipgloss.Color(m.palette[iterations%len(m.palette)]))).String()
				}
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
