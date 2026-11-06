package main

import tea "github.com/charmbracelet/bubbletea"

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Step size per movement based on current zoom and screen dimensions
		subHeight := m.height * 2
		if subHeight <= 0 {
			subHeight = 1
		}
		width := m.width
		if width <= 0 {
			width = 1
		}
		stepY := ((YMAX - YMIN) / (m.zoom * float64(subHeight))) * float64(subHeight/20+1)
		stepX := ((XMAX - XMIN) / (m.zoom * float64(width))) * float64(width/20+1)

		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "w", "W":
			m.centerY -= stepY
		case "a", "A":
			m.centerX -= stepX
		case "s", "S":
			m.centerY += stepY
		case "d", "D":
			m.centerX += stepX
		case "i", "I":
			m.zoom *= 1.5
			m.resetCache()
		case "o", "O":
			m.zoom /= 1.5
			m.resetCache()
		case "c", "C":
			m.nextPalette()
		case "n", "N":
			m.nextState()
			m.resetCache()
		case "+":
			m.maxIterations += 5
			m.resetCache()
		case "-":
			if m.maxIterations > 5 {
				m.maxIterations -= 5
				m.resetCache()
			}
		case "b", "B":
			m.bailout += 0.1
			m.resetCache()
		case "v", "V":
			if m.bailout > 0.1 {
				m.bailout -= 0.1
				m.resetCache()
			}
		case "x", "X":
			m.showCoordinates = !m.showCoordinates
		case "?", "h", "H":
			m.showHelp = !m.showHelp
		case "esc":
			if m.showHelp {
				m.showHelp = false
			}
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.resetCache()
	}

	return m, nil
}
