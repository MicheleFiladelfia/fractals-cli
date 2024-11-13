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
			fmt.Println("fractals-cli (v1.1.0), created by @MicheleFiladelfia")
		case "--help":
			fmt.Println("fractals-cli: a simple fractals terminal explorer.")
			fmt.Println("Usage: fractals-cli [OPTION]")
			fmt.Println()
			fmt.Println("Options:")
			fmt.Println("  --version\t\tPrint version information and exit.")
			fmt.Println("  --help\t\tPrint this help message and exit.")
			fmt.Println("  --usecache\t\tRuns fractal-cli with dynamic programming optimization (a little bit faster but may be unstable)")
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
			fmt.Println("  B\t\t\tIncrease bailout by 0.1.")
			fmt.Println("  V\t\t\tDecrease bailout by 0.1.")
			fmt.Println("  ctrl+c\t\tQuit.")
			fmt.Println()
			fmt.Println("Report bugs to <https://github.com/MicheleFiladelfia>")
		case "--usecache":
			var mod = initialModel()
			mod.usecache = true
			startWithModel(mod)
		default:
			fmt.Println("fractals-cli: unrecognized option '", os.Args[1], "'")
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
