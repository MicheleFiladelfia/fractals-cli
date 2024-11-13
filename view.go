package main

import (
	"fmt"
	"runtime"
	"sync"
	"bufio"
	"os"

	"github.com/charmbracelet/lipgloss"
)

const (
	XMAX = 2.4
	YMAX = 1.6
	XMIN = -2.4
	YMIN = -1.6
)

func (m *model) View() string {
	pixels := make([]string, 0)

	buf := bufio.NewWriter(os.Stdout)
	defer buf.Flush()

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
				
				hash := hashParameters(r,i,m.maxIterations,m.bailout)
				var iterations int

				if m.usecache {
					itscached, found := m.cache[hash]

					if found == false {
						iterations = m.fractalFunction(r, i, m.maxIterations, m.bailout)
						m.cache[hash] = iterations
					} else {
						iterations = itscached
					}
				} else {
					iterations = m.fractalFunction(r, i, m.maxIterations, m.bailout)
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
			fmt.Fprint(buf,"\n")
		}

		fmt.Fprint(buf,pixels[i])
	}

	fmt.Fprint(buf,"\x1b[H")

	return ""
}
