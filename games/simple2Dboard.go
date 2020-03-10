package games

// Simple2DBoard is a simple 2D board with a 2D string array
type Simple2DBoard struct {
	board [][]string
}

// Init initializes a Simple2DBoard given two dimensions
func (b *Simple2DBoard) Init(m, n uint) {
	b.board = make([][]string, m)
	for i := range b.board {
		b.board[i] = make([]string, n)
	}
}

// Set gaming piece at Position p
func (b *Simple2DBoard) Set(p Position, s string) {
	b.board[p.X][p.Y] = s
}

// Get gaming piece Position p
func (b *Simple2DBoard) Get(p Position) string {
	return b.board[p.X][p.Y]
}

// Remove gaming piece
func (b *Simple2DBoard) Remove(p Position) {
	b.Set(p, ``)
}

// Move gaming piece
func (b *Simple2DBoard) Move(x, y Position) {
	b.Set(y, b.Get(x))
}

// Reset board
func (b *Simple2DBoard) Reset() {
	b.Init(uint(len(b.board)), uint(len(b.board[0])))
}
