# fractals-cli

Multiplatform, Elegant and clean terminal fractals explorer written in pure golang with concurrency to make the computing of the fractal more efficient.

![fractals-cli](https://user-images.githubusercontent.com/86882607/181925302-a46801f0-bef7-44c9-ac13-95481472127b.png)
![fractals-cli](https://user-images.githubusercontent.com/86882607/181925339-26222e9d-737f-404f-a200-c055b1674d72.png)
![fractals-cli](https://user-images.githubusercontent.com/86882607/181925342-0fa345ff-29d5-46fe-bfa6-f4bdf10b578f.png)
![fractals-cli](https://user-images.githubusercontent.com/86882607/181925350-7f31e072-861e-4552-bf4e-b42de4dc8082.png)
![fractals-cli](https://user-images.githubusercontent.com/86882607/181925358-047d9ecf-7db2-44d5-97dd-277c3ca802f0.png)
![fractals-cli](https://user-images.githubusercontent.com/86882607/181925376-ab1aa80b-6594-4e69-a64e-5f757dea2a5d.png)
![fractals-cli](https://user-images.githubusercontent.com/86882607/181925437-d523ac86-a52c-4b68-a5a2-1f124f10daa4.png)
![fractals-cli](https://user-images.githubusercontent.com/86882607/181925452-11120aaf-cdcb-4ee6-a1af-18b1dea4e9ba.png)
![fractals-cli](https://user-images.githubusercontent.com/86882607/181925489-056a74bb-d036-4770-9798-8148b86f3d05.png)
![fractals-cli](https://user-images.githubusercontent.com/86882607/181925538-d5cda89e-4906-449a-a895-f30aa1fa9f4a.png)

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
- **Change color palette:** C
- **Change Fractal Set:** N (Switch between Mandelbrot, Julia, and Burning Ship sets)
- **Quit:** ctrl+c

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

## The best terminals where to run fractals-cli

You can run fractals-cli in most of *nix terminals, but if you want the best possible experience, you have to run it in one how the tested terminals below:

- gnome-terminal
- konsole
- terminator
- yakuake
- kitty
- alacritty (Fastest in rendering and best quality)
- xfce4-terminal
- windows terminals

***

## Acknowledgments

In this software is used [The ELM Architecture](https://guide.elm-lang.org/architecture/) provided by [bubbletea](https://github.com/charmbracelet/bubbletea) TUI framework.

***

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
