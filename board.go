/*
 *
 *		GO BOARD
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

/*
	TYPES
*/
type Colour string

// Interface of things to be put on the board
// Chars is a string array of characters to display
// in the Tile when printed
type Tile interface {
	SetColours(Colour, Colour) // fg & bg
	SetChars([]string)
	GetColourCode() string // ANSI colour code
	GetChars() []string
	GetSize() (int, int) // width & height
}

type Board struct {
	data [][]Tile
}

/*
	FUNCTIONS
*/

// Creates a new board with `initTile` as the
// default tile everything is set as
func NewBoard(width, height int, initTile Tile) *Board {
	b := Board{}
	for y := 0; y < height; y++ {
		b.data = append(b.data, []Tile{})
		for x := 0; x < width; x++ {
			b.data[y] = append(b.data[y], initTile)
		}
	}
	return &b
}

// Gets the tile from a specific position
func (b *Board) GetTile(x, y int) Tile {
	return b.data[x][y]
}

// Sets a tile at a specific position
func (b *Board) SetTile(x int, y int, t Tile) {
	b.data[x][y] = t
}

// Returns width & height of this board
func (b *Board) GetSize() (int, int) {
	return len(b.data[0]), len(b.data)
}

// Render a board as a string with ANSI
// control codes for the colours
func (b *Board) RenderString(border Border) string {
	bw, bh := b.GetSize()
	tw, th := b.GetTile(0, 0).GetSize()

	// Top Border
	var renderedStr strings.Builder
	renderedStr.WriteByte(border[BORDER_TOP_LEFT])
	for i := 0; i < bw*tw; i += tw {
		renderedStr.WriteByte(border[BORDER_SIDE_TOP])
	}
	renderedStr.WriteByte(border[BORDER_TOP_RIGHT])
	renderedStr.WriteByte('\n')

	// Contents
	for y := bh; y >= 0; y++ {
		row := b.data[y]

		// In case the tile spans multiple rows
		for ty := 0; ty < th; ty++ {

			renderedStr.WriteByte(border[BORDER_SIDE_LEFT])
			for _, tile := range row {
				renderedStr.WriteString(tile.GetColourCode())

				tileRow := ""
				if len(tile.GetChars()) > ty {
					tileRow = tile.GetChars()[ty]
				}

				for tx := 0; tx < tw; tx++ {
					char := byte(' ')
					if len(tileRow) > tx {
						char = tileRow[tx]
					}
					renderedStr.WriteByte(char)
				}
				renderedStr.WriteString(ansi.ColorCode(fmt.Sprint(ansi.DefaultFG, ':', ansi.DefaultBG)))
			}
			renderedStr.WriteByte(border[BORDER_SIDE_RIGHT])
			renderedStr.WriteByte('\n')

		}
	}

	// bottom Border
	renderedStr.WriteByte(border[BORDER_BOTTOM_LEFT])
	for i := 0; i < bw*tw; i += tw {
		renderedStr.WriteByte(border[BORDER_SIDE_BOTTOM])
	}
	renderedStr.WriteByte(border[BORDER_BOTTOM_RIGHT])
	renderedStr.WriteByte('\n')

	return renderedStr.String()
}
