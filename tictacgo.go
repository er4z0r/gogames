package main

import (
	"fmt"

	. "github.com/er4z0r/tictacgo/games"
)

func clear() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	fmt.Println("Hello Tic Tac Go!")
	var x, y int

	b, _ := NewSimple2DBoard(3, 3)

	//Prompt for player names
	p1 := new(Player)
	p2 := new(Player)

	fmt.Println("Player1 Name:")
	fmt.Scan(&p1.Name)
	p1.Symbol = "o"

	fmt.Println("Player2 Name:")
	fmt.Scan(&p2.Name)
	p2.Symbol = "x"

	var g GameLogic
	g, _ = NewBaseLogic(b, p1, p2)
	for {
		clear()
		fmt.Printf("\n%v\n", b)

		//While the game is not over
		if g.IsOver() {
			break
		}
		p := g.WhoseTurn()
		//Prompt Player for coordinates
		fmt.Printf("%s's turn. Please enter coordinates:", p.Name)
		fmt.Scanln(&x, &y)

		if g.IsLegal(Place, p, x, y) {
			b.Set(x, y, p.Symbol)
		}
	}

	//Identify the winner
	clear()
	fmt.Println("\t== GAME OVER! ==")
	fmt.Printf("\n%v\n", b)
	if w := g.GetWinner(); w != nil {
		fmt.Printf("Player %s wins!\n", w.Name)
	} else {
		fmt.Println("It is a draw. Nobody wins!")
	}

}
