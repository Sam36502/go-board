/*
 *
 *	Implementation of BoardRenderer
 *
 *	Handles managing and displaying
 *	of 2D ASCII boards.
 *
 */
package board

import (
	"sort"
)

type Board struct {
	width  int
	height int
	layers map[int]LayerRenderer
	border Border
}

// Creates a new Board with one layer of the size provided
// And the border set
func NewBoard(width, height int, border Border) *Board {
	return &Board{
		width:  width,
		height: height,
		layers: make(map[int]LayerRenderer),
		border: border,
	}
}

// Will raise an error if Layer doesn't implement LayerRenderer
var _ BoardRenderer = (*Board)(nil)

// Sets a layer at a certain index
func (b *Board) SetLayer(index int, layer LayerRenderer) {
	b.layers[index] = layer
}

// Gets a layer from a certain index
func (b *Board) GetLayer(index int) (LayerRenderer, bool) {
	l, ok := b.layers[index]
	return l, ok
}

// Sets this board's border to be displayed when rendered
func (b *Board) SetBorder(border Border) {
	b.border = border
}

// Get this board's border
func (b *Board) GetBorder() Border {
	return b.border
}

// Get this board's width
func (b *Board) GetWidth() int {
	return b.width
}

// Get this board's width
func (b *Board) GetHeight() int {
	return b.height
}

// Renders this board as a string with ANSI codes
// Stacks each layer visually so any layers with
// empty squares will show the layer beneath it.
func (b *Board) RenderString() string {
	renderBuf := NewLayer(b.GetWidth(), b.GetHeight(), SQUARE_WIDTH, SQUARE_HEIGHT, b.border)

	// Get sorted indices
	indices := make([]int, 0)
	for k := range b.layers {
		indices = append(indices, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(indices)))

	for _, li := range indices {
		lar, ok := b.GetLayer(li)
		if !ok {
			continue
		}
		for y := 0; y < b.GetWidth(); y++ {
			for x := 0; x < b.GetWidth(); x++ {
				pos := Coord{x, y}
				if renderBuf.GetSquare(pos) == nil && lar.GetSquare(pos) != nil {
					renderBuf.SetSquare(pos, lar.GetSquare(pos))
				}
			}
		}
	}
	return renderBuf.RenderString()
}

// Prints this board to the console rendered in
// glorious 2D with ANSI control codes for colour
// (Should also works on Windows)
func (b *Board) PrintBoard() {
	PrintANSIString(b.RenderString() + "\n")
}

//// Utility Functions ////

// Creates a new layer with the same size/border as this board
func (b *Board) CreateLayer(sqWidth, sqHeight int) *Layer {
	return NewLayer(b.GetWidth(), b.GetHeight(), sqWidth, sqHeight, b.GetBorder())
}
