/*
 *
 *	Basic Implementation
 *	of Tile interface
 *
 */

package board

const (
	PIXEL_WIDTH  = 2 // Makes a roughly
	PIXEL_HEIGHT = 1 // square tile
)

type Pixel struct {
	Colour Colour
	Chars  []string
}

func (p *Pixel) GetWidth() int {
	return PIXEL_WIDTH
}

func (p *Pixel) GetHeight() int {
	return PIXEL_HEIGHT
}

func (p *Pixel) SetColour(colour Colour) {
	p.Colour = colour
}

func (p *Pixel) SetChars(chars []string) {
	p.Chars = chars
}

// Can just return an empty string for no formatting
// if you don't want to use ANSI
func (p *Pixel) GetANSIString() string {
	return p.Colour.GetANSIString()
}

func (p *Pixel) GetChars() []string {
	return p.Chars
}

// Will raise an error if Pixel doesn't implement Tile
var _ Tile = (*Pixel)(nil)
