/*
 *
 *	Implementation of BoardRenderer
 *
 *	Handles managing and displaying
 *	of 2D ASCII boards.
 *
 */
package board

type Board struct {
	width  int
	height int
	layers []LayerRenderer
	border Border
}

// Creates a new Board with one layer of the size provided
// And the border set
func NewBoard(width, height int, border Border) *Board {
	return &Board{
		width:  width,
		height: height,
		layers: []LayerRenderer{
			NewLayer(width, height, border),
		},
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
func (b *Board) GetLayer(index int) LayerRenderer {
	return b.layers[index]
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
	// TODO
	return ""
}

// Prints this board to the console rendered in
// glorious 2D with ANSI control codes for colour
// (Should also works on Windows)
func (b *Board) PrintBoard() {
	PrintANSIString(b.RenderString() + "\n")
}
