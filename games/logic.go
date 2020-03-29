package games

import "fmt"

type Action int

const (
	Place Action = iota
	Remove
	Move
)

type Direction int

const (
	LeftRight Direction = iota
	RightLeft
)

//GameLogic defines the functions required to decide
type GameLogic interface {

	//IsOver returns true, if a winner exists or there are no moves left
	IsOver() bool

	//IsOver returns true if Action with piece s can be conducted at the coords
	IsLegal(a Action, p *Player, coords ...int) bool

	//MovesRemaining returns an array of all remaining positions that the given player
	//can set a game piece on
	MovesRemaining() int

	//GetWinner returns a pointer to the player that has won the game according
	//to the internal rules
	GetWinner() *Player
}

type BaseLogic struct {
	players map[string]*Player
	board   Board
}

//NewBaseLogic returns an initialized BaseLogic struct
func NewBaseLogic(b Board, players ...*Player) (*BaseLogic, error) {
	bl := new(BaseLogic)
	var p *Player

	if players[0] == players[1] {
		return nil, fmt.Errorf("you must supply two different players")
	}
	if players[0].Symbol == players[1].Symbol {
		return nil, fmt.Errorf("the two players must not have the same symbol")
	}

	bl.board = b
	bl.players = make(map[string]*Player)
	for _, p = range players {
		bl.players[p.Symbol] = p
	}
	return bl, nil
}

//MovesRemaining implements the GameLogic interface
func (bl *BaseLogic) MovesRemaining() int {
	moves := 0
	for y := 0; y < bl.board.Height(); y++ {
		for x := 0; x < bl.board.Width(); x++ {
			if bl.board.IsEmpty(x, y) {
				moves++
			}
		}
	}
	return moves
}

//IsOver implement the GameLogic interface
func (bl *BaseLogic) IsOver() bool {
	return bl.GetWinner() != nil || bl.MovesRemaining() == 0
}

//IsLegal impelments the GameLogic interface
func (bl *BaseLogic) IsLegal(a Action, p *Player, coords ...int) bool {
	var legal bool

	switch a {
	case Place:
		//		fmt.Printf("Testing if Placement is legal: Position %v Player %v\n", coords, p)
		legal = bl.board.IsEmpty(coords[0], coords[1])
	case Remove:
		//		fmt.Printf("Testing if Removing is legal: Position %v Player %v\n", coords, p)
		//we can only remove a piece, if there is a piece and
		//we only may remove a piece, if it is ours
		legal = !bl.board.IsEmpty(coords[0], coords[1]) && bl.board.Get(coords[0], coords[1]) == p.Symbol
	case Move:
		//		fmt.Printf("Testing if Moving is legal: From %v To %v Player %v\n", coords[0:2], coords[2:], p)
		// see if we may take whats in place A and move it to place B
		legal = bl.IsLegal(Remove, p, coords[0], coords[1]) && bl.IsLegal(Place, p, coords[2], coords[3])
	default:
		legal = false
	}
	return legal
}

//GetWinner implements the GameLogic interface
func (bl *BaseLogic) GetWinner() *Player {
	var winner *Player
	//check horizontally
	winner = bl.checkHorizontally()

	if winner == nil {
		//check vertically
		winner = bl.checkVertically()
	}

	if winner == nil {
		//check diagnoally
		winner = bl.checkDiagonally()
	}
	return winner
}

func (bl *BaseLogic) checkHorizontally() *Player {
	var winner *Player
	var s string
	var streaking string = ""
	streakLen := 0
	//fmt.Printf("---- Started Horizontal Check (%d,%d)----\n", bl.board.Height(), bl.board.Width())
	for x := 0; x < bl.board.Height(); x++ {
		//iterate over all fields in a row until
		for y := 0; y < bl.board.Width(); y++ {
			//if you find a non-empty field
			if !(bl.board.IsEmpty(x, y)) {
				s = bl.board.Get(x, y)
				//if the symbol machtes the current streak
				if s == streaking {
					//increase length of current streak
					streakLen++
					//fmt.Printf("Increased streakLen to %d\n", streakLen)
				} else {
					//start a new streak with that symbol
					//fmt.Printf("Started new streak for: %s. Starting at (%d,%d)\n", s, x, y)
					streaking = s
					streakLen = 1
				}
			} else {
				//fmt.Printf("Resetting streaks\n")
				streaking = ""
				streakLen = 0
			}
			if streakLen == 3 {
				return bl.players[streaking]
			}
		} //end y
	} //end x

	//fmt.Printf("---- Finished Horizontal Check ----\n")
	//return a pointer to the player that has the winning symbol
	return winner
}

func (bl *BaseLogic) checkVertically() *Player {
	var winner *Player
	var s string
	var streaking string = ""
	streakLen := 0
	//fmt.Printf("---- Started Vertical Check ----\n")
	for y := 0; y < bl.board.Width(); y++ {
		//iterate over all fields in a row until
		for x := 0; x < bl.board.Height(); x++ {
			//if you find a non-empty field
			if !(bl.board.IsEmpty(x, y)) {
				s = bl.board.Get(x, y)
				//if the symbol machtes the current streak
				if s == streaking {
					//increase length of current streak
					streakLen++
					//fmt.Printf("Increased streakLen to %d\n", streakLen)
				} else {
					//start a new streak with that symbol
					//fmt.Printf("Started new streak for: %s. Starting at (%d,%d)\n", s, x, y)
					streaking = s
					streakLen = 1
				}
			} else {
				//fmt.Printf("Resetting streaks\n")
				streaking = ""
				streakLen = 0
			}
			if streakLen == 3 {
				return bl.players[streaking]
			}
		}
	}
	//fmt.Printf("---- Finished Vertical Check ----\n")
	//return a pointer to the player that has the winning symbol
	return winner
}

func (bl *BaseLogic) checkDiagonally() *Player {
	var diagonal []string
	var p *Player
	for y := 0; y < bl.board.Height(); y++ {
		for x := 0; x < bl.board.Width(); x++ {
			diagonal = bl.getDiagonal(x, y, LeftRight)
			p = bl.checkSlice(diagonal)
			if p != nil {
				return p
			}
			diagonal = bl.getDiagonal(x, y, RightLeft)
			p = bl.checkSlice(diagonal)
			if p != nil {
				return p
			}
		}
	}
	return nil
}

func (bl *BaseLogic) checkSlice(s []string) *Player {
	var win *Player
	fmt.Printf("CheckSlice: %v\n", s)
	var streaking string
	streakLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] != streaking {
			streaking = s[i]
			streakLen = 1
		} else {
			streakLen++
		}
	}
	if streakLen == 3 {
		win = bl.players[streaking]
	}
	return win
}

func (bl *BaseLogic) getDiagonal(x, y int, d Direction) []string {
	var pieces []string
	var x1, y1 int
	if d == LeftRight {
		fmt.Printf("Getting the LeftRight diagonal starting at (%d,%d)\n", x, y)
		for x1, y1 = x, y; x1 < bl.board.Height() && y1 < bl.board.Width(); x1, y1 = x1+1, y1+1 {
			pieces = append(pieces, bl.board.Get(x1, y1))
		}
	} else if d == RightLeft {
		fmt.Printf("Getting the RightLeft diagonal starting at (%d,%d)\n", x, y)
		for x1, y1 = x, y; x1 < bl.board.Height() && y1 >= 0; x1, y1 = x1+1, y1-1 {
			pieces = append(pieces, bl.board.Get(x1, y1))
		}
	}
	fmt.Printf("Diagonal (%d,%d):%v\n", x, y, pieces)
	return pieces
}
