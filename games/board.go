package games

//Player has a name and Symbol
type Player struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

//Board defines functions that a generic game board should implement.
//A board is considered dumb. It does not implement any game logic.
//It is merely a wrapper around the internal data structre.
//The logic of the game should be enforced uinsg an implementation of the
//GameLogic interface
type Board interface {

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
