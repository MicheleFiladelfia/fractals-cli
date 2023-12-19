package main

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version":
			fmt.Println("mandelbrot-cli (v1.0.0), created by @MicheleFiladelfia")
		case "--help":
			fmt.Println("mandelbrot-cli: a simple mandelbrot set terminal explorer.")
			fmt.Println("Usage: mandelbrot-cli [OPTION]")
			fmt.Println()
			fmt.Println("Options:")
			fmt.Println("  --version\t\tPrint version information and exit.")
			fmt.Println("  --help\t\tPrint this help message and exit.")
			fmt.Println()
			fmt.Println("Controls:")
			fmt.Println("  W\t\t\tMove up.")
			fmt.Println("  S\t\t\tMove down.")
			fmt.Println("  A\t\t\tMove left.")
			fmt.Println("  D\t\t\tMove right.")
			fmt.Println("  I\t\t\tZoom in.")
			fmt.Println("  O\t\t\tZoom out.")
			fmt.Println("  +\t\t\tIncrease iterations.")
			fmt.Println("  -\t\t\tDecrease iterations.")
			fmt.Println("  C\t\t\tChange color palette.")
			fmt.Println("  N\t\t\tChange fractal set.")
			fmt.Println("  ctrl+c\t\tQuit.")
			fmt.Println()
			fmt.Println("Report bugs to <https://github.com/MicheleFiladelfia>")
		case "--fractal":
			if len(os.Args) > 2 {
				choice, err := strconv.Atoi(os.Args[2])
				if err != nil {
					fmt.Println("mandelbrot-cli: invalid fractal set option")
					fmt.Println("Please make sure the provided option is a valid integer.")
					fmt.Println("Try '--fractal' for more information.")
					os.Exit(1)
				}
				if State(choice) < NumStates {
					m := initialModel()
					m.state = State(choice)
					startWithModel(m)
				} else {
					var fractals = map[int]string{
						0: "Mandelbrot Set",
						1: "Julia Set",
						2: "Burning Ship",
					}
					fmt.Printf("mandelbrot-cli: unrecognized fractal set option '%d'\n", choice)
					fmt.Println("Valid options are:")
					for key, name := range fractals {
						fmt.Printf("  %d\t%s\n", key, name)
					}
				}
			} else {
				fmt.Println("Usage: mandelbrot-cli --fractal [OPTION]")
				fmt.Println()
				fmt.Println("Current fractal set options:")
				fmt.Println(" 0\t\tMandelbrot Set (the default option)")
				fmt.Println(" 1\t\tJulia Set")
				fmt.Println(" 2\t\tBurning Ship")
				fmt.Println()
				fmt.Println("Report bugs to <https://github.com/MicheleFiladelfia>")
			}
		default:
			fmt.Println("mandelbrot-cli: unrecognized option '", os.Args[1], "'")
			fmt.Println("Try '--help' for more information.")
		}

		os.Exit(0)
	}
	startWithModel(initialModel())
}

func startWithModel(m *model) {
	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
