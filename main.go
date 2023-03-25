package main

import (
	"fmt"
	"os"

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
			fmt.Println("  ctrl+c\t\tQuit.")
			fmt.Println()
			fmt.Println("Report bugs to <https://github.com/MicheleFiladelfia>")
		default:
			fmt.Println("mandelbrot-cli: unrecognized option '", os.Args[1], "'")
			fmt.Println("Try '--help' for more information.")
		}

		os.Exit(0)
	}

	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
