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

func (m *model) computePoint(r, i float64, misses *[]cacheEntry) int {
	if !m.usecache {
		return m.fractalFunction(r, i, m.maxIterations, m.bailout)
	}
	hash := hashParameters(r, i, m.maxIterations, m.bailout)
	if itscached, found := m.cache[hash]; found {
		return itscached
	}
	iterations := m.fractalFunction(r, i, m.maxIterations, m.bailout)
	*misses = append(*misses, cacheEntry{hash: hash, iterations: iterations})
	return iterations
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

	subHeight := m.height * 2

	m.cacheMu.RLock()
	// Generating fractal colored image with double vertical resolution (upper half-block '▀')
	for rowIdx := 0; rowIdx < m.height; rowIdx++ {
		yTop := m.yStart + rowIdx*2
		yBot := yTop + 1
		iTop := (YMIN + (YMAX-YMIN)*(float64(yTop)/float64(subHeight))) / m.zoom
		iBot := (YMIN + (YMAX-YMIN)*(float64(yBot)/float64(subHeight))) / m.zoom

		go func(rowIdx int, iTop, iBot float64) {
			defer wg.Done()

			var row strings.Builder
			row.Grow(m.width * 25)

			var misses []cacheEntry

			for x := m.xStart; x < m.width+m.xStart; x++ {
				r := (XMIN + (XMAX-XMIN)*(float64(x)/float64(m.width))) / m.zoom

				iterTop := m.computePoint(r, iTop, &misses)
				iterBot := m.computePoint(r, iBot, &misses)

				topIdx := m.colorIndex(iterTop)
				botIdx := m.colorIndex(iterBot)

				row.WriteString(m.pairStyles[topIdx][botIdx])
			}

			if len(misses) > 0 {
				rowMisses[rowIdx] = misses
			}

			pixels[rowIdx] = row.String()
		}(rowIdx, iTop, iBot)
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

