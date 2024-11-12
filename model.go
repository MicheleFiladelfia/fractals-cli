package main

import tea "github.com/charmbracelet/bubbletea"

type State int

const (
	MandelbrotSet State = iota
	JuliaSet
	BurningShipSet
	NumStates // represents the total number of states
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
	fractal_function func(cr float64, ci float64, maxIterations int, bailout float64) int
	palette       []string
	bailout       float64
}

func initialModel() *model {
	return &model{
		zoom:          1.0,
		yStart:        0,
		xStart:        0,
		maxIterations: 20,
		color:         DEF,
		state:         MandelbrotSet,
		fractal_function: mandelbrot,
		palette:       []string{"#421E0F", "#19071A", "#09012F", "#040449", "#000764", "#0C2C8A", "#1852B1", "#397DD1", "#86B5E5", "#D3ECF8", "#F1E9BF", "#F8C95F", "#FFAA00", "#CC8000", "#995700", "#6A3403"},
		bailout:       4.0,
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) nextState() {
	m.state = (m.state + 1) % NumStates
	switch m.state {
	case MandelbrotSet:
		m.fractal_function = mandelbrot
	case JuliaSet:
		m.fractal_function = julia
	case BurningShipSet:
		m.fractal_function = burningShip
	}
}
