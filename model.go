package main

import tea "github.com/charmbracelet/bubbletea"

type State int

const (
	MandelbrotSet State = iota
	JuliaSet
	BurningShip
	numStates // represents the total number of states
)

type model struct {
	height        int
	width         int
	zoom          float64
	yStart        int
	xStart        int
	maxIterations int
	color         int
	state         State
	palette       []string
}

func initialModel() *model {
	return &model{
		zoom:          1.0,
		yStart:        0,
		xStart:        0,
		maxIterations: 20,
		color:         DEF,
		state:         MandelbrotSet,
		palette:       []string{"#421E0F", "#19071A", "#09012F", "#040449", "#000764", "#0C2C8A", "#1852B1", "#397DD1", "#86B5E5", "#D3ECF8", "#F1E9BF", "#F8C95F", "#FFAA00", "#CC8000", "#995700", "#6A3403"},
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) nextState() {
	m.state = (m.state + 1) % numStates
}
