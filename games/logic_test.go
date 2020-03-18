package games

import (
	"testing"
)

func TestNewBaseLogic(t *testing.T) {

	var b Board
	b, _ = NewSimple2DBoard(3, 4)

	p1 := Player{Name: "Alice", Symbol: "o"}
	p2 := Player{Name: "Bob", Symbol: "x"}

	l, _ := NewBaseLogic(&b, &p1, &p2)
	if l.board != &b {
		t.Errorf("NewBaseLogic failed. Pointer to board not assigned correctly. ")
	}

	//test that we may not use the same player twice
	pa := Player{Name: "Alice", Symbol: "o"}
	l1, _ := NewBaseLogic(&b, &pa, &pa)
	if l1 != nil {
		t.Errorf("NewBaseLogic failed. Supplied (%v,%v,%v) Expected: %v Received: %v", &b, &pa, &pa, nil, l1)
	}

	//test that we may not use players with the same symbol
	pa2 := Player{Name: "Alice", Symbol: "o"}
	pb2 := Player{Name: "Bob", Symbol: "o"}
	l2, _ := NewBaseLogic(&b, &pa2, &pb2)
	if l2 != nil {
		t.Errorf("NewBaseLogic failed. Supplied two players with same symbol. Expected: %v Received: %v", nil, l2)
	}

	e := 2
	a := len(l.players)
	if a != e {
		t.Errorf("NewBaseLogic failed. Players not assigned corretly. Expected len(l.players)=%d. Got %d\n", e, a)
		t.Errorf("%v", l.players)
	}

	if l.players[p1.Symbol] != &p1 || l.players[p2.Symbol] != &p2 {
		t.Error("NewBaseLogic failed. Players not stored correctly")
	}
}

//MovesRemaining returns an array of all remaining positions that the given player
//can set a game piece on
func TestMovesRemaining(t *testing.T) {
	var b Board
	b, _ = NewSimple2DBoard(2, 2)

	p1 := Player{Name: "Alice", Symbol: "o"}
	p2 := Player{Name: "Bob", Symbol: "x"}

	l, _ := NewBaseLogic(&b, &p1, &p2)

	e1 := 4
	a1 := l.MovesRemaining()

	if e1 != a1 {
		t.Errorf("Remaining moves not correct. Expected: %d. Actual: %d", e1, a1)
		t.Logf("\n%v\n", b)

	}

	b.Set(0, 0, p1.Symbol)

	e2 := 3
	a2 := l.MovesRemaining()

	if e2 != a2 {
		t.Errorf("Remaining moves not correct. Expected: %d. Actual: %d", e1, a1)
		t.Logf("\n%v\n", b)

	}
}

//IsOver returns true, if a winner exists or there are no moves left
func TestIsOver(t *testing.T) {
	var b Board
	b, _ = NewSimple2DBoard(3, 3)

	p1 := Player{Name: "Alice", Symbol: "o"}
	p2 := Player{Name: "Bob", Symbol: "x"}

	l, _ := NewBaseLogic(&b, &p1, &p2)
	//Only test a draw as we are testing
	//the other cases in the TestCheck...
	//functions
	//o|x|o
	//x|o|x
	//o|x|o
	b.Set(0, 0, p1.Symbol)
	b.Set(0, 1, p2.Symbol)
	b.Set(0, 2, p1.Symbol)

	b.Set(1, 0, p2.Symbol)
	b.Set(1, 1, p1.Symbol)
	b.Set(1, 2, p2.Symbol)

	b.Set(2, 0, p1.Symbol)
	b.Set(2, 1, p2.Symbol)
	b.Set(2, 2, p1.Symbol)

	e := true
	a := l.IsOver()
	if e != a {
		t.Errorf("IsOver failed. The game has no moves left and no winner. Expected %t Got: %t", e, a)
		t.Logf("\n%v\n", b)
	}

}

