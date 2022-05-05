package board

import (
	"fmt"
	"testing"
)

var (
	sqCyan = NewSquare(
		2, 1,
		Colour{
			Foreground: Black,
			Background: Cyan,
		},
		[]string{},
	)
	sqWhite = NewSquare(
		2, 1,
		Colour{
			Foreground: Black,
			Background: White,
		},
		[]string{},
	)
	sqBlack = NewSquare(
		2, 1,
		Colour{
			Foreground: Black,
			Background: Black,
		},
		[]string{},
	)
)

func TestRenderSquare(t *testing.T) {

	squa := NewSquare(
		SQUARE_WIDTH,
		SQUARE_HEIGHT,
		Colour{
			Foreground: Blue,
			Background: Cyan,
		},
		[]string{
			"~~outside",
			"the render area",
		},
	)
	PrintANSIString(squa.RenderString())
	fmt.Println()

	squa = NewSquare(
		5,
		5,
		Colour{
			Foreground: Blue,
			Background: Cyan,
		},
		[]string{
			"~~outside",
			"the render area",
		},
	)
	PrintANSIString(squa.RenderString())

}

func TestCheckerboard(t *testing.T) {

	brd := NewBoard(8, 8, BrdrASCIIBevel)

	// Squares

	lyHidden := brd.CreateLayer()
	lyHidden.FillLayer(sqCyan)
	lyHidden.FillArea(
		Coord{4, 4},
		Coord{7, 7},
		nil,
	)

	lyWhite := brd.CreateLayer()
	lyWhite.FillLayer(sqWhite)
	lyWhite.FillArea(
		Coord{4, 4},
		Coord{7, 7},
		nil,
	)
	lyWhite.FillArea(
		Coord{0, 0},
		Coord{3, 3},
		nil,
	)

	lyBlack := brd.CreateLayer()
	lyBlack.FillPattern(PtrnCheckerboard(sqBlack, 1))

	brd.SetLayer(0, lyHidden)
	brd.SetLayer(1, lyWhite)
	brd.SetLayer(2, lyBlack)

	brd.PrintBoard()

}

func TestPatterns(t *testing.T) {

	brd := NewBoard(10, 10, BrdrSingle)
	back := brd.CreateLayer()
	back.FillLayer(sqWhite)
	brd.SetLayer(0, back)

	lyPattern := brd.CreateLayer()
	brd.SetLayer(1, lyPattern)

	lyPattern.FillPattern(PtrnStripesHoriz(sqBlack, 0))
	brd.PrintBoard()
	lyPattern.FillLayer(nil)

	lyPattern.FillPattern(PtrnStripesVert(sqBlack, 0))
	brd.PrintBoard()
	lyPattern.FillLayer(nil)

	lyPattern.FillPattern(PtrnRandom(sqBlack, -1, 0.2))
	brd.PrintBoard()
	lyPattern.FillLayer(nil)

}

func TestPieces(t *testing.T) {

	brd := NewBoard(7, 7, BrdrDouble)

	lyBackground := brd.CreateLayer()
	lyBackground.FillLayer(sqBlack)

	lyPieces := NewPieceLayer(brd.GetWidth(), brd.GetHeight(), brd.GetBorder())

	odys := NewSquare(
		2, 1,
		sqWhite.Colour,
		[]string{
			"oo",
		})
	agam := NewSquare(
		2, 1,
		sqCyan.Colour,
		[]string{
			"oo",
		})

	lyPieces.SetPiece("Odysseus", Coord{1, 1}, odys)
	lyPieces.SetPiece("Agammemnon", Coord{1, 1}, agam)

	brd.PrintBoard()

}
