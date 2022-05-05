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

import (
	"math/rand"
	"time"
)

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

// Random Patter
// Can optionally provide a Seed (-1 for no seed)
// Can also determine the percent of results to be set
func PtrnRandom(sq SquareRenderer, seed int64, fillPercent float32) PatternFunc {
	if seed != -1 {
		rand.Seed(seed)
	} else {
		rand.Seed(time.Now().Unix())
	}
	return func(c Coord) SquareRenderer {
		if rand.Float32() <= fillPercent {
			return sq
		} else {
			return nil
		}
	}
}
