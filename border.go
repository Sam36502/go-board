/*
 *
 *	Border Character arrays
 *
 *	Defines how the edges of the rendered
 *	board should look, following this pattern:
 *
 *	011112
 *	7    3
 *	7    3
 *	655554
 *
 */
package board

type Border []rune

// Indices
const (
	BORDER_TOP_LEFT     = 0
	BORDER_SIDE_TOP     = 1
	BORDER_TOP_RIGHT    = 2
	BORDER_SIDE_RIGHT   = 3
	BORDER_BOTTOM_RIGHT = 4
	BORDER_SIDE_BOTTOM  = 5
	BORDER_BOTTOM_LEFT  = 6
	BORDER_SIDE_LEFT    = 7
)

var (
	NoBorder         = Border([]rune("        "))
	ASCIIBorder      = Border([]rune("+-+|+-+|"))
	ASCIIBevelBorder = Border([]rune("/-\\|/-\\|"))
	SingleBorder     = Border([]rune("┌─┐│┘─└│"))
	DoubleBorder     = Border([]rune("╔═╗║╝═╚║"))
)
