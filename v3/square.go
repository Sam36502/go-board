/*
 *
 *	Implementation of SquareRenderer
 *
 *	(Feel free to overwrite with non-
 *	square squares :D)
 *
 */

package board

import (
	"strings"
)

const (
	SQUARE_WIDTH  = 2 // Makes a roughly square square
	SQUARE_HEIGHT = 1 // (depending on font)
)

type Square struct {
	width  int
	height int
	Colour Colour
	Chars  []string
}

func NewSquare(w, h int, c Colour, cs []string) *Square {
	return &Square{
		width:  w,
		height: h,
		Colour: c,
		Chars:  cs,
	}
}

func (p *Square) SetColour(colour Colour) {
	p.Colour = colour
}

func (p *Square) GetColour() Colour {
	return p.Colour
}

func (p *Square) SetChars(chars []string) {
	p.Chars = chars
}

func (p *Square) GetChars() []string {
	return p.Chars
}

func (p *Square) GetWidth() int {
	if p.width <= 0 {
		p.width = SQUARE_WIDTH
	}
	return p.width
}

func (p *Square) GetHeight() int {
	if p.height <= 0 {
		p.height = SQUARE_HEIGHT
	}
	return p.height
}

// Can just return an empty string for no formatting
// if you don't want to use ANSI
func (p *Square) GetANSIString() string {
	return p.Colour.GetANSIString()
}

func (p *Square) RenderString() string {
	var render strings.Builder
	for y := 0; y < p.GetHeight(); y++ {
		render.WriteString(p.GetANSIString())
		row := ""
		if y < len(p.GetChars()) {
			row = p.GetChars()[y]
		}

		for x := 0; x < p.GetWidth(); x++ {
			if x < len(row) {
				render.WriteByte(row[x])
			} else {
				render.WriteRune(' ')
			}
		}
		render.WriteString(ANSI_RESET_COLOUR)
		render.WriteRune('\n')
	}
	return render.String()
}

// Will raise an error if Pixel doesn't implement Tile
var _ SquareRenderer = (*Square)(nil)
