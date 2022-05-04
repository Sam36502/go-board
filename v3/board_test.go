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
		2, 2,
		Colour{
			Foreground: Black,
			Background: Cyan,
		},
		[]string{},
	)
	sqWhite := NewSquare(
		2, 2,
		Colour{
			Foreground: Black,
			Background: White,
		},
		[]string{},
	)
	sqBlack := NewSquare(
		2, 2,
		Colour{
			Foreground: Black,
			Background: Black,
		},
		[]string{},
	)

	lyHidden := brd.CreateLayer()
	lyHidden.FillLayer(sqHidden)

	lyWhite := brd.CreateLayer()
	lyWhite.FillLayer(sqWhite)

	lyBlack := brd.CreateLayer()

	brd.SetLayer(0, hiddenLayer)
	brd.SetLayer(1, hiddenLayer)
	brd.SetLayer(2, hiddenLayer)

}
