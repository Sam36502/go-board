/*
 *
 *		ANSI Control Code Utils
 *
 */
package board

import (
	"bufio"
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

func printAnsi(code string) {
	stdOut := bufio.NewWriter(colorable.NewColorableStdout())
	fmt.Fprint(stdOut, code)
}

func ClearScreen() {
	printAnsi(ANSI_CLEAR_SCREEN)
}

func ResetColour() {
	printAnsi(ANSI_RESET_COLOUR)
}
