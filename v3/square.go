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

	"github.com/mgutz/ansi"
)

const (
	SQUARE_WIDTH  = 2 // Makes a roughly square square
	SQUARE_HEIGHT = 1 // (depending on font)
)

var DEFAULT_SQUARE = Square{
	width:  SQUARE_WIDTH,
	height: SQUARE_HEIGHT,
	Colour: Colour{
		Foreground: ansi.DefaultFG,
		Background: ansi.DefaultBG,
	},
	Chars: []string{},
}

type Square struct {
	width  int
	height int
	Colour Colour
	Chars  []string
}

// Will raise an error if Square doesn't implement SquareRenderer
var _ SquareRenderer = (*Square)(nil)

// Creates a new Square with the specified size, colour & characters
func NewSquare(w, h int, c Colour, cs []string) *Square {
	return &Square{
		width:  w,
		height: h,
		Colour: c,
		Chars:  cs,
	}
}

// Set this square colour
func (p *Square) SetColour(colour Colour) {
	p.Colour = colour
}

// Get this square's colour
func (p *Square) GetColour() Colour {
	return p.Colour
}

// Set this square's chars
func (p *Square) SetChars(chars []string) {
	p.Chars = chars
}

// Get this square's chars
func (p *Square) GetChars() []string {
	return p.Chars
}

// Get this square's width
func (p *Square) GetWidth() int {
	if p.width <= 0 {
		p.width = SQUARE_WIDTH
	}
	return p.width
}

// Get this square's width
func (p *Square) GetHeight() int {
	if p.height <= 0 {
		p.height = SQUARE_HEIGHT
	}
	return p.height
}

// Gets this square's colour's ANSI code
// When implementing, you can just return an empty
// string for no formatting if you don't want to use ANSI
func (p *Square) GetANSIString() string {
	return p.Colour.GetANSIString()
}

// Renders this square as a string
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
