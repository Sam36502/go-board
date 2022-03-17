/*
 *
 *	Basic Implementation
 *	of Tile interface
 *
 */

package board

import (
	"fmt"

	"github.com/mgutz/ansi"
)

const (
	PIXEL_WIDTH  = 2 // Makes a roughly
	PIXEL_HEIGHT = 1 // square tile
)

type Pixel struct {
	fg    Colour
	bg    Colour
	chars []string
}

func (p *Pixel) GetSize() (int, int) {
	return PIXEL_WIDTH, PIXEL_HEIGHT
}

func (p *Pixel) SetColours(fg, bg Colour) {
	p.fg = fg
	p.bg = bg
}

func (p *Pixel) SetChars(chars []string) {
	p.chars = chars
}

// Can just return an empty string for no formatting
// if you don't want to use ANSI
func (p *Pixel) GetColourCode() string {
	return ansi.ColorCode(fmt.Sprint(p.fg, ":", p.bg))
}

func (p *Pixel) GetChars() []string {
	return p.chars
}
