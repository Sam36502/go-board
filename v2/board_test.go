package board

import (
	"fmt"
	"testing"
)

var (
	tBlack *Pixel = &Pixel{
		Colour: Colour{
			White,
			Black,
		},
	}
	tWhite *Pixel = &Pixel{
		Colour: Colour{
			Black,
			White,
		},
	}
)

func TestCheckerboard(t *testing.T) {

	brd := NewBoard(8, 8, tWhite)
	blk := false
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if blk {
				brd.SetTile(Coord{x, y}, tBlack)
			}
			blk = !blk
		}
		blk = !blk
	}

	fmt.Print(brd.RenderString(DoubleBorder))

	knight := &Pixel{
		Colour: Colour{
			Black,
			Red,
		},
		Chars: []string{
			"K1",
		},
	}
	brd.SetPiece(Coord{3, 1}, knight)
	brd.SetPiece(Coord{4, 3}, knight)

	// Try to move the piece to the same position as the other one
	fmt.Print(brd.RenderString(DoubleBorder))
	brd.MovePiece(Coord{3, 1}, Vector{1, 2})
	fmt.Print(brd.RenderString(DoubleBorder))

	// Remove the piece blocking the move
	brd.DeletePiece(Coord{4, 3})

	// Try the move again
	fmt.Print(brd.RenderString(DoubleBorder))
	brd.MovePiece(Coord{3, 1}, Vector{1, 2})
	fmt.Print(brd.RenderString(DoubleBorder))

}
