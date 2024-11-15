package main

import tea "github.com/charmbracelet/bubbletea"

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "w", "W":
			m.yStart -= m.height / 20
		case "a", "A":
			m.xStart -= m.width / 20
		case "s", "S":
			m.yStart += m.height / 20
		case "d", "D":
			m.xStart += m.width / 20
		case "i", "I":
			m.zoom *= 2
			m.xStart *= 2
			m.yStart *= 2
			m.resetCache()
		case "o", "O":
			m.zoom /= 2
			m.xStart /= 2
			m.yStart /= 2
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
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.resetCache()
	}

	return m, nil
}
