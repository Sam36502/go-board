/*
 *
 *	Basic Implementation
 *	of Piece interface
 *
 */

package board

const (
	TOKEN_WIDTH  = 2 // Makes a roughly
	TOKEN_HEIGHT = 1 // square Piece
)

type Token struct {
	Position Coord
	Colour   Colour
	Chars    []string
}

func (t *Token) GetWidth() int {
	return PIXEL_WIDTH
}

func (t *Token) GetHeight() int {
	return PIXEL_HEIGHT
}

func (t *Token) SetColour(colour Colour) {
	t.Colour = colour
}

func (t *Token) SetChars(chars []string) {
	t.Chars = chars
}

// Can just return an empty string for no formatting
// if you don't want to use ANSI
func (t *Token) GetANSIString() string {
	return t.Colour.GetANSIString()
}

func (t *Token) GetChars() []string {
	return t.Chars
}

func (t *Token) MovePiece(coord Coord) {
	t.Position = coord
}

func (t *Token) GetPosition() Coord {
	return t.Position
}

// Will raise an error if Token doesn't implement Piece
var _ Piece = (*Token)(nil)