func TestIsLegal(t *testing.T) {
	var b Board
	b, _ = NewSimple2DBoard(3, 3)

	p1 := Player{Name: "Alice", Symbol: "o"}
	p2 := Player{Name: "Bob", Symbol: "x"}

	l, _ := NewBaseLogic(&b, &p1, &p2)

	//test, that we may place at an empty field
	e := true
	a := l.IsLegal(Place, &p1, 0, 0)
	if a != e {
		t.Errorf("IsLegal failed. Placing (%d,%d,%s) returned %t, but expected %t", 0, 0, p1.Symbol, a, e)
		t.Logf("\n%v\n", b)
	}
	b.Set(0, 0, p1.Symbol)

	//test, that we may not place on an occupied field
	e = false
	a = l.IsLegal(Place, &p2, 0, 0)
	if a != e {
		t.Errorf("IsLegal failed. Placing (%d,%d,%s) returned %t, but expected %t", 0, 0, p2.Symbol, a, e)
		t.Logf("\n%v\n", b)
	}

	//test, that we may not remove a piece that is not ours
	e = false
	a = l.IsLegal(Remove, &p2, 0, 0)
	if a != e {
		t.Errorf("IsLegal failed. Removing (%d,%d,%s) as %s returned %t, but expected %t", 0, 0, p1.Symbol, p2.Symbol, a, e)
		t.Logf("\n%v\n", b)
	}

	//test, that we may remove a piece that is ours
	e = true
	a = l.IsLegal(Remove, &p1, 0, 0)
	if a != e {
		t.Errorf("IsLegal failed. Removing (%d,%d) returned %t, but expected %t", 0, 0, a, e)
		t.Logf("\n%v\n", b)
	}

	//test the we may move a piece that is ours to an empty field
	e = true
	a = l.IsLegal(Move, &p1, 0, 0, 0, 1)
	if a != e {
		t.Errorf("IsLegal failed. Moving (%d,%d) to (%d,%d) returned %t, but expected %t", 0, 0, 0, 1, a, e)
		t.Logf("\n%v\n", b)
	}

	//test the we may move a piece that not ours to an empty field
	e = false
	a = l.IsLegal(Move, &p2, 0, 0, 0, 1)
	if a != e {
		t.Errorf("IsLegal failed. Moving (%d,%d) to (%d,%d) returned %t, but expected %t", 0, 0, 0, 1, a, e)
		t.Logf("\n%v\n", b)
	}

	//test that we may not move a piece that is ours to o non-empty field
	b.Set(0, 1, p1.Symbol)
	e = false
	a = l.IsLegal(Move, &p1, 0, 0, 0, 1)
	if a != e {
		t.Errorf("IsLegal failed. Moving (%d,%d) to (%d,%d) returned %t, but expected %t", 0, 0, 0, 1, a, e)
		t.Logf("\n%v\n", b)
	}

	//test that we may not move a piece that is not there
	b.Remove(0, 1)
	b.Remove(0, 0)
	e = false
	a = l.IsLegal(Move, &p1, 0, 0, 0, 1)
	if a != e {
		t.Errorf("IsLegal failed. Moving (%d,%d) to (%d,%d) returned %t, but expected %t", 0, 0, 0, 1, a, e)
		t.Logf("\n%v\n", b)
	}

}

//GetWinner returns a pointer to the player that has won the game according
//to the internal rules
func TestCheckHorizontally(t *testing.T) {
	var b Board
	b, _ = NewSimple2DBoard(3, 3)

	p1 := Player{Name: "Alice", Symbol: "o"}
	p2 := Player{Name: "Bob", Symbol: "x"}

	l, _ := NewBaseLogic(&b, &p1, &p2)

	var e1 *Player
	b.Set(0, 0, p1.Symbol)
	a1 := l.checkHorizontally()
	if a1 != e1 {
		t.Errorf("Game ended too early. CheckHorizontally returned %p expected: %p", a1, e1)
		t.Logf("\n%v\n", b)

	}

	var e2 *Player
	b.Set(1, 0, p2.Symbol)
	a2 := l.checkHorizontally()
	if a2 != e2 {
		t.Errorf("Game ended too early. CheckHorizontally returned %p expected: %p", a2, e2)
		t.Logf("\n%v\n", b)

	}

	var e3 *Player
	b.Set(0, 1, p1.Symbol)
	a3 := l.checkHorizontally()
	if a3 != e3 {
		t.Errorf("Game ended too early. CheckHorizontally returned %p expected: %p", a3, e3)
		t.Logf("\n%v\n", b)

	}

	var e4 *Player
	b.Set(2, 0, p2.Symbol)
	a4 := l.checkHorizontally()
	if a4 != e4 {
		t.Errorf("Game ended too late. CheckHorizontally returned %p expected: %p", a4, e4)
		t.Logf("\n%v\n", b)

	}

	var e5 = &p1
	b.Set(0, 2, p1.Symbol)
	a5 := l.checkHorizontally()
	if a5 != e5 {
		t.Errorf("Game ended too late. CheckHorizontally returned %p expected: %p", a5, e5)
		t.Logf("\n%v\n", b)
	}
}

func TestCheckVerticaally(t *testing.T) {
	var b Board
	b, _ = NewSimple2DBoard(3, 3)

	p1 := Player{Name: "Alice", Symbol: "o"}
	p2 := Player{Name: "Bob", Symbol: "x"}

	l, _ := NewBaseLogic(&b, &p1, &p2)

	var e1 *Player
	b.Set(0, 0, p1.Symbol)
	a1 := l.checkVertically()
	if a1 != e1 {
		t.Errorf("Game ended too early. CheckVerticaally returned %p expected: %p", a1, e1)
		t.Logf("\n%v\n", b)

	}

	var e2 *Player
	b.Set(0, 1, p2.Symbol)
	a2 := l.checkVertically()
	if a2 != e2 {
		t.Errorf("Game ended too early. CheckVerticaally returned %p expected: %p", a2, e2)
		t.Logf("\n%v\n", b)

	}

	var e3 *Player
	b.Set(1, 0, p1.Symbol)
	a3 := l.checkVertically()
	if a3 != e3 {
		t.Errorf("Game ended too early. CheckVerticaally returned %p expected: %p", a3, e3)
		t.Logf("\n%v\n", b)

	}

	var e4 *Player
	b.Set(1, 1, p2.Symbol)
	a4 := l.checkVertically()
	if a4 != e4 {
		t.Errorf("Game ended too late. CheckVerticaally returned %p expected: %p", a4, e4)
		t.Logf("\n%v\n", b)

	}

	var e5 = &p1
	b.Set(2, 0, p1.Symbol)
	a5 := l.checkVertically()
	if a5 != e5 {
		t.Errorf("Game ended too late. CheckVerticaally returned %p expected: %p", a5, e5)
		t.Logf("\n%v\n", b)
	}
}
