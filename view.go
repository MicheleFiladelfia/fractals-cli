package main

import (
	"strings"
	"sync"
)

const (
	XMAX = 2.4
	YMAX = 1.6
	XMIN = -2.4
	YMIN = -1.6
)

type cacheEntry struct {
	hash       uint64
	iterations int
}

func (m *model) View() string {
	if m.height <= 0 || m.width <= 0 {
		return ""
	}

	pixels := make([]string, m.height)
	var rowMisses [][]cacheEntry
	if m.usecache {
		rowMisses = make([][]cacheEntry, m.height)
	}

	var wg sync.WaitGroup
	wg.Add(m.height)

	paletteLen := len(m.renderedPalette)

	m.cacheMu.RLock()
	// Generating fractal colored "image" by using concurrency (one goroutine per row)
	for y := m.yStart; y < m.height+m.yStart; y++ {
		// Formula for imaginary part
		i := (YMIN + (YMAX-YMIN)*(float64(y)/float64(m.height))) / m.zoom
		rowIdx := y - m.yStart

		go func(rowIdx int, i float64) {
			defer wg.Done()

			var row strings.Builder
			row.Grow(m.width * 20)

			var misses []cacheEntry

			for x := m.xStart; x < m.width+m.xStart; x++ {
				// Formula for real part
				r := (XMIN + (XMAX-XMIN)*(float64(x)/float64(m.width))) / m.zoom

				var iterations int
				if m.usecache {
					hash := hashParameters(r, i, m.maxIterations, m.bailout)
					if itscached, found := m.cache[hash]; found {
						iterations = itscached
					} else {
						iterations = m.fractalFunction(r, i, m.maxIterations, m.bailout)
						misses = append(misses, cacheEntry{hash: hash, iterations: iterations})
					}
				} else {
					iterations = m.fractalFunction(r, i, m.maxIterations, m.bailout)
				}

				if iterations == m.maxIterations {
					row.WriteString(m.setPixel)
				} else {
					row.WriteString(m.renderedPalette[iterations%paletteLen])
				}
			}

			if len(misses) > 0 {
				rowMisses[rowIdx] = misses
			}

			pixels[rowIdx] = row.String()
		}(rowIdx, i)
	}

	wg.Wait()
	m.cacheMu.RUnlock()

	if m.usecache {
		m.cacheMu.Lock()
		for _, misses := range rowMisses {
			for _, entry := range misses {
				m.cache[entry.hash] = entry.iterations
			}
		}
		m.cacheMu.Unlock()
	}

	return strings.Join(pixels, "\n")
}

