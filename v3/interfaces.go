/*
 *
 *	Board Interfaces
 *
 *	The various interfaces for the different
 *	tiers of board information
 *
 */
package board

// Used for rendering the individual board "squares".
// Note: Doesn't necessarily need to be square.
// Chars is the strings to render inside the square.
// They will render as empty strings if not set.
type SquareRenderer interface {
	SetColour(Colour)
	GetColour() Colour
	SetChars([]string)
	GetChars() []string
	GetWidth() int
	GetHeight() int
	GetANSIString() string
	RenderString() string
}

// Used for rendering the individual layers of the board.
// Any empty tiles will show the layer underneath when
// rendered, or the default colours if there's no lower layer.
type LayerRenderer interface {
	SetSquare(Coord, SquareRenderer)
	GetSquare(Coord) SquareRenderer
	GetWidth() int
	GetHeight() int
	RenderString() string
}

// Used for rendering a complete board with many layers.
// Any layers outside the board's size will be truncated.
type BoardRenderer interface {
	SetLayer(int, LayerRenderer)
	GetLayer(int) LayerRenderer
	SetBorder(Border)
	GetBorder(Border)
	GetWidth() int
	GetHeight() int
	RenderString() string
	PrintBoard()
}
