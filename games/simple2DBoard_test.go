package games

import "testing"
import "fmt"

func TestBoardInit(t *testing.T) {
	t.Log("Creating 3x3 board")
	b := new(Simple2DBoard)
	b.Init(3, 3)
	h := cap(b.board)
	if h != 3 {
		t.Error(fmt.Sprintf("Width incorrect after call to Init. Expected: %d , got %d", 3, h))
	}
	w := cap(b.board[0])
	if w != 3 {
		t.Error(fmt.Sprintf("Width incorrect after call to Init. Expected: %d , got %d", 3, w))
	}
}

func TestSet(t *testing.T) {
	t.Log("Creating 3x3 board")
	b := new(Simple2DBoard)
	b.Init(3, 3)
	b.Set(0, 0, "x")
	s := b.board[0][0]
	if s != "x" {
		t.Error(fmt.Sprintf("Invalid result after call to Set"))
	}
}

func TestGet(t *testing.T) {
	t.Log("Creating 3x3 board")
	b := new(Simple2DBoard)
	b.Init(3, 3)
	b.board[0][0] = "x"
	s := b.Get(0, 0)
	if s != "x" {
		t.Error(fmt.Sprintf("Invalid resuult returned by Get"))
	}
}

func TestRemove(t *testing.T) {
	t.Log("Creating 3x3 board")
	b := new(Simple2DBoard)
	b.Init(3, 3)
	b.board[0][0] = "x"
	b.Remove(0, 0)
	if b.board[0][0] == "x" {
		t.Error("Remove not successful.")
	}
}

func TestMove(t *testing.T) {
	t.Log("Creating 3x3 board")
	b := new(Simple2DBoard)
	b.Init(3, 3)
	b.board[0][0] = "x"
	b.Move(0, 0, 0, 1)
	if b.board[0][0] == "x" {
		t.Error("Move not successful: Original piece not removed.")
	}
	if b.board[0][1] != "x" {
		t.Error("Move not successful: New piece not placed.")
	}
}

func TestHeight(t *testing.T) {
	t.Log("Creating 3x4 board")
	b := new(Simple2DBoard)
	exp := 4
	b.Init(3, exp)
	h := b.Height()
	if h != exp {
		t.Error(fmt.Sprintf("Height incorrect after call to Init. Expected: %d , got %d", exp, h))
	}
}

func TestWidth(t *testing.T) {
	t.Log("Creating 3x4 board")
	b := new(Simple2DBoard)
	exp := 3
	b.Init(exp, 4)
	w := b.Width()
	if w != exp {
		t.Error(fmt.Sprintf("Height incorrect after call to Init. Expected: %d , got %d", exp, w))
	}
}

func TestReset(t *testing.T) {
	t.Log("Creating 3x3 board")
	b := new(Simple2DBoard)
	b.Init(3, 3)

	b.board[0][0] = "x"
	b.board[1][1] = "x"
	b.board[2][2] = "x"

	addr := &b.board
	b.Reset()
	if &b.board != addr {
		t.Error("Reset failed. Board set to new array")
	}
}
