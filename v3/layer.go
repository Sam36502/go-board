/*
 *
 *	Implementation of LayerRenderer
 *
 *	Handles managing and displaying
 *	of 2D ASCII boards.
 *
 */

package board

import (
	"fmt"
	"strings"

	"github.com/mgutz/ansi"
)

type Layer struct {
	squares [][]SquareRenderer
	border  Border
}

// Will raise an error if Layer doesn't implement LayerRenderer
var _ LayerRenderer = (*Layer)(nil)

// Creates a new layer with the given width, height and border
func NewLayer(width, height int, border Border) *Layer {
	squares := make([][]SquareRenderer, height)
	for i := 0; i < len(squares); i++ {
		squares[i] = make([]SquareRenderer, width)
	}

	return &Layer{
		squares: squares,
		border:  border,
	}
}

// Returns width of this layer
func (l *Layer) GetWidth() int {
	return len(l.squares[0])
}

// Returns height of this layer
func (b *Layer) GetHeight() int {
	return len(b.squares)
}

// Sets the border to use when rendering this layer
func (l *Layer) SetBorder(border Border) {
	l.border = border
}

// Gets the square from a specific position
func (l *Layer) GetSquare(pos Coord) SquareRenderer {
	if !pos.IsInBounds(l.GetWidth(), l.GetHeight()) {
		return nil
	}
	return l.squares[pos.Y][pos.X]
}

// Sets a square at a specific position
func (l *Layer) SetSquare(pos Coord, s SquareRenderer) {
	if !pos.IsInBounds(l.GetWidth(), l.GetHeight()) {
		return
	}
	l.squares[pos.Y][pos.X] = s
}

// Render a layer as a string with ANSI
// control codes for the colours
func (l *Layer) RenderString() string {

	// Get an example tile as size reference
	// TODO: mixed boards with different tiles wouldn't work with this
	var exTile SquareRenderer = nil
	if l.GetWidth() > 0 && l.GetHeight() > 0 {
		exTile = l.GetSquare(Coord{0, 0})
	} else {
		return fmt.Sprint(
			l.border[BORDER_TOP_LEFT],
			l.border[BORDER_TOP_RIGHT],
			'\n',
			l.border[BORDER_BOTTOM_LEFT],
			l.border[BORDER_BOTTOM_RIGHT],
			'\n',
		)
	}

	// Top Border
	var renderedStr strings.Builder
	renderedStr.WriteRune(l.border[BORDER_TOP_LEFT])
	for i := 0; i < l.GetWidth()*exTile.GetWidth(); i++ {
		renderedStr.WriteRune(l.border[BORDER_SIDE_TOP])
	}
	renderedStr.WriteRune(l.border[BORDER_TOP_RIGHT])
	renderedStr.WriteByte('\n')

	// Contents
	for y := l.GetHeight() - 1; y >= 0; y-- {
		row := l.squares[y]

		// In case the tile spans multiple rows
		for ty := 0; ty < exTile.GetHeight(); ty++ {

			renderedStr.WriteRune(l.border[BORDER_SIDE_LEFT])
			for _, tile := range row {

				if tile == nil {
					tile = &DEFAULT_SQUARE
				}

				renderedStr.WriteString(tile.GetANSIString())

				tileRow := ""
				if len(tile.GetChars()) > ty {
					tileRow = tile.GetChars()[ty]
				}

				for tx := 0; tx < exTile.GetWidth(); tx++ {
					char := byte(' ')
					if len(tileRow) > tx {
						char = tileRow[tx]
					}
					renderedStr.WriteByte(char)
				}
				renderedStr.WriteString(ansi.Reset)
			}
			renderedStr.WriteRune(l.border[BORDER_SIDE_RIGHT])
			renderedStr.WriteByte('\n')

		}
	}

	// bottom Border
	renderedStr.WriteRune(l.border[BORDER_BOTTOM_LEFT])
	for i := 0; i < l.GetWidth()*exTile.GetWidth(); i++ {
		renderedStr.WriteRune(l.border[BORDER_SIDE_BOTTOM])
	}
	renderedStr.WriteRune(l.border[BORDER_BOTTOM_RIGHT])
	renderedStr.WriteByte('\n')

	return renderedStr.String()
}

// Prints the rendered board with ANSI control chars
func (l *Layer) PrintBoard() {
	PrintANSIString(l.RenderString() + "\n")
}

//// Util Functions ////

// Loops through the layer and sets the square based on a function
func (l *Layer) FillPattern(patternFunc func(Coord) SquareRenderer) {
	for y := 0; y < l.GetHeight(); y++ {
		for x := 0; x < l.GetWidth(); x++ {
			pos := Coord{x, y}
			l.SetSquare(pos, patternFunc(pos))
		}
	}
}

// Fills a layer with a single square
func (l *Layer) FillLayer(sq SquareRenderer) {
	l.FillPattern(func(c Coord) SquareRenderer {
		return sq
	})
}

// Fills a certain area with a square
// TODO: Make more efficient with fancy mafs
func (l *Layer) FillArea(a Coord, b Coord, sq SquareRenderer) {
	l.FillPattern(func(c Coord) SquareRenderer {
		if !c.IsInBounds(l.GetWidth(), l.GetHeight()) ||
			!a.IsInBounds(l.GetWidth(), l.GetHeight()) ||
			!b.IsInBounds(l.GetWidth(), l.GetHeight()) {
			return l.GetSquare(c)
		}
		x1 := a.X
		x2 := b.X
		if a.X > b.X {
			x1 = b.X
			x2 = a.X
		}
		y1 := a.Y
		y2 := b.Y
		if a.Y > b.Y {
			y1 = b.Y
			y2 = a.Y
		}

		if c.X >= x1 && c.X <= x2 && c.Y >= y1 && c.Y <= y2 {
			return sq
		} else {
			return l.GetSquare(c)
		}
	})
}
