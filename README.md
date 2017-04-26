# go-game

A 2d game framwork for the go programming language.

## Notes for OS X

Xcode 8.3 had some changes so that the linking of lib sdl was broken. If you get the error `signal: killed` run or build your game with the flags `-ldflags -s`. See: [Github Issue](https://github.com/golang/go/issues/19734)

## Used Font

The used font OpenSans-Regular was downloaded from [www.fontsquirrel.com](https://www.fontsquirrel.com/fonts/open-sans).

It is licensed under the Apache License v2.00.
