# fractals-cli

A multiplatform, high-performance terminal fractals explorer written in pure Go, powered by concurrent rendering and memoization for smooth real-time exploration.

<p align="center">
  <img src="https://github.com/MicheleFiladelfia/fractals-cli/assets/86882607/85cac069-fb07-4605-9270-1638a29c9568" width="48%" />
  <img src="https://github.com/user-attachments/assets/633018d5-26db-42e5-b07d-378d56e33eb2" width="48%" />
</p>
<p align="center">
  <img src="https://github.com/user-attachments/assets/77acb0a2-98ed-4c56-afc5-1907c1e4e9df" width="48%" />
  <img src="https://user-images.githubusercontent.com/86882607/181925437-d523ac86-a52c-4b68-a5a2-1f124f10daa4.png" width="48%" />
</p>
<p align="center">
  <img src="https://user-images.githubusercontent.com/86882607/181925452-11120aaf-cdcb-4ee6-a1af-18b1dea4e9ba.png" width="48%" />
  <img src="https://user-images.githubusercontent.com/86882607/181925538-d5cda89e-4906-449a-a895-f30aa1fa9f4a.png" width="48%" />
</p>

***

## Controls

- **Up:** W
- **Down:** S
- **Left:** A
- **Right:** D
- **Zoom in:** I
- **Zoom out:** O
- **More Iterations:** +
- **Less Iterations:** -
- **Increase Bailout:** B
- **Decrease Bailout:** V
- **Toggle coordinates bar:** X
- **Toggle controls help:** ? or H
- **Change color palette:** C
- **Change Fractal Set:** N (Switch between 13 fractal sets: Mandelbrot, Julia, Burning Ship, Tricorn, Celtic, Buffalo, Multibrot 3/4, Newton, Perpendicular Ship, Celtic Ship, Lambda, and Shark Fin)
- **Quit:** ctrl+c

## Command line flags
- **--help** small guide on how to use fractal-cli
- **--version** get version number
- **--nocache** disable memoization cache optimization (enabled by default)


## Render quality of the fractal
The quality of the fractal depends on how many pixels you have in your terminal, if you want to display more pixels, on most terminals you can do so with ctrl+- or by scrolling the mouse wheel.

Note: The more pixels you display in your terminal, the slower the rendering will be

***

## Compiling and running

Firstly you have to [install golang](https://go.dev/doc/install) and then...

```bash
git clone https://github.com/MicheleFiladelfia/fractals-cli

cd fractals-cli

#*nix
go build -o fractals-cli
#windows
go build -o fractals-cli.exe

./fractals-cli
```

***

## Recommended Terminals

You can run `fractals-cli` in most Unix and Windows terminals. However, for the best visual experience and performance, we recommend using one of the tested terminals below:

- ghostty (Fastest native macOS/Linux GPU-accelerated terminal)
- alacritty (Fastest in rendering and more pixels)
- kitty
- gnome-terminal
- konsole
- terminator
- yakuake
- xfce4-terminal
- windows terminals

***

## Acknowledgments

In this software is used [The ELM Architecture](https://guide.elm-lang.org/architecture/) provided by [bubbletea](https://github.com/charmbracelet/bubbletea) TUI framework.

***

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
