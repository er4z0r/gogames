package games

import "fmt"

//GameLogic defines the functions required to decide
type GameLogic interface {

	//Init initializes
	Init(b *Board, players ...*Player)

	//IsOver returns true, if a winner exists or there are no moves left
	IsOver() bool

	//MovesRemaining returns an array of all remaining positions that the given player
	//can set a game piece on
	MovesRemaining() int

	//GetWinner returns a pointer to the player that has won the game according
	//to the internal rules
	GetWinner() *Player
}

type BaseLogic struct {
	players map[string]*Player
	board   *Board
}

//Init implements the GameLogic interface
func (bl *BaseLogic) Init(b *Board, players ...*Player) {
	var p *Player
	bl.board = b
	bl.players = make(map[string]*Player)
	for _, p = range players {
		bl.players[p.Symbol] = p
	}
}

//MovesRemaining implements the GameLogic interface
func (bl *BaseLogic) MovesRemaining() int {
	moves := 0
	var board Board
	board = *bl.board
	for y := 0; y < board.Height(); y++ {
		for x := 0; x < board.Width(); x++ {
			if (*bl.board).IsEmpty(x, y) {
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

//GetWinner implements the GameLogic interface
func (bl *BaseLogic) GetWinner() *Player {
	var winner *Player
	winner = nil
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
	var board Board
	var s string
	board = *bl.board
	var streaking string = ""
	streakLen := 0
	fmt.Printf("---- Started Horizontal Check ----\n")
	for x := 0; x < board.Height(); x++ {
		//iterate over all fields in a row until
		for y := 0; y < board.Width(); y++ {
			//if you find a non-empty field
			if !(board.IsEmpty(x, y)) {
				s = board.Get(x, y)
				//if the symbol machtes the current streak
				if s == streaking {
					//increase length of current streak
					streakLen++
					fmt.Printf("Increased streakLen to %d\n", streakLen)
				} else {
					//start a new streak with that symbol
					fmt.Printf("Started new streak for: %s. Starting at (%d,%d)\n", s, x, y)
					streaking = s
					streakLen = 1
				}
			} else {
				fmt.Printf("Resetting streaks\n")
				streaking = ""
				streakLen = 0
			}
			if streakLen == 3 {
				return bl.players[streaking]
			}
		} //end y
	} //end x

	fmt.Printf("---- Finished Horizontal Check ----\n")
	//return a pointer to the player that has the winning symbol
	return winner
}

func (bl *BaseLogic) checkVertically() *Player {
	var winner *Player
	var board Board
	var s string
	board = *bl.board
	var streaking string = ""
	streakLen := 0
	fmt.Printf("---- Started Vertical Check ----\n")
	for y := 0; y < board.Width(); y++ {
		//iterate over all fields in a row until
		for x := 0; x < board.Height(); x++ {
			//if you find a non-empty field
			if !(board.IsEmpty(x, y)) {
				s = board.Get(x, y)
				//if the symbol machtes the current streak
				if s == streaking {
					//increase length of current streak
					streakLen++
					fmt.Printf("Increased streakLen to %d\n", streakLen)
				} else {
					//start a new streak with that symbol
					fmt.Printf("Started new streak for: %s. Starting at (%d,%d)\n", s, x, y)
					streaking = s
					streakLen = 1
				}
			} else {
				fmt.Printf("Resetting streaks\n")
				streaking = ""
				streakLen = 0
			}
			if streakLen == 3 {
				return bl.players[streaking]
			}
		}
	}
	fmt.Printf("---- Finished Vertical Check ----\n")
	//return a pointer to the player that has the winning symbol
	return winner
}

func (bl *BaseLogic) checkDiagonally() *Player {
	return nil
}
