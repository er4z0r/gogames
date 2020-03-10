package main

import (
	"fmt"

	. "github.com/er4z0r/games"
)

func main() {
	fmt.Println("Hello Tic Tac Go!")

	var b Board
	b = new(Simple2DBoard)

	//Prompt for player names

	//Prompt for board size
	b.Init(3, 3)

	//While the game is not over
	//Prompt Player for coordinates
	//Identify the winner

	b.Set(Position{X: 0, Y: 0}, "o")
	b.Set(Position{X: 0, Y: 1}, "x")
	b.Set(Position{X: 0, Y: 2}, "x")

	b.Set(Position{X: 1, Y: 0}, "x")
	b.Set(Position{X: 1, Y: 1}, "o")
	b.Set(Position{X: 1, Y: 2}, "x")

	b.Set(Position{X: 2, Y: 0}, "x")
	b.Set(Position{X: 2, Y: 1}, "x")
	b.Set(Position{X: 2, Y: 2}, "o")

	fmt.Printf("%v\n", b)
}
