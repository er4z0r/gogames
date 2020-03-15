package games

import "fmt"

// Simple2DBoard is a simple 2D board with a 2D string array
type Simple2DBoard struct {
	board [][]string
}

// Init initializes a Simple2DBoard given two dimensions
func (b *Simple2DBoard) Init(m, n int) {
	b.board = make([][]string, m)
	for i := range b.board {
		b.board[i] = make([]string, n)
	}
}

// Set gaming piece at Position p
func (b *Simple2DBoard) Set(x, y int, s string) {
	b.board[x][y] = s
}

// Get gaming piece Position p
func (b *Simple2DBoard) Get(x, y int) string {
	return b.board[x][y]
}

func (b *Simple2DBoard) IsEmpty(x, y int) bool {
	return b.board[x][y] == ""
}

//Width implements the Board infetface
func (b *Simple2DBoard) Width() int {
	return cap(b.board)
}

//Height implements the Board infetface
func (b *Simple2DBoard) Height() int {
	return cap(b.board[0])
}

// Remove gaming piece
func (b *Simple2DBoard) Remove(x, y int) {
	b.Set(x, y, ``)
}

// Move gaming piece
func (b *Simple2DBoard) Move(x1, y1, x2, y2 int) {
	s := b.Get(x1, y1)
	b.Remove(x1, y1)
	b.Set(x2, y2, s)
}

// Reset board
func (b *Simple2DBoard) Reset() {
	b.Init(cap(b.board), cap(b.board[0]))
}

func (b *Simple2DBoard) String() string {
	var ret string
	for i := range b.board {
		for j := range b.board[i] {
			if b.board[i][j] == "" {
				ret += fmt.Sprintf("| |")
			} else {
				ret += fmt.Sprintf("|%s|", b.board[i][j])
			}
		}
		ret += "\n"
	}
	return ret
}
