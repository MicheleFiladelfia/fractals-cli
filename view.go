package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/charmbracelet/lipgloss"
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

func (m *model) renderCoordinatesBar() string {
	coordText := fmt.Sprintf(" %s │ Pos: %.6f, %.6fi │ Zoom: %.1fx │ Iter: %d │ Bailout: %.1f │ Palette: %s │ [?] Help ",
		m.stateName(), m.centerX, m.centerY, m.zoom, m.maxIterations, m.bailout, m.paletteName())

	barStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#202020")).
		MaxWidth(m.width)

	return barStyle.Render(coordText)
}

func (m *model) renderHelpOverlay() string {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFD700")).MarginBottom(1)
	keyStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FFFF"))
	descStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#E0E0E0"))

	lines := []string{
		titleStyle.Render("=== FRACTALS-CLI CONTROLS ==="),
		fmt.Sprintf("%s  %s", keyStyle.Render("W, A, S, D "), descStyle.Render("Move camera")),
		fmt.Sprintf("%s  %s", keyStyle.Render("I / O      "), descStyle.Render("Zoom in / out")),
		fmt.Sprintf("%s  %s", keyStyle.Render("+ / -      "), descStyle.Render("Iterations (+/- 5)")),
		fmt.Sprintf("%s  %s", keyStyle.Render("B / V      "), descStyle.Render("Bailout (+/- 0.1)")),
		fmt.Sprintf("%s  %s", keyStyle.Render("C          "), descStyle.Render("Next color palette")),
		fmt.Sprintf("%s  %s", keyStyle.Render("N          "), descStyle.Render("Next fractal set")),
		fmt.Sprintf("%s  %s", keyStyle.Render("X          "), descStyle.Render("Toggle coordinates bar")),
		fmt.Sprintf("%s  %s", keyStyle.Render("? / H      "), descStyle.Render("Toggle this help")),
		fmt.Sprintf("%s  %s", keyStyle.Render("ctrl+c     "), descStyle.Render("Quit")),
		"",
		descStyle.Render("Press ? or Esc to close"),
	}

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#888888")).
		Background(lipgloss.Color("#181818")).
		Padding(1, 2)

	return boxStyle.Render(strings.Join(lines, "\n"))
}

func (m *model) View() string {
	if m.height <= 0 || m.width <= 0 {
		return ""
	}

	renderHeight := m.height
	if m.showCoordinates && renderHeight > 1 {
		renderHeight--
	}

	pixels := make([]string, renderHeight)
	var rowMisses [][]cacheEntry
	if m.usecache {
		rowMisses = make([][]cacheEntry, renderHeight)
	}

	var wg sync.WaitGroup
	wg.Add(renderHeight)

	subHeight := renderHeight * 2

	m.cacheMu.RLock()
	// Generating fractal colored image with double vertical resolution (upper half-block '▀')
	for rowIdx := 0; rowIdx < renderHeight; rowIdx++ {
		yTop := rowIdx * 2
		yBot := yTop + 1
		iTop := m.centerY + ((float64(yTop)-float64(subHeight)/2.0)/(m.zoom*float64(subHeight)))*(YMAX-YMIN)
		iBot := m.centerY + ((float64(yBot)-float64(subHeight)/2.0)/(m.zoom*float64(subHeight)))*(YMAX-YMIN)

		go func(rowIdx int, iTop, iBot float64) {
			defer wg.Done()

			var row strings.Builder
			row.Grow(m.width * 25)

			var misses []cacheEntry

			for px := 0; px < m.width; px++ {
				r := m.centerX + ((float64(px)-float64(m.width)/2.0)/(m.zoom*float64(m.width)))*(XMAX-XMIN)

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

	output := strings.Join(pixels, "\n")
	if m.showCoordinates {
		output += "\n" + m.renderCoordinatesBar()
	}

	if m.showHelp {
		overlay := m.renderHelpOverlay()
		output = lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, overlay,
			lipgloss.WithWhitespaceChars(" "),
			lipgloss.WithWhitespaceForeground(lipgloss.NoColor{}),
		)
	}

	return output
}


