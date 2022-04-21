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

func (b *Board) SetBorder(border Border) {
	b.border = border
}

// Gets the tile from a specific position
func (b *Board) GetTile(pos Coord) Tile {
	if !pos.IsInBounds(b.GetWidth(), b.GetHeight()) {
		return nil
	}
	return b.tiles[pos.Y][pos.X]
}

// Sets a tile at a specific position
func (b *Board) SetTile(pos Coord, t Tile) {
	if !pos.IsInBounds(b.GetWidth(), b.GetHeight()) {
		return
	}
	b.tiles[pos.Y][pos.X] = t
}

// Gets a list of all piece coordinates on the board
func (b *Board) GetPieceCoords() []Coord {
	coords := make([]Coord, len(b.pieces))
	i := 0
	for k := range b.pieces {
		coords[i] = k
		i++
	}
	return coords
}

// Gets the piece from a specific position and returns whether it exists or not
func (b *Board) GetPiece(pos Coord) (Tile, bool) {
	if !pos.IsInBounds(b.GetWidth(), b.GetHeight()) {
		return nil, false
	}
	piece, exists := b.pieces[pos]
	return piece, exists
}

// Sets a piece at a specific position
func (b *Board) SetPiece(pos Coord, p Tile) {
	if !pos.IsInBounds(b.GetWidth(), b.GetHeight()) {
		return
	}
	b.pieces[pos] = p
}

// Remove a piece from the board
func (b *Board) DeletePiece(pos Coord) {
	delete(b.pieces, pos)
}

// Checks if a certain move is valid (within bounds &
// not blocked by any existing pieces) and returns a bool
func (b *Board) IsMoveValid(startPos Coord, endPos Coord) bool {
	_, pieceExists := b.GetPiece(startPos)
	_, spaceOccupied := b.GetPiece(endPos)

	return true &&
		startPos.IsInBounds(b.GetWidth(), b.GetHeight()) &&
		pieceExists &&
		endPos.IsInBounds(b.GetWidth(), b.GetHeight()) &&
		!spaceOccupied
}

// Moves a piece at the given position by the given vector
func (b *Board) MovePiece(startPos Coord, direction Vector) {
	endPos := startPos
	endPos = endPos.Add(direction)

	if !b.IsMoveValid(startPos, endPos) {
		return
	}

	b.pieces[endPos] = b.pieces[startPos]
	delete(b.pieces, startPos)
}

// Moves a piece at the given position to a specific position
func (b *Board) MovePieceTo(startPos Coord, endPos Coord) {
	if !b.IsMoveValid(startPos, endPos) {
		return
	}

	b.pieces[endPos] = b.pieces[startPos]
	delete(b.pieces, startPos)
}

// Returns width & height of this board
func (b *Board) GetWidth() int {
	return len(b.tiles[0])
}

func (b *Board) GetHeight() int {
	return len(b.tiles)
}

// Render a board as a string with ANSI
// control codes for the colours
func (b *Board) RenderString() string {

	// Get an example tile as size reference
	// TODO: mixed boards with different tiles wouldn't work with this
	var exTile Tile = nil
	if b.GetWidth() > 0 && b.GetHeight() > 0 {
		exTile = b.GetTile(Coord{0, 0})
	} else {
		return fmt.Sprint(
			b.border[BORDER_TOP_LEFT],
			b.border[BORDER_TOP_RIGHT],
			'\n',
			b.border[BORDER_BOTTOM_LEFT],
			b.border[BORDER_BOTTOM_RIGHT],
			'\n',
		)
	}

	// Top Border
	var renderedStr strings.Builder
	renderedStr.WriteRune(b.border[BORDER_TOP_LEFT])
	for i := 0; i < b.GetWidth()*exTile.GetWidth(); i++ {
		renderedStr.WriteRune(b.border[BORDER_SIDE_TOP])
	}
	renderedStr.WriteRune(b.border[BORDER_TOP_RIGHT])
	renderedStr.WriteByte('\n')

	// Contents
	for y := b.GetHeight() - 1; y >= 0; y-- {
		row := b.tiles[y]

		// In case the tile spans multiple rows
		for ty := 0; ty < exTile.GetHeight(); ty++ {

			renderedStr.WriteRune(b.border[BORDER_SIDE_LEFT])
			for x, tile := range row {

				// Check if a piece is here
				piece, exists := b.GetPiece(Coord{x, y})
				if exists {
					tile = piece
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
			renderedStr.WriteRune(b.border[BORDER_SIDE_RIGHT])
			renderedStr.WriteByte('\n')

		}
	}

	// bottom Border
	renderedStr.WriteRune(b.border[BORDER_BOTTOM_LEFT])
	for i := 0; i < b.GetWidth()*exTile.GetWidth(); i++ {
		renderedStr.WriteRune(b.border[BORDER_SIDE_BOTTOM])
	}
	renderedStr.WriteRune(b.border[BORDER_BOTTOM_RIGHT])
	renderedStr.WriteByte('\n')

	return renderedStr.String()
}

// Prints the rendered board with ANSI control chars
func (b *Board) PrintBoard() {
	printAnsi(b.RenderString() + "\n")
}
