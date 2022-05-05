package board

import (
	"fmt"
	"testing"
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

	brd := NewBoard(8, 8, ASCIIBevelBorder)

	// Squares
	sqHidden := NewSquare(
		2, 1,
		Colour{
			Foreground: Black,
			Background: Cyan,
		},
		[]string{},
	)
	sqWhite := NewSquare(
		2, 1,
		Colour{
			Foreground: Black,
			Background: White,
		},
		[]string{},
	)
	sqBlack := NewSquare(
		2, 1,
		Colour{
			Foreground: Black,
			Background: Black,
		},
		[]string{},
	)

	lyHidden := brd.CreateLayer()
	lyHidden.FillLayer(sqHidden)
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
	lyBlack.FillPattern(func(c Coord) SquareRenderer {
		if (c.X%2 == 0 && c.Y%2 == 0) || (c.X%2 != 0 && c.Y%2 != 0) {
			return sqBlack
		} else {
			return nil
		}
	})

	brd.SetLayer(0, lyHidden)
	brd.SetLayer(1, lyWhite)
	brd.SetLayer(2, lyBlack)

	brd.PrintBoard()

}
