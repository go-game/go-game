# go-game

A 2D game framework for the go programming language.

## Installation
```bash
go get -u -v github.com/go-gl/gl/v2.1/gl
go get -u -v github.com/veandco/go-sdl2/sdl
go get -u -v github.com/veandco/go-sdl2/img
go get -u -v github.com/veandco/go-sdl2/mix
go get -u -v github.com/veandco/go-sdl2/ttf
```

## Notes for OS X

### Compilation
Xcode 8.3 had some changes so that the linking of lib sdl was broken. If you get the error `signal: killed` run or build your game with the flags `-ldflags -s`. See: [Github Issue](https://github.com/golang/go/issues/19734)

### Dependencies
You need pkg-config and SDL2 libs.

Install with Homebrew:
```bash
brew install pkg-config sdl2 sdl2_image sdl2_mixer sdl2_ttf
```

### Controller support

For XBox 360 controller support on Mac, see: https://github.com/360Controller/360Controller/releases
## Used Font

The used font OpenSans-Regular was downloaded from [www.fontsquirrel.com](https://www.fontsquirrel.com/fonts/open-sans).

It is licensed under the Apache License v2.00.
