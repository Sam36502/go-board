/*
 *
 *		GO BOARD
 *
 *	Handles managing and displaying
 *	of 2D ASCII boards.
 *
 */

package main

import (
	"fmt"
	"strings"

	"github.com/mgutz/ansi"
)

/*
	TYPES
*/

type Coord struct {
	X int
	Y int
}

// Interface of things to be put on the board
// Chars is a string array of characters to display
// in the Tile when printed
type Tile interface {
	SetColour(Colour)
	SetChars([]string)
	GetANSIString() string
	GetChars() []string
	GetWidth() int
	GetHeight() int
}

// Holds all the board-related data
type Board struct {
	tiles  [][]Tile
	pieces map[Coord]Tile
}

/*
	FUNCTIONS
*/

// Creates a new board with `initTile` as the
// default tile everything is set as
func NewBoard(width, height int, initTile Tile) *Board {
	b := Board{}
	for y := 0; y < height; y++ {
		b.tiles = append(b.tiles, []Tile{})
		for x := 0; x < width; x++ {
			b.tiles[y] = append(b.tiles[y], initTile)
		}
	}
	b.pieces = map[Coord]Tile{}
	return &b
}

// Gets the tile from a specific position
func (b *Board) GetTile(pos Coord) Tile {
	if pos.X < 0 || pos.X >= b.GetWidth() || pos.Y < 0 || pos.Y >= b.GetHeight() {
		return nil
	}
	return b.tiles[pos.Y][pos.X]
}

// Sets a tile at a specific position
func (b *Board) SetTile(pos Coord, t Tile) {
	if pos.X < 0 || pos.X >= b.GetWidth() || pos.Y < 0 || pos.Y >= b.GetHeight() {
		return
	}
	b.tiles[pos.Y][pos.X] = t
}

// Gets the piece from a specific position
func (b *Board) GetPiece(pos Coord) (Tile, bool) {
	if pos.X < 0 || pos.X >= b.GetWidth() || pos.Y < 0 || pos.Y >= b.GetHeight() {
		return nil, false
	}
	piece, exists := b.pieces[pos]
	return piece, exists
}

// Sets a piece at a specific position
func (b *Board) SetPiece(pos Coord, p Tile) {
	if pos.X < 0 || pos.X >= b.GetWidth() || pos.Y < 0 || pos.Y >= b.GetHeight() {
		return
	}
	b.pieces[pos] = p
}

// Moves a piece on a board
func (b *Board) MovePiece(start, end Coord) {
	if start.X < 0 || start.X >= b.GetWidth() || start.Y < 0 || start.Y >= b.GetHeight() {
		return
	}
	if end.X < 0 || end.X >= b.GetWidth() || end.Y < 0 || end.Y >= b.GetHeight() {
		return
	}
	b.pieces[end] = b.pieces[start]
	delete(b.pieces, start)
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
func (b *Board) RenderString(border Border) string {

	// Get an example tile as size reference
	// TODO: mixed boards with different tiles wouldn't work with this
	var exTile Tile = nil
	if b.GetWidth() > 0 && b.GetHeight() > 0 {
		exTile = b.GetTile(Coord{0, 0})
	} else {
		return fmt.Sprint(
			border[BORDER_TOP_LEFT],
			border[BORDER_TOP_RIGHT],
			'\n',
			border[BORDER_BOTTOM_LEFT],
			border[BORDER_BOTTOM_RIGHT],
			'\n',
		)
	}

	// Top Border
	var renderedStr strings.Builder
	renderedStr.WriteRune(border[BORDER_TOP_LEFT])
	for i := 0; i < b.GetWidth()*exTile.GetWidth(); i++ {
		renderedStr.WriteRune(border[BORDER_SIDE_TOP])
	}
	renderedStr.WriteRune(border[BORDER_TOP_RIGHT])
	renderedStr.WriteByte('\n')

	// Contents
	for y := b.GetHeight() - 1; y >= 0; y-- {
		row := b.tiles[y]

		// In case the tile spans multiple rows
		for ty := 0; ty < exTile.GetHeight(); ty++ {

			renderedStr.WriteRune(border[BORDER_SIDE_LEFT])
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
			renderedStr.WriteRune(border[BORDER_SIDE_RIGHT])
			renderedStr.WriteByte('\n')

		}
	}

	// bottom Border
	renderedStr.WriteRune(border[BORDER_BOTTOM_LEFT])
	for i := 0; i < b.GetWidth()*exTile.GetWidth(); i++ {
		renderedStr.WriteRune(border[BORDER_SIDE_BOTTOM])
	}
	renderedStr.WriteRune(border[BORDER_BOTTOM_RIGHT])
	renderedStr.WriteByte('\n')

	return renderedStr.String()
}
