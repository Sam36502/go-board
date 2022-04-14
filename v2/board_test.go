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

	brd := NewBoard(8, 8, DoubleBorder, tWhite)
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

	fmt.Print(brd.RenderString())

	knight := &Pixel{
		Colour: Colour{
			Black,
			Red,
		},
		Chars: []string{
			"K1",
		},
	}

	StartPos := Coord{3, 1}
	KnightMove := Vector{1, 2}

	brd.SetPiece(StartPos, knight)
	brd.SetPiece(StartPos.Add(KnightMove.Scale(2)), knight)

	// Try to move the piece to the same position as the other one
	fmt.Print(brd.RenderString())
	brd.MovePiece(StartPos, KnightMove.Scale(2))
	fmt.Print(brd.RenderString())

	// Remove the piece blocking the move
	brd.DeletePiece(StartPos.Add(KnightMove.Scale(2)))

	ClearScreen()

	// Try the move again
	fmt.Print(brd.RenderString())
	brd.MovePiece(StartPos, KnightMove.MirrorX())
	fmt.Print(brd.RenderString())

}
