package games

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestNewSimple2DBoard(t *testing.T) {
	t.Log("Creating 3x3 board")
	b, _ := NewSimple2DBoard(3, 3)
	h := cap(b.board)
	if h != 3 {
		t.Error(fmt.Sprintf("Width incorrect after call to Init. Expected: %d , got %d", 3, h))
	}
	w := cap(b.board[0])
	if w != 3 {
		t.Error(fmt.Sprintf("Width incorrect after call to Init. Expected: %d , got %d", 3, w))
	}

	b1, _ := NewSimple2DBoard(-1, 3)
	if b1 != nil {
		t.Error("The width of a Simple2DBoard must not be negative")
	}

	b2, _ := NewSimple2DBoard(3, -1)
	if b2 != nil {
		t.Error("The height of a Simple2DBoard must not be negative")
	}
}

func TestSet(t *testing.T) {
	t.Log("Creating 3x3 board")
	b, _ := NewSimple2DBoard(3, 3)
	b.Set(0, 0, "x")
	s := b.Get(0, 0)
	if s != "x" {
		t.Error(fmt.Sprintf("Invalid result after call to Set"))
	}
}

func TestGet(t *testing.T) {
	t.Log("Creating 3x3 board")
	b, _ := NewSimple2DBoard(3, 3)
	b.Set(0, 0, "x")
	s := b.Get(0, 0)
	if s != "x" {
		t.Error(fmt.Sprintf("Invalid resuult returned by Get"))
	}
}

func TestRemove(t *testing.T) {
	t.Log("Creating 3x3 board")
	b, _ := NewSimple2DBoard(3, 3)
	b.Set(0, 0, "x")
	b.Remove(0, 0)
	if b.Get(0, 0) == "x" {
		t.Error("Remove not successful.")
	}
}

func TestMove(t *testing.T) {
	t.Log("Creating 3x3 board")
	b, _ := NewSimple2DBoard(3, 3)
	b.Set(0, 0, "x")
	b.Move(0, 0, 0, 1)
	if b.Get(0, 0) == "x" {
		t.Error("Move not successful: Original piece not removed.")
	}
	if b.Get(0, 1) != "x" {
		t.Error("Move not successful: New piece not placed.")
	}
}

func TestHeight(t *testing.T) {
	t.Log("Creating 3x4 board")
	exp := 4
	b, _ := NewSimple2DBoard(3, exp)
	h := b.Height()
	if h != exp {
		t.Error(fmt.Sprintf("Height incorrect after call to Init. Expected: %d , got %d", exp, h))
	}
}

func TestWidth(t *testing.T) {
	t.Log("Creating 3x4 board")
	exp := 3
	b, _ := NewSimple2DBoard(exp, 3)
	w := b.Width()
	if w != exp {
		t.Error(fmt.Sprintf("Height incorrect after call to Init. Expected: %d , got %d", exp, w))
	}
}

func TestReset(t *testing.T) {
	t.Log("Creating 3x3 board")
	b, _ := NewSimple2DBoard(3, 3)

	b.board[0][0] = "x"
	b.board[1][1] = "x"
	b.board[2][2] = "x"

	addr := &b.board
	b.Reset()
	if &b.board != addr {
		t.Error("Reset failed. Board set to new array")
	}

	if &b.board == nil {
		t.Error("Reset failed. Board set to new array")
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Log("Creating 3x3 board")
	b, _ := NewSimple2DBoard(3, 3)

	b.board[0][0] = "x"
	b.board[1][1] = "x"
	b.board[2][2] = "x"

	j, _ := json.Marshal(b)
	fmt.Println(string(j))

}

func TestUnMarshalJSON(t *testing.T) {
	var e *Simple2DBoard
	var a Simple2DBoard

	e, _ = NewSimple2DBoard(3, 3)

	e.board[0][0] = "x"
	e.board[1][1] = "x"
	e.board[2][2] = "x"

	j := []byte(`{"Board":[
		["x","",""],
		["","x",""],
		["","","x"]], "Width":3, "Height":3 }`)
	err := json.Unmarshal(j, &a)

	if err != nil || !reflect.DeepEqual(e, &a) {
		t.Errorf("Unmarshalling previously marshalled Simpel2DBoard failed.JSON: %s Expected: \n%v\n (%T) got \n%v\n (%T)", j, e, e, &a, &a)
	}

	eh := 3
	if a.Height() != eh {
		t.Errorf("Unmarshaling failed: Board Height is incorrect. Exepected: %d Received: %d", eh, a.Height())
	}

	ew := 3
	if a.Width() != ew {
		t.Errorf("Unmarshaling failed: Board Width is incorrect. Exepected: %d Received: %d", ew, a.Width())
	}
}
