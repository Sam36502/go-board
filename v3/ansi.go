/*
 *
 *		ANSI Control Code Utils
 *		and helper functions
 *
 */
package board

import (
	"fmt"

	"github.com/mattn/go-colorable"
	"github.com/mgutz/ansi"
)

const (
	ANSI_CLEAR_SCREEN = "\x1b[2J\x1b[H"
	ANSI_RESET_COLOUR = ansi.Reset
)

type Colour struct {
	Foreground string
	Background string
}

// Gets the ANSI control code for this colour
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

// Util function which handles ANSI codes for Windows
func PrintANSIString(code string) {
	stdOut := colorable.NewColorableStdout()
	fmt.Fprint(stdOut, code)
}

func ClearScreen() {
	PrintANSIString(ANSI_CLEAR_SCREEN)
}

func ResetColour() {
	PrintANSIString(ANSI_RESET_COLOUR)
}
