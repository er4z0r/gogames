package games

//Position is a wrapper for coordinates on a board
type Position struct {
	X int
	Y int
}

//Player has a name and Symbol
type Player struct {
	Name   string
	Symbol string
}

//Board defines functions that a generic game board should implement
type Board interface {

	// Init initalizes the board with a width of m and height of n fields
	Init(m, n int)

	// Height returns the vertical dimension (number of fields)
	Height() int

	//  Width return the horizontal dimenstion (number of fields)
	Width() int

	// Set gaming piece
	Set(x, y int, s string)

	// Get gaming piece
	Get(x, y int) string

	// IsEmpty returns true, if no game piece is placed there
	IsEmpty(x, y int) bool

	// Remove gaming piece
	Remove(x, y int)

	// Move gaming piece
	Move(x1, y1, x2, y2 int)

	// Reset board
	Reset()
}
