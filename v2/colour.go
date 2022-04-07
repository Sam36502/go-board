package main

import (
	"fmt"

	"github.com/mgutz/ansi"
)

type Colour struct {
	Foreground string
	Background string
}

func (c *Colour) GetANSIString() string {
	return ansi.ColorCode(fmt.Sprint(c.Foreground, ":", c.Background))
}

// Colour Constants
const (
	Black        = "black"
	Red          = "red"
	Green        = "green"
	Yellow       = "yellow"
	Blue         = "blue"
	Magenta      = "magenta"
	Cyan         = "cyan"
	White        = "white"
	LightBlack   = "black+h"
	LightRed     = "red+h"
	LightGreen   = "green+h"
	LightYellow  = "yellow+h"
	LightBlue    = "blue+h"
	LightMagenta = "magenta+h"
	LightCyan    = "cyan+h"
	LightWhite   = "white+h"
)
