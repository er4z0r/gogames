package games

import (
	"testing"
)

func TestLogicInit(t *testing.T) {

	var b Board
	b = new(Simple2DBoard)
	b.Init(3, 4)

	p1 := Player{Name: "Alice", Symbol: "o"}
	p2 := Player{Name: "Bob", Symbol: "x"}

	l := new(BaseLogic)
	l.Init(&b, &p1, &p2)
	if l.board != &b {
		t.Errorf("Init failed. Pointer to board not assigned correctly. ")
	}

	//test that we may not use players with the same symbol
	l1 := new(BaseLogic)
	pa := Player{Name: "Alice", Symbol: "o"}
	pb := Player{Name: "Bob", Symbol: "o"}
	l1.Init(&b, &pa, &pb)

	e := 2
	a := len(l.players)
	if a != e {
		t.Errorf("Init failed. Players not assigned corretly. Expected len(l.players)=%d. Got %d\n", e, a)
		t.Errorf("%v", l.players)
	}

	if l.players[p1.Symbol] != &p1 || l.players[p2.Symbol] != &p2 {
		t.Error("Init failed. Players not stored correctly")
	}
}

//MovesRemaining returns an array of all remaining positions that the given player
//can set a game piece on
func TestMovesRemaining(t *testing.T) {
	var b Board
	b = new(Simple2DBoard)
	b.Init(2, 2)

	p1 := Player{Name: "Alice", Symbol: "o"}
	p2 := Player{Name: "Bob", Symbol: "o"}

	var l GameLogic
	l = new(BaseLogic)
	l.Init(&b, &p1, &p2)

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

}

func TestIsLegal(t *testing.T) {
	var b Board
	b = new(Simple2DBoard)
	b.Init(3, 3)

	p1 := Player{Name: "Alice", Symbol: "o"}
	p2 := Player{Name: "Bob", Symbol: "x"}

	var l GameLogic
	l = new(BaseLogic)
	l.Init(&b, &p1, &p2)

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
	b = new(Simple2DBoard)
	b.Init(3, 3)

	p1 := Player{Name: "Alice", Symbol: "o"}
	p2 := Player{Name: "Bob", Symbol: "x"}

	var l *BaseLogic
	l = new(BaseLogic)
	l.Init(&b, &p1, &p2)

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
	b = new(Simple2DBoard)
	b.Init(3, 3)

	p1 := Player{Name: "Alice", Symbol: "o"}
	p2 := Player{Name: "Bob", Symbol: "x"}

	var l *BaseLogic
	l = new(BaseLogic)
	l.Init(&b, &p1, &p2)

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
