package games

import (
	"encoding/json"
	"fmt"
)

// Simple2DBoard is a simple 2D board with a 2D string array
type Simple2DBoard struct {
	board  [][]string
	height int
	width  int
}

// JSONSimple2DBoard is used when generating a JSON representation
type JSONSimple2DBoard struct {
	Board  [][]string
	Height int
	Width  int
}

func (js JSONSimple2DBoard) Simple2DBoard() Simple2DBoard {
	return Simple2DBoard{board: js.Board, height: js.Height, width: js.Width}
}

// NewSimple2DBoard initializes a Simple2DBoard given two dimensions
func NewSimple2DBoard(m, n int) (*Simple2DBoard, error) {
	if n < 0 || m < 0 {
		return nil, fmt.Errorf("Both width and height must be positive")
	}
	b := new(Simple2DBoard)
	b.board = make([][]string, m)
	for i := range b.board {
		b.board[i] = make([]string, n)
	}
	b.height = n
	b.width = m
	return b, nil
}

func NewJSONSimple2DBoard(s Simple2DBoard) JSONSimple2DBoard {
	return JSONSimple2DBoard{s.board, s.height, s.width}
}

// Set gaming piece at Position p
func (b *Simple2DBoard) Set(x, y int, s string) {
	b.board[y][x] = s
}

// Get gaming piece Position p
func (b *Simple2DBoard) Get(x, y int) string {
	return b.board[y][x]
}

//IsEmpty returns true is the board at position x,y
//is not occupied by a game piece
func (b *Simple2DBoard) IsEmpty(x, y int) bool {
	return b.board[y][x] == ""
}

//Width implements the Board infetface
func (b *Simple2DBoard) Width() int {
	return b.width
}

//Height implements the Board infetface
func (b *Simple2DBoard) Height() int {
	return b.height
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
	b.board = make([][]string, cap(b.board))
	for i := range b.board {
		b.board[i] = make([]string, cap(b.board[0]))
	}
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

// MarshalJSON implements the Marshaler interface
func (s *Simple2DBoard) MarshalJSON() ([]byte, error) {
	return json.Marshal(NewJSONSimple2DBoard(*s))
}

func (s *Simple2DBoard) UnmarshalJSON(data []byte) error {
	var js JSONSimple2DBoard
	if err := json.Unmarshal(data, &js); err != nil {
		return err
	}
	*s = js.Simple2DBoard()
	return nil
}
