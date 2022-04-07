/*
 *
 *	Simple coordinate struct
 *	for making movement easier
 *
 */
package board

import "math/rand"

const (
	INVERT_SCALAR = -1
)

// Direction vectors for each cardinal direction
var DIRECTIONS_CARD []Vector = []Vector{
	{0, 1},  // North
	{1, 0},  // East
	{0, -1}, // South
	{-1, 0}, // West
}

// Direction vectors for each diagonal direction
var DIRECTIONS_DIAG []Vector = []Vector{
	{1, 1},   // NE
	{1, -1},  // SE
	{-1, -1}, // SW
	{-1, 1},  // NW
}

// Direction vectors for every possible direction in a square grid
var DIRECTIONS_ALL []Vector = []Vector{
	{0, 1},   // North
	{1, 0},   // East
	{0, -1},  // South
	{-1, 0},  // West
	{1, 1},   // NE
	{1, -1},  // SE
	{-1, -1}, // SW
	{-1, 1},  // NW
}

// Holds position information
type Coord struct {
	X int
	Y int
}

// Holds position transform information
type Vector struct {
	X int
	Y int
}

// Returns a random vector from the list provided.
// Use with the various `DIRECTIONS` vars included.
func RandomDirection(directions []Vector) Vector {
	return directions[rand.Intn(len(directions))]
}

// Returns a random coordinate within the provided bounds.
func RandomPos(wid, hei int) Coord {
	return Coord{
		rand.Intn(wid),
		rand.Intn(hei),
	}
}

// Checks whether a coordiate is within given bounds
func (p *Coord) IsInBounds(wid, hei int) bool {
	return p.X >= 0 && p.X < wid && p.Y >= 0 && p.Y < hei
}

// Adds a vector to this vector and then returns the current vector
func (a *Coord) Add(b Vector) *Coord {
	a.X += b.X
	a.Y += b.Y
	return a
}

// Scales this vector and returns it
func (a *Coord) Scale(s int) *Coord {
	a.X *= s
	a.Y *= s
	return a
}

// Inverts this vector and returns it
func (a *Coord) Invert() *Coord {
	return a.Scale(INVERT_SCALAR)
}

// Converts this position to a vector
// originating in the bottom-left corner
func (c *Coord) ToVector() *Vector {
	return &Vector{
		c.X,
		c.Y,
	}
}

// Converts this vector to the position of the
// end of the vector placed in the bottom-left corner
func (v *Vector) ToCoord() *Coord {
	return &Coord{
		v.X,
		v.Y,
	}
}
