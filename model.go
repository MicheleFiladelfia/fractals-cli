package main

import (
	"sync"

	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	MandelbrotSet State = iota
	JuliaSet
	BurningShipSet
	TricornSet
	CelticSet
	BuffaloSet
	Multibrot3Set
	NewtonSet
	PerpBurningShipSet
	CelticBurningShipSet
	Multibrot4Set
	LambdaSet
	SharkFinSet
	NumStates // represents the total number of states
)

type model struct {
	height          int
	width           int
	zoom            float64
	centerX         float64
	centerY         float64
	maxIterations   int
	color           int
	state           State
	fractalFunction func(cr float64, ci float64, maxIterations int, bailout float64) int
	palette         []string
	pairStyles      [][]string
	bailout         float64
	usecache        bool
	cache           map[uint64]int
	cacheMu         sync.RWMutex
	showCoordinates bool
	showHelp        bool
}

func initialModel() *model {
	m := &model{
		zoom:            1.0,
		centerX:         0.0,
		centerY:         0.0,
		maxIterations:   20,
		color:           DEF,
		state:           MandelbrotSet,
		fractalFunction: mandelbrot,
		palette:         []string{"#421E0F", "#19071A", "#09012F", "#040449", "#000764", "#0C2C8A", "#1852B1", "#397DD1", "#86B5E5", "#D3ECF8", "#F1E9BF", "#F8C95F", "#FFAA00", "#CC8000", "#995700", "#6A3403"},
		bailout:         4.0,
		usecache:        true,
		cache:           make(map[uint64]int),
	}
	m.buildPaletteStyles()
	return m
}

func (m *model) Init() tea.Cmd {
	m.resetCache()
	return nil
}

func (m *model) stateName() string {
	switch m.state {
	case MandelbrotSet:
		return "Mandelbrot"
	case JuliaSet:
		return "Julia"
	case BurningShipSet:
		return "Burning Ship"
	case TricornSet:
		return "Tricorn"
	case CelticSet:
		return "Celtic"
	case BuffaloSet:
		return "Buffalo"
	case Multibrot3Set:
		return "Multibrot (z^3)"
	case NewtonSet:
		return "Newton (z^3-1)"
	case PerpBurningShipSet:
		return "Perpendicular Burning Ship"
	case CelticBurningShipSet:
		return "Celtic Burning Ship"
	case Multibrot4Set:
		return "Multibrot (z^4)"
	case LambdaSet:
		return "Lambda"
	case SharkFinSet:
		return "Shark Fin"
	default:
		return "Unknown"
	}
}

func (m *model) paletteName() string {
	switch m.color {
	case DEF:
		return "Default"
	case PURPLE:
		return "Purple"
	case YELLOW:
		return "Yellow"
	case GREY:
		return "Grey"
	case FIRE:
		return "Fire"
	case BLUE:
		return "Blue"
	case RAINBOW:
		return "Rainbow"
	case CYBERPUNK:
		return "Cyberpunk"
	case EMERALD:
		return "Emerald"
	case SUNSET:
		return "Sunset"
	case ELECTRIC_BLUE:
		return "Electric Blue"
	case NEON_ACID:
		return "Neon Acid"
	default:
		return "Custom"
	}
}

func (m *model) nextState() {
	m.state = (m.state + 1) % NumStates
	switch m.state {
	case MandelbrotSet:
		m.fractalFunction = mandelbrot
	case JuliaSet:
		m.fractalFunction = julia
	case BurningShipSet:
		m.fractalFunction = burningShip
	case TricornSet:
		m.fractalFunction = tricorn
	case CelticSet:
		m.fractalFunction = celticMandelbrot
	case BuffaloSet:
		m.fractalFunction = buffalo
	case Multibrot3Set:
		m.fractalFunction = multibrot3
	case NewtonSet:
		m.fractalFunction = newton
	case PerpBurningShipSet:
		m.fractalFunction = perpendicularBurningShip
	case CelticBurningShipSet:
		m.fractalFunction = celticBurningShip
	case Multibrot4Set:
		m.fractalFunction = multibrot4
	case LambdaSet:
		m.fractalFunction = lambda
	case SharkFinSet:
		m.fractalFunction = sharkFin
	}
}

func (m *model) resetCache() {
	m.cacheMu.Lock()
	defer m.cacheMu.Unlock()
	m.cache = make(map[uint64]int)
}
