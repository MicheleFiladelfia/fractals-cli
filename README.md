# mandelbrot-cli

Multiplatform, Elegant and clean terminal mandelbrot fractal explorer written in pure golang with concurrency to make the computing of the fractal more efficient.

//images

***

## Controls

- **Up:** W
- **Down:** S
- **Left:** A
- **Right:** D
- **Zoom in:** Q
- **Zoom out:** E
- **More Iterations:** +
- **Less Iterations:** -
- **Change color palette:** C
- **Quit:** ctrl+c

The quality of the mandelbrot set depends on how many pixels you have in your terminal, if you want to display more pixels, on most terminals you can do so with ctrl+- or by scrolling the mouse wheel.

Note: The more pixels you display in your terminal, the slower the rendering will be

***

## Compiling and installing

Firstly you have to [install golang](https://go.dev/doc/install) and then...

```bash
git clone https://github.com/MicheleFiladelfia/mandelbrot-cli

cd mandelbrot-cli

#*nix
go build -o temp
#windows
go build -o temp.exe

go install

mandelbrot-cli
```

***

## The best terminals where to run mandelbrot-cli

You can run mandelbrot-cli in most of *nix terminals, but if you want the best possible experience, you have to run it in one how the tested terminals below:

- gnome-terminal
- konsole
- terminator
- yakuake
- kitty
- alacritty (Fastest in rendering and best quality but a little bit glitchy when rendering)
- xfce4-terminal

Note: mandelbrot-cli is not fully compatible with windows terminals, so is recommended to use a linux distribution.

***

## Acknowledgments

In this software is used [The ELM Architecture](https://guide.elm-lang.org/architecture/) provided by [bubbletea](https://github.com/charmbracelet/bubbletea) TUI framework.

***

## Contributing

If you want to help me improve this software and have an idea of what can be improved or want to make/add a [TODO](TODO.md), please contact me from my [website](https://michelefiladelfia.github.io/) or on my [telegram](https://t.me/Mic04_7)/[discord](https://discord.com/users/450266882394554378) and describe me your idea, then if i accept, you can do a pr.

***

## License

mandelbrot-cli is released under the [MIT LICENSE](LICENSE).

***
