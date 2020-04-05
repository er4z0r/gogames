package games

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

	//WhoseTurn returns a pointer to the player,
	//that may make the next move
	WhoseTurn() *Player

	//MovesRemaining returns an array of all remaining positions that the given player
	//can set a game piece on
	MovesRemaining() int

	//GetWinner returns a pointer to the player that has won the game according
	//to the internal rules
	GetWinner() *Player

	//StartRound notifies the game logic that a new round has begun
	BeginTurn()

	//NextRound notifies the game logic that a round has ended
	EndTurn()
}
