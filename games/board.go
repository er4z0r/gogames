package games

//Position is a wrapper for coordinates on a board
type Position struct {
	X uint8
	Y uint8
}

//Board defines functions that a generic game board should implement
type Board interface {

	// Set gaming piece
	Set(p Position, s string) bool

	// Remove gaming piece
	Remove(p Position) bool

	// Move gaming piece
	Move(a, b Position) bool

	// Reset board
	Reset()
}
