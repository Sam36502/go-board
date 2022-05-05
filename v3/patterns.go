/*
 *
 *	Provides pattern functions for
 *	filling in layers with given
 *	patterns
 *
 *	Can be used like so:
 *		lyr.FillPattern(PtrnCheckerboard(sq))
 *
 */

package board

type PatternFunc func(Coord) SquareRenderer

// Classic Checkerboard pattern
// Set offset to move board left/right for specific pattern starts
func PtrnCheckerboard(sq SquareRenderer, offset int) PatternFunc {
	return func(c Coord) SquareRenderer {
		if ((c.X+offset)%2 == 0 && c.Y%2 == 0) || ((c.X+offset)%2 != 0 && c.Y%2 != 0) {
			return sq
		} else {
			return nil
		}
	}
}

// Horizontal Stripe pattern
// Set offset to change which line starts the pattern
func PtrnStripesHoriz(sq SquareRenderer, offset int) PatternFunc {
	return func(c Coord) SquareRenderer {
		if (c.Y+offset)%2 == 0 {
			return sq
		} else {
			return nil
		}
	}
}

// Vertical Stripe pattern
// Set offset to change which line starts the pattern
func PtrnStripesVert(sq SquareRenderer, offset int) PatternFunc {
	return func(c Coord) SquareRenderer {
		if (c.X+offset)%2 == 0 {
			return sq
		} else {
			return nil
		}
	}
}
