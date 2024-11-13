# fractals-cli

Multiplatform, Elegant and clean terminal fractals explorer written in pure golang with concurrency and dynamic programming to make the computing of the fractal more efficient.

![fractals-cli](https://github.com/MicheleFiladelfia/fractals-cli/assets/86882607/85cac069-fb07-4605-9270-1638a29c9568)
![fractals-cli](https://github.com/MicheleFiladelfia/fractals-cli/assets/86882607/ec2d9b36-9006-4ac7-a236-1a03579eed88)
![fractals-cli](https://github.com/MicheleFiladelfia/fractals-cli/assets/86882607/83bb9686-6d1b-4a6f-b386-a61d4af42f1c)
![fractals-cli](https://github.com/MicheleFiladelfia/fractals-cli/assets/86882607/d8050f5d-7a9e-48ce-afcd-b3136a4b483b)
![fractals-cli](https://user-images.githubusercontent.com/86882607/181925350-7f31e072-861e-4552-bf4e-b42de4dc8082.png)
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
- **Increase Bailout:** B
- **Decrease Bailout:** V
- **Change color palette:** C
- **Change Fractal Set:** N (Switch between Mandelbrot, Julia, and Burning Ship sets)
- **Quit:** ctrl+c

## Command line flags
- **--help** small guide on how to use fractal-cli
- **--version** get version number
- **--usecache** uses dynamic programming optimizations (faster but probably unstable)


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

## The best terminals where to run fractals-cli

You can run fractals-cli in most of *nix terminals, but if you want the best possible experience, you have to run it in one how the tested terminals below:

- gnome-terminal
- konsole
- terminator
- yakuake
- kitty
- alacritty (Fastest in rendering and more pixels)
- xfce4-terminal
- windows terminals

***

## Acknowledgments

In this software is used [The ELM Architecture](https://guide.elm-lang.org/architecture/) provided by [bubbletea](https://github.com/charmbracelet/bubbletea) TUI framework.

***

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
