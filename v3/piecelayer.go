/*
 *
 *	Implementation of LayerRenderer
 *	that also handles moving pieces around
 *
 *	Handles managing and displaying
 *	of 2D ASCII boards.
 *
 */

package board

type PieceLayer struct {
	renderLayer LayerRenderer
	pieces      map[string]Coord
}

// Adds a piece to this layer to be rendered
func (l *PieceLayer) SetPiece(id string, pos Coord, p SquareRenderer) {
	if oldPos, exists := l.pieces[id]; exists {
		l.renderLayer.SetSquare(oldPos, nil)
	}

	l.pieces[id] = pos
	l.renderLayer.SetSquare(pos, p)
}

// Removes a piece from this layer
func (l *PieceLayer) RemPiece(id string) {
	if _, exists := l.pieces[id]; exists {
		l.renderLayer.SetSquare(l.pieces[id], nil)
		delete(l.pieces, id)
	}
}

// Move a piece along a given vector
func (l *PieceLayer) MovePiece(id string, move Vector) {
	if oldPos, exists := l.pieces[id]; exists {
		p := l.renderLayer.GetSquare(oldPos)
		l.renderLayer.SetSquare(oldPos, nil)

		newPos := oldPos.Add(move)
		l.pieces[id] = newPos
		l.renderLayer.SetSquare(newPos, p)
	}
}

// Get a piece by its ID
func (l *PieceLayer) GetPiece(id string) SquareRenderer {
	return l.renderLayer.GetSquare(l.pieces[id])
}

// Name wrapper for consistency
func (l *PieceLayer) GetPieceAt(pos Coord) SquareRenderer {
	return l.renderLayer.GetSquare(pos)
}

//// Layer Wrapper Functions ////

// Will raise an error if Layer doesn't implement LayerRenderer
var _ LayerRenderer = (*PieceLayer)(nil)

// Creates a new layer with the given width, height and border
func NewPieceLayer(width, height int, border Border) *PieceLayer {
	return &PieceLayer{
		renderLayer: NewLayer(width, height, border),
		pieces:      make(map[string]Coord, 0),
	}
}

// Returns width of this layer
func (l *PieceLayer) GetWidth() int {
	return l.renderLayer.GetWidth()
}

// Returns height of this layer
func (l *PieceLayer) GetHeight() int {
	return l.renderLayer.GetHeight()
}

// Sets the border to use when rendering this layer
func (l *PieceLayer) SetBorder(border Border) {
	l.renderLayer.SetBorder(border)
}

// Gets the border of this layer
func (l *PieceLayer) GetBorder() Border {
	return l.renderLayer.GetBorder()
}

// Gets the square from a specific position
func (l *PieceLayer) GetSquare(pos Coord) SquareRenderer {
	return l.renderLayer.GetSquare(pos)
}

// Sets a square at a specific position
// (Disabled as render layer is read-only)
func (l *PieceLayer) SetSquare(pos Coord, s SquareRenderer) {
}

// Render a layer as a string with ANSI
// control codes for the colours
func (l *PieceLayer) RenderString() string {
	return l.renderLayer.RenderString()
}

// Prints the rendered board with ANSI control chars
func (l *PieceLayer) PrintBoard() {
	PrintANSIString(l.RenderString() + "\n")
}
