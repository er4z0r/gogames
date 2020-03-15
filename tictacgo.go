package main

import (
	"fmt"

	. "github.com/er4z0r/tictacgo/games"
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

	b.Set(0, 0, "o")
	b.Set(0, 1, "x")
	b.Set(0, 2, "x")

	b.Set(1, 0, "x")
	b.Set(1, 1, "o")
	b.Set(1, 2, "x")

	b.Set(2, 0, "x")
	b.Set(2, 1, "x")
	b.Set(2, 2, "o")

	fmt.Printf("%v\n", b)
}
